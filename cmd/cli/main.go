package main

import (
	"fmt"

	"github.com/m-ariany/gptcli/internal/adapter/chatgpt"
	"github.com/m-ariany/gptcli/internal/config"
	"github.com/m-ariany/gptcli/internal/delivery/cli"
	"github.com/m-ariany/gptcli/internal/interactor/assistant"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	out, err := config.LogOutput()
	if err != nil {
		panic(err)
	}

	defer func() {
		if out != nil {
			_ = out.Close()
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("Oops! %v", err)
			log.Error().Msg(string(debug.Stack()))
			fmt.Printf("%v.\n", err)
		}
	}()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: out})
	zerolog.SetGlobalLevel(config.DefaultLogLevel)
}
