package config

import (
	"strconv"

	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/sirupsen/logrus"
)

func NewLogger(config util.DotEnvConfig) *logrus.Logger {
	log := logrus.New()

	logLevel, _ := strconv.Atoi(config["LOG_LEVEL"])

	log.SetLevel(logrus.Level(logLevel))
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)

	return log
}
