package main

import (
	"fmt"

	"github.com/m-ariany/gptcli/internal/adapter/chatgpt"
	"github.com/m-ariany/gptcli/internal/config"
	"github.com/m-ariany/gptcli/internal/delivery/cli"
	"github.com/m-ariany/gptcli/internal/interactor/assistant"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	cfg := config.NewConfig()

	chatGPT := chatgpt.New(cfg.ChatGPT)

	assistant := assistant.New(chatGPT)

	shell := cli.New(cfg.Shell, assistant)
	shell.Run()
}
