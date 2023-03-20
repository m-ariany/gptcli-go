package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	chatCommand = "chat"
)

func (shell ChatShell) Run() {
	rootCmd := &cobra.Command{
		Use:   chatCommand,
		Short: "ChatGPT CLI",
		Run:   shell.run,
	}

	rootCmd.Execute()
}

func (shell ChatShell) run(cmd *cobra.Command, args []string) {
	shell.writeToStdout(fmt.Sprintf("Type 'bye' to exit%s", shell.You))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		statement := strings.TrimSpace(scanner.Text())
		if len(statement) == 0 {
			continue
		}

		if statement == "bye" {
			break
		}

		shell.writeToStdout(shell.AI)
		shell.writeToStdout(shell.You)
	}
}

func (shell ChatShell) writeToStdout(text string) {
	io.Copy(os.Stdout, strings.NewReader(text))
}
