package test

import (
	"context"

	user_config "github.com/TesyarRAz/penggerak/internal/app/user/config"
	user_migration "github.com/TesyarRAz/penggerak/internal/app/user/db"
	"github.com/TesyarRAz/penggerak/internal/pkg/config"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var (
	app       *fiber.App
	db        *sqlx.DB
	dotenvcfg util.DotEnvConfig
	log       *logrus.Logger
	validate  *validator.Validate
)

func init() {
	ctx := context.Background()
	dotenvcfg = config.NewDotEnv("../../.env.test")
	log = config.NewLogger(dotenvcfg)
	validate = config.NewValidator()
	app = config.NewFiber(dotenvcfg)
	db = config.NewDatabase(dotenvcfg, log)

	user_config.Bootstrap(&user_config.BootstrapConfig{
		App:      app,
		DB:       db,
		Config:   dotenvcfg,
		Log:      log,
		Validate: validate,
	})

	migration, err := user_migration.New(&user_migration.MigrationConfig{
		Dsn:       config.GenerateDSNFromConfig(dotenvcfg),
		SourceURL: "file://../../internal/app/user/db/migrations",
		Logger:    log,
		DB:        db,
	})
	if err != nil {
		log.Fatalf("Failed to create migration: %v", err)
		return
	}

	migration.Down(0)
	if err := migration.Up(ctx, true); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
		return
	}
}
