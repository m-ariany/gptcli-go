package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	ApiKey           string `env:"OPENAI_API_KEY,required"`
	MaxResponseToken int    `env:"OPENAI_MAX_RESPONSE_TOKEN,default=1048576"`
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
