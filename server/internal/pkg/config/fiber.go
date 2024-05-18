package config

import (
	"strconv"

	"github.com/TesyarRAz/penggerak/internal/pkg/errors"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewFiber(config util.DotEnvConfig) *fiber.App {
	prefork, _ := strconv.ParseBool(config["WEB_PREFORK"])

	app := fiber.New(fiber.Config{
		AppName:      config["APP_NAME"],
		ErrorHandler: NewErrorHandler(),
		Prefork:      prefork,
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		ers := make(map[string]string)
		msg := err.Error()
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		if e, ok := err.(validator.ValidationErrors); ok {
			msg = "Validation error"
			for _, fe := range e {
				ers[fe.Field()] = fe.Tag()
			}
		}
		if _, ok := err.(errors.Unauthorized); ok {
			code = fiber.StatusUnauthorized
			msg = "Unauthorized"
		}
		if _, ok := err.(errors.NotFound); ok {
			code = fiber.StatusNotFound
			msg = "Not found"
		}
		if _, ok := err.(errors.InternalServerError); ok {
			code = fiber.StatusInternalServerError
			msg = "Internal server error"
		}

		return ctx.Status(code).JSON(fiber.Map{
			"message": msg,
			"errors":  ers,
		})
	}
}
