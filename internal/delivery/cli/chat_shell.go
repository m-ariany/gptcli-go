package cli

import (
	"fmt"

	"github.com/m-ariany/gptcli-go/internal/config"

	"github.com/m-ariany/gptcli-go/internal/contract"
)

type ChatShell struct {
	You       string
	AI        string
	assistant contract.AssistantInteractor
}

func New(cfg config.ShellConfig, assistant contract.AssistantInteractor) ChatShell {
	return ChatShell{
		You:       fmt.Sprintf("\n\n%s: ", cfg.You),
		AI:        fmt.Sprintf("\n%s🤖\n\n ", cfg.AI),
		assistant: assistant,
	}
}
