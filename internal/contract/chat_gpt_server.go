package contract

import (
	"context"

	"github.com/m-ariany/gptcli-go/internal/entity"
)

type ChatGPTServer interface {
	Chat(context.Context, string) <-chan entity.Message
}
