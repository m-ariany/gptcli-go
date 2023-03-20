package cli

import (
	"fmt"
)

type ChatShell struct {
	You string
	AI  string
}

func New() ChatShell {
	return ChatShell{
		You: fmt.Sprintf("\n\n%s: ", "You"),
		AI:  fmt.Sprintf("\n%s: ", "ChatGPT"),
	}
}
