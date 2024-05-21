package config

import (
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type BootstrapConfig struct {
	Fiber    *fiber.App
	DB       *sqlx.DB
	Log      *logrus.Logger
	Validate *validator.Validate
	Env      model.DotEnvConfig
	Redis    *redis.Client
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
