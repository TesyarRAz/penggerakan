package monolith_migration

import (
	"context"

	course_migration "github.com/TesyarRAz/penggerak/internal/app/course/db"
	user_migration "github.com/TesyarRAz/penggerak/internal/app/user/db"
	migration "github.com/TesyarRAz/penggerak/internal/pkg/db"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type MigrationConfig struct {
	UserSourceURL   string
	CourseSourceURL string

	Logger *logrus.Logger
	DB     *sqlx.DB
}

type MonolithMigration struct {
	config     *MigrationConfig
	migrations []*migration.Migration

	migration.Migration
}

func New(config *MigrationConfig) (*migration.Migration, error) {
	userMigration, err := user_migration.New(&migration.MigrationConfig{
		Logger:    config.Logger,
		DB:        config.DB,
		SourceURL: config.UserSourceURL,
	})
	if err != nil {
		return nil, err
	}
	courseMigration, err := course_migration.New(&migration.MigrationConfig{
		Logger:    config.Logger,
		DB:        config.DB,
		SourceURL: config.CourseSourceURL,
	})
	if err != nil {
		return nil, err
	}

	migration := migration.Migration(&MonolithMigration{
		config: config,
		migrations: []*migration.Migration{
			userMigration,
			courseMigration,
		},
	})

	return &migration, nil
}

func (u *MonolithMigration) Up(ctx context.Context, withSeed bool) error {
	for _, m := range u.migrations {
		if err := (*m).Up(ctx, withSeed); err != nil {
			return err
		}
	}

	return nil
}

func (u *MonolithMigration) Down() error {
	for _, m := range u.migrations {
		if err := (*m).Down(); err != nil {
			return err
		}
	}

	return nil
}

func (u *MonolithMigration) Drop() error {
	for _, m := range u.migrations {
		if err := (*m).Drop(); err != nil {
			return err
		}
	}

	return nil
}
