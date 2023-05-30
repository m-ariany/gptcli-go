package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type ChatGPTConfig struct {
	APIKey    string `env:"OPENAI_API_KEY,required"`
	Model     string `env:"OPENAI_MODEL,default=gpt-3.5-turbo"`
	MaxTokens int    `env:"OPENAI_MAX_TOKENS,default=1024"`
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
