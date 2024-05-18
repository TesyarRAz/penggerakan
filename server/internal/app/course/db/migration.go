package course_migration

import (
	migration "github.com/TesyarRAz/penggerak/internal/pkg/db"

	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
)

func New(config *migration.MigrationConfig) (*migration.Migration, error) {
	return migration.New(config, &pgx.Config{
		MigrationsTable: "course_schema_migrations",
	}, nil)
}
