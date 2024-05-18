package config

import (
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type BootstrapConfig struct {
	Fiber    *fiber.App
	DB       *sqlx.DB
	Log      *logrus.Logger
	Validate *validator.Validate
	Env      util.DotEnvConfig
}

type App interface {
	Provider() Provider
	Service(Provider)
}

func Bootstrap(app App) {
	providers := app.Provider()
	for _, provider := range providers {
		provider.Boot()
	}

	app.Service(providers)
}
