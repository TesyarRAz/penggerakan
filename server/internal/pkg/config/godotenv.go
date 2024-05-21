package config

import (
	"fmt"

	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/joho/godotenv"
)

func NewDotEnv(filenames ...string) model.DotEnvConfig {
	var config model.DotEnvConfig

	config, err := godotenv.Read(filenames...)

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return config
}
