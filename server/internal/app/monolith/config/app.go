package monolith_config

import (
	course_config "github.com/TesyarRAz/penggerak/internal/app/course/config"
	user_config "github.com/TesyarRAz/penggerak/internal/app/user/config"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type BootstrapConfig struct {
	App      *fiber.App
	DB       *sqlx.DB
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   util.DotEnvConfig
}

func Bootstrap(config *BootstrapConfig) {
	user_config.Bootstrap(&user_config.BootstrapConfig{
		App:      config.App,
		DB:       config.DB,
		Log:      config.Log,
		Validate: config.Validate,
		Config:   config.Config,
	})

	course_config.Bootstrap(&course_config.BootstrapConfig{
		App:      config.App,
		DB:       config.DB,
		Log:      config.Log,
		Validate: config.Validate,
		Config:   config.Config,
	})
}
