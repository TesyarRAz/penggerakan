package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func NewDatabase(config util.DotEnvConfig, log *logrus.Logger) *sqlx.DB {
	idleConnection, _ := strconv.Atoi(config["DB_POOL_IDLE"])
	maxConnection, _ := strconv.Atoi(config["DB_POOL_MAX"])
	maxLifeTimeConnection, _ := strconv.Atoi(config["DB_POOL_LIFETIME"])

	dsn := GenerateDSNFromConfig(config)

	db, err := sqlx.Connect("postgres", dsn)
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

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, database, sslMode)
}
