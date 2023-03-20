package cli

import (
	"fmt"

	"github.com/m-ariany/gptcli/internal/contract"
)

type ChatShell struct {
	You       string
	AI        string
	assistant contract.AssistantInteractor
}

func New(assistant contract.AssistantInteractor) ChatShell {
	return ChatShell{
		You:       fmt.Sprintf("\n\n%s: ", "You"),
		AI:        fmt.Sprintf("\n%s: ", "ChatGPT"),
		assistant: assistant,
	}
}
