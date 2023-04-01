package contract

import (
	"context"

	"github.com/hjhussaini/gptcli-go/internal/entity"
)

type ChatGPTServer interface {
	Chat(context.Context, string) <-chan entity.Message
}
