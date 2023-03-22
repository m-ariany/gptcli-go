package config

import (
	"github.com/caarlos0/env/v6"
	ai "github.com/sashabaranov/go-openai"
)

type Config struct {
	ApiKey           string `env:"OPENAI_API_KEY,required"`
	MaxResponseToken int    `env:"OPENAI_MAX_RESPONSE_TOKEN"`
	Model            string `env:"OPENAI_CHAT_COMPLETIONS_MODEL"`
}

// TODO: be replaced by viper

func NewConfig() Config {

	// Set defaults in case the corresponding ENV_VAR is not presented
	config := Config{
		MaxResponseToken: DefaultMaxResponseToken,
		Model:            ai.GPT3Dot5Turbo,
	}

	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	return config
}
