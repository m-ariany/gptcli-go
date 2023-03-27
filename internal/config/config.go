package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type ChatGPTConfig struct {
	APIKey    string `env:"CHATGPT_API_KEY,required"`
	Model     string `env:"CHATGPT_MODEL,default=gpt-3.5-turbo"`
	MaxTokens int    `env:"CHATGPT_MAX_TOKENS,default=0"`
}

type ShellConfig struct {
	You string `env:"SHELL_YOU_PROMPT,default=You"`
	AI  string `env:"SHELL_AI_PROMPT,default=ChatGPT"`
}

type Config struct {
	ChatGPT ChatGPTConfig
	Shell   ShellConfig
}

func NewConfig() Config {
	var cfg Config

	ctx := context.Background()
	err := envconfig.Process(ctx, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
