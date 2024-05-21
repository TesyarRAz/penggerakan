package config

import (
	"fmt"
	"net/url"

	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func NewDatabase(config model.DotEnvConfig, log *logrus.Logger) *sqlx.DB {
	dsn := GenerateDSNFromConfig(config)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}

func GenerateDSNFromConfig(config model.DotEnvConfig) string {
	username := config.DBUsername()
	password := config.DBPassword()
	host := config.DBHost()
	port := config.DBPort()
	database := config.DBName()
	sslMode := config.DBSSLMode()

	q := url.Values{}
	q.Set("sslmode", sslMode)

	return fmt.Sprintf("postgres://%s:%s@%s:%v/%s?%s", username, password, host, port, database, q.Encode())
}
