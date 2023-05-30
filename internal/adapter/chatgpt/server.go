package chatgpt

import (
	"github.com/m-ariany/gptcli-go/internal/config"

	"github.com/sashabaranov/go-openai"
)

// Server implements contract.ChatGPTServer interface
type Server struct {
	cfg     config.ChatGPTConfig
	client  *openai.Client
	history []openai.ChatCompletionMessage
}

func New(cfg config.ChatGPTConfig) *Server {
	return &Server{
		cfg:    cfg,
		client: openai.NewClient(cfg.APIKey),
		// Typically, a conversation is formatted with an system message first,
		// followed by alternating user and assistant messages.
		// Ref: https://platform.openai.com/docs/guides/chat/introduction
		//
		// TODO: the system instructing message must be taken from the user
		history: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are a helpful AI that answers my questions",
			},
		},
	}
}
