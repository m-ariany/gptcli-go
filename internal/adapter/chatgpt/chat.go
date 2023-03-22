package chatgpt

import (
	"context"
	"errors"

	"github.com/m-ariany/gptcli/internal/entity"

	"github.com/sashabaranov/go-openai"
)

func (server *Server) Chat(ctx context.Context, text string) <-chan entity.Message {
	message := make(chan entity.Message)

	go func() {
		defer close(message)

		message <- entity.Message{
			Error: errors.New("Unimplemented method"),
		}
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
