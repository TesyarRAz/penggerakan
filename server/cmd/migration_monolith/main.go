package main

import (
	"context"
	"flag"

	monolith_migration "github.com/TesyarRAz/penggerak/internal/app/monolith/db"
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
	defer db.Close()

	flag.Parse()

	migration, err := monolith_migration.New(&monolith_migration.MigrationConfig{
		UserSourceURL:   "file://internal/app/user/db/migrations",
		CourseSourceURL: "file://internal/app/course/db/migrations",
		Logger:          logger,
		DB:              db,
	})
	if err != nil {
		logger.Fatalf("Failed to create migration: %v", err)
		return
	}

	if err := (*migration).Up(ctx, *withSeed); err != nil {
		logger.Fatalf("Failed to migrate user: %v", err)
		return
	}

	logger.Info("Migration done")
}
