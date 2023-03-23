package chatgpt

import (
	"context"
	"fmt"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/m-ariany/gptcli/internal/entity"

	"github.com/sashabaranov/go-openai"
)

func (server *Server) Chat(ctx context.Context, text string) <-chan entity.Message {
	message := make(chan entity.Message)

	go func() {
		defer close(message)

		request := server.newChatCompletionRequest(text)
		response, err := server.client.CreateChatCompletionStream(ctx, request)
		if err != nil {
			message <- entity.Message{Error: err}

			return
		}
		defer response.Close()

		statusCode := response.GetResponse().StatusCode
		if statusCode >= 400 {

			message <- entity.Message{
				Error: fmt.Errorf(
					"%d %s",
					statusCode,
					http.StatusText(statusCode),
				),
			}

			return
		}

		data := openai.ChatCompletionStreamResponse{}
		builder := strings.Builder{}
		for {
			data, err = response.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				message <- entity.Message{Error: err}

				return
			}

			chunk := data.Choices[0].Delta.Content
			builder.WriteString(chunk)
			message <- entity.Message{Data: chunk}
		}

		server.history = append(
			server.history,
			openai.ChatCompletionMessage{
				Role:    "assistant",
				Content: builder.String(),
			},
		)
	}()

	return message
}

func (server *Server) newChatCompletionRequest(
	statement string,
) openai.ChatCompletionRequest {
	// Ref: https://platform.openai.com/docs/guides/chat/introdution
	//
	// Including the conversation history helps the models to give relevantanswers
	// to the prior conversation.
	// Because the models have no memory of past requests, all relevant information
	// must be supplied via the conversation.
	server.history = append(
		server.history,
		openai.ChatCompletionMessage{
			Role:    "user",
			Content: statement,
		},
	)

	return openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: server.cfg.MaxResponseToken,
		Messages:  server.history,
		Stream:    true,
	}
}
