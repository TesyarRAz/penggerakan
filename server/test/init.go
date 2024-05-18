package test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	monolith_config "github.com/TesyarRAz/penggerak/internal/app/monolith/config"
	monolith_migration "github.com/TesyarRAz/penggerak/internal/app/monolith/db"
	"github.com/TesyarRAz/penggerak/internal/pkg/config"
	migration "github.com/TesyarRAz/penggerak/internal/pkg/db"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	app       *fiber.App
	db        *sqlx.DB
	dotenvcfg util.DotEnvConfig
	log       *logrus.Logger
	validate  *validator.Validate

	m *migration.Migration
)

func init() {
	ctx := context.Background()
	dotenvcfg = config.NewDotEnv("../.env.test")
	log = config.NewLogger(dotenvcfg)
	validate = config.NewValidator()
	app = config.NewFiber(dotenvcfg)
	db = config.NewDatabase(dotenvcfg, log)

	monolith_config.Bootstrap(&monolith_config.BootstrapConfig{
		App:      app,
		DB:       db,
		Config:   dotenvcfg,
		Log:      log,
		Validate: validate,
	})

	var err error
	m, err = monolith_migration.New(&monolith_migration.MigrationConfig{
		UserSourceURL:   "file://../internal/app/user/db/migrations",
		CourseSourceURL: "file://../internal/app/course/db/migrations",
		Logger:          log,
		DB:              db,
	})
	if err != nil {
		log.Fatalf("Failed to create migration: %v", err)
		return
	}
	(*m).Down()
	if err = (*m).Up(ctx, true); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
		return
	}
}

func GetAdmin(t *testing.T) (*model.LoginUserRequest, *http.Response, *model.WebResponse[model.LoginUserResponse]) {
	requestBody := model.LoginUserRequest{
		Email:    "admin@example.com",
		Password: "password",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.LoginUserResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	return &requestBody, response, responseBody
}

func GetAdminToken(t *testing.T) string {
	_, _, responseBody := GetAdmin(t)

	token := responseBody.Data.Token
	assert.NotNil(t, token)

	return token
}
