package config

import (
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/sirupsen/logrus"
)

func NewLogger(config model.DotEnvConfig) *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(config.LogLevel()))
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)

	return log
}
