package config

import (
	"fmt"

	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/joho/godotenv"
)

func NewDotEnv(filenames ...string) util.DotEnvConfig {
	var config util.DotEnvConfig

	config, err := godotenv.Read(filenames...)

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return config
}
