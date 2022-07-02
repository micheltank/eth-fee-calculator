package config

import (
	"github.com/Netflix/go-env"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type Config struct {
	LogLevel       string `env:"LOG_LEVEL"`
	Port           int    `env:"PORT" validate:"required"`
	SoChainApiHost string `env:"SO_CHAIN_API_HOST"`
	DbConfig       DbConfig
	Extras         env.EnvSet
}

func NewConfig() (config Config, err error) {
	es, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		return Config{}, errors.Wrap(err, "failed to initialize config: %v")
	}
	err = validator.New().Struct(config)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		return Config{}, errors.Wrap(validationErrors, "failed on config validation")
	}

	// Remaining environment variables.
	config.Extras = es

	return config, nil
}
