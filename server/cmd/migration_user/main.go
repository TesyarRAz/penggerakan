package main

import (
	"context"
	"flag"

	user_migration "github.com/TesyarRAz/penggerak/internal/app/user/db"
	"github.com/TesyarRAz/penggerak/internal/pkg/config"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var (
	dotenv util.DotEnvConfig
	logger *logrus.Logger
	db     *sqlx.DB

	withSeed = flag.Bool("seed", false, "Seed database")
)

func main() {
	ctx := context.Background()
	dotenv = config.NewDotEnv()
	logger = config.NewLogger(dotenv)
	db = config.NewDatabase(dotenv, logger)

	flag.Parse()

	migration, err := user_migration.New(&user_migration.MigrationConfig{
		Dsn:       config.GenerateDSNFromConfig(dotenv),
		SourceURL: "file://internal/app/user/db/migration",
		Logger:    logger,
		DB:        db,
	})

	if err != nil {
		logger.Fatalf("Failed to create migration: %v", err)
		return
	}

	migration.Up(ctx, *withSeed)

	logger.Info("Migration done")
}
