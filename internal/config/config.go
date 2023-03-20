package config

import (
	"github.com/caarlos0/env/v6"
	ai "github.com/sashabaranov/go-openai"
)

type ChatGPTConfig struct {
	APIKey           string `env:"CHATGPT_API_KEY,required"`
	MaxResponseToken int    `env:"CHATGPT_MAX_RESPONSE_TOKEN,default=1048576"`
}

type ShellConfig struct {
	You string `env:"SHELL_YOU_PROMPT,default=You"`
	AI  string `env:"SHELL_AI_PROMPT,default=ChatGPT"`
}

type Config struct {
	ApiKey           string `env:"OPENAI_API_KEY,required"`
	MaxResponseToken int    `env:"OPENAI_MAX_RESPONSE_TOKEN"`
	Model            string `env:"OPENAI_CHAT_COMPLETIONS_MODEL"`
}

func NewConfig() Config {
	var cfg Config

	// Set defaults in case the corresponding ENV_VAR is not presented
	config := Config{
		MaxResponseToken: DefaultMaxResponseToken,
		Model:            ai.GPT3Dot5Turbo,
	}

	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	return cfg
}
