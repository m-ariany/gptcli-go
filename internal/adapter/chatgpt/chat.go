package chatgpt

import (
	"context"
	"errors"

	"github.com/m-ariany/gptcli/internal/entity"
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
