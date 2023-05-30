package assistant

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/m-ariany/gptcli-go/internal/config"
	"github.com/rs/zerolog/log"
	ai "github.com/sashabaranov/go-openai"
)

const (
	Instruct   = "Instruct ChatGPT by giving context of conversation: "
	YouAndExit = `You (write 'exit' to exit): `
	You        = "You: "
	ChatGPT    = "ChatGPT: "
	ErrToUser  = "Oops! something went wrong, please try again."
)

type Assistant struct {
	config           config.ChatGPTConfig
	client           *ai.Client
	model            string
	scanner          *bufio.Scanner
	maxResponseToken int
	history          []ai.ChatCompletionMessage
}

func NewAssistant(cnf config.ChatGPTConfig) *Assistant {
	return &Assistant{
		config:           cnf,
		client:           ai.NewClient(cnf.APIKey),
		model:            cnf.Model,
		scanner:          bufio.NewScanner(os.Stdin),
		maxResponseToken: cnf.MaxTokens,
		// Typically, a conversation is formatted with a system message first,
		// followed by alternating user and assistant messages.
		// Ref: https://platform.openai.com/docs/guides/chat/introduction
		//
		// TODO: the system instructing message must be taken from the user
		history: []ai.ChatCompletionMessage{},
	}
}

func (a *Assistant) Run() {

	a.instruct()
	writeToStdout(YouAndExit)
	for a.scanner.Scan() {
		question := strings.TrimSpace(a.scanner.Text())
		if len(question) < 1 {
			continue
		}

		if question == "exit" {
			os.Exit(0)
		}

		a.chat(question)
	}
}

func (a *Assistant) instruct() {
	writeToStdout(Instruct)
	if a.scanner.Scan() {
		instruction := strings.TrimSpace(a.scanner.Text())
		a.history = append(a.history, ai.ChatCompletionMessage{
			Role:    "system",
			Content: instruction,
		})
	}
}

func (a *Assistant) chat(prompt string) {
	req := a.newChatCompletionRequest(prompt)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	answer, err := a.doChatCompletionStream(ctx, req)
	if !errors.Is(err, io.EOF) {
		return
	}

	a.history = append(a.history, ai.ChatCompletionMessage{
		Role:    "assistant",
		Content: answer,
	})
}

func (a *Assistant) newChatCompletionRequest(question string) ai.ChatCompletionRequest {

	/*
		Ref: https://platform.openai.com/docs/guides/chat/introduction
		Including the conversation history helps the models to give relevant answers to the prior conversation.
		Because the models have no memory of past requests, all relevant information must be supplied via the conversation.
	*/
	a.history = append(a.history, ai.ChatCompletionMessage{
		Role:    "user",
		Content: question,
	})

	return ai.ChatCompletionRequest{
		Model:     a.model,
		MaxTokens: a.maxResponseToken,
		Messages:  a.history,
		Stream:    true,
	}
}

func (a *Assistant) doChatCompletionStream(ctx context.Context, request ai.ChatCompletionRequest) (answer string, err error) {

	defer func() {
		writeToStdout(fmt.Sprintf("\n\n%s", You))
	}()

	resp, err := a.client.CreateChatCompletionStream(ctx, request)

	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		writeToStdout(ErrToUser)
		return
	}
	defer resp.Close()

	if resp.GetResponse().StatusCode >= 400 {
		statusCode := resp.GetResponse().StatusCode
		b, _ := ioutil.ReadAll(resp.GetResponse().Body)
		writeToStdout(fmt.Sprintf("%d %s - %s\n", statusCode, http.StatusText(statusCode), string(b)))
		os.Exit(1)
	}

	writeToStdout(ChatGPT)
	sb := strings.Builder{}
	for {
		var data ai.ChatCompletionStreamResponse
		data, err = resp.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				// TODO: Answer might be corrupted. Inform user about the error.
				log.Error().Err(err).Msg("stream error")
				return
			}
			break
		}
		respChunk := data.Choices[0].Delta.Content
		writeToStdout(respChunk)
		sb.WriteString(respChunk)
	}

	answer = sb.String()

	return
}

func writeToStdout(s string) {
	io.Copy(os.Stdout, strings.NewReader(s))
}
