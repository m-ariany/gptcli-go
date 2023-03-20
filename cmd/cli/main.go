package main

import (
	"fmt"
	"runtime/debug"

	"github.com/m-ariany/gptcli/internal/assistant"
	"github.com/m-ariany/gptcli/internal/config"
	"github.com/m-ariany/gptcli/internal/delivery/cli"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func main() {
	cfg := config.NewConfig()
	_ = cfg

	shell := cli.New()
	shell.Run()
}

func run(cmd *cobra.Command, args []string) {
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

	config := loadConfiguration()
	assistant := assistant.NewAssistant(config)
	assistant.Run()
}

func loadConfiguration() config.Config {
	return config.NewConfig()
}
