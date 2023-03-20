package assistant

import (
	"github.com/m-ariany/gptcli/internal/contract"
)

type Assistant struct {
	server contract.ChatGPTServer
}

func New(server contract.ChatGPTServer) *Assistant {
	return &Assistant{
		server: server,
	}
}
