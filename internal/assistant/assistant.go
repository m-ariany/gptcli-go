package assistant

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/m-ariany/gptcli/internal/config"
	"github.com/rs/zerolog/log"
	ai "github.com/sashabaranov/go-openai"
)

const (
	IceBreaker = "You (type `exit` to exit): "
	You        = "You: "
	ErrToUser  = "Oops! ChatGPT failed to responde. Please try again😅"
)

type Assistant struct {
	config   config.Config
	client   *ai.Client
	maxToken int
	history  []ai.ChatCompletionMessage
}

func NewAssistant(cnf config.Config) *Assistant {
	return &Assistant{
		config:   cnf,
		client:   ai.NewClient(cnf.ApiKey),
		maxToken: cnf.MaxToken,
		// Typically, a conversation is formatted with a system message first,
		// followed by alternating user and assistant messages.
		// Ref: https://platform.openai.com/docs/guides/chat/introduction
		//
		// TODO: the system instructing message must be taken from the user
		history: []ai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are a helpful AI that answers my questions to its best knowledge and candor",
			},
		},
	}
}

func (a *Assistant) Run() {

	writeToStdout(IceBreaker)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		question := strings.TrimSpace(scanner.Text())
		if len(question) < 1 {
			continue
		}

		if question == "exit" {
			os.Exit(0)
		}

		a.chat(question)
	}
}

func (a *Assistant) chat(prompt string) {
	req := a.newChatCompletionRequest(prompt)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	a.doChatCompletionStream(ctx, req)
}

func (a *Assistant) newChatCompletionRequest(prompt string) ai.ChatCompletionRequest {

	msg := ai.ChatCompletionMessage{
		Role:    "user",
		Content: prompt,
	}
	/*
		Ref: https://platform.openai.com/docs/guides/chat/introduction
		Including the conversation history helps the models to give relevant answers to the prior conversation.
		Because the models have no memory of past requests, all relevant information must be supplied via the conversation.
	*/
	a.history = append(a.history, msg)

	return ai.ChatCompletionRequest{
		Model:     ai.GPT4,
		MaxTokens: a.maxToken,
		Messages:  a.history,
		Stream:    true,
	}
}

func (a *Assistant) doChatCompletionStream(ctx context.Context, request ai.ChatCompletionRequest) {

	defer func() {
		writeToStdout(fmt.Sprintf("\n%s", You))
	}()

	resp, err := a.client.CreateChatCompletionStream(ctx, request)

	if err != nil {
		log.Error().Err(err)
		writeToStdout(ErrToUser)
		return
	}
	defer resp.Close()

	writeToStdout("ChatGPT: ")
	for {
		data, err := resp.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				log.Error().Err(err).Msg("stream error")
			}
			break
		}
		writeToStdout(data.Choices[0].Delta.Content)
	}
}

func writeToStdout(s string) {
	io.Copy(os.Stdout, strings.NewReader(s))
}
