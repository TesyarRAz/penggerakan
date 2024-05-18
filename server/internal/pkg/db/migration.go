package migration

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type SeedHandler = func(ctx context.Context, config *MigrationConfig) error

type Migration interface {
	Up(ctx context.Context, withSeed bool) error
	Down() error
	Drop() error
	Seed(ctx context.Context) error
}

type internalMigration struct {
	config  *MigrationConfig
	migrate *migrate.Migrate

	seed SeedHandler

	Migration
}

type MigrationConfig struct {
	Dsn       string
	SourceURL string
	Logger    *logrus.Logger
	DB        *sqlx.DB
}

func New(config *MigrationConfig, pgConfig *pgx.Config, seed SeedHandler) (*Migration, error) {
	driver, err := pgx.WithInstance(config.DB.DB, pgConfig)
	if err != nil {
		config.Logger.Errorf("Error when creating driver: %v", err)
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(config.SourceURL, "postgres", driver)
	if err != nil {
		config.Logger.Errorf("Error when creating migration: %v", err)
		return nil, err
	}

	migration := Migration(&internalMigration{
		config:  config,
		migrate: m,
		seed:    seed,
	})

	return &migration, nil
}

func (m *internalMigration) Up(ctx context.Context, withSeed bool) error {
	if err := m.migrate.Up(); err != nil {
		return err
	}

	if withSeed {
		if err := m.Seed(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (m *internalMigration) Down() error {
	return m.migrate.Down()
}
func (m *internalMigration) Drop() error {
	return m.migrate.Drop()
}
func (m *internalMigration) Seed(ctx context.Context) error {
	if m.seed != nil {
		if err := m.seed(ctx, m.config); err != nil {
			m.config.Logger.Errorf("Error when seeding: %v", err)
			return err
		}
	}

	return nil
}
