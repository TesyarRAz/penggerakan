package config

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func NewDatabase(config util.DotEnvConfig, log *logrus.Logger) *sqlx.DB {
	idleConnection, _ := strconv.Atoi(config["DB_POOL_IDLE"])
	maxConnection, _ := strconv.Atoi(config["DB_POOL_MAX"])
	maxLifeTimeConnection, _ := strconv.Atoi(config["DB_POOL_LIFETIME"])

	dsn := GenerateDSNFromConfig(config)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.SetMaxIdleConns(idleConnection)
	db.SetMaxOpenConns(maxConnection)
	db.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}

func GenerateDSNFromConfig(config util.DotEnvConfig) string {
	username := config["DB_USERNAME"]
	password := config["DB_PASSWORD"]
	host := config["DB_HOST"]
	port := config["DB_PORT"]
	database := config["DB_NAME"]
	sslMode := config["DB_SSLMODE"]

	q := url.Values{}
	q.Set("sslmode", sslMode)

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", username, password, host, port, database, q.Encode())
}
