package contract

import (
	"context"

	"github.com/m-ariany/gptcli/internal/entity"
)

type ChatGPTServer interface {
	Chat(context.Context, string) <-chan entity.Message
}
