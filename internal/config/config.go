package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ApiKey   string `env:"OPENAI_API_KEY,required"`
	MaxToken int    `env:"OPENAI_MAX_TOKEN"`
}

// TODO: be replaced by viper

func NewConfig() Config {

	config := Config{
		MaxToken: DefaultMaxToken,
	}

	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	return config
}
