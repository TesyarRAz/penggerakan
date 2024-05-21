package config

import (
	"strconv"

	"github.com/TesyarRAz/penggerak/internal/pkg/errors"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewFiber(config model.DotEnvConfig) *fiber.App {
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
		ers := make(map[string]interface{})
		msg := err.Error()
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		if e, ok := err.(validator.ValidationErrors); ok {
			msg = "Validation error"
			for _, fe := range e {
				ers[fe.Field()] = fiber.Map{
					"tag": fe.Tag(),
					"val": fe.Param(),
				}
			}
			code = fiber.StatusBadRequest
		}
		if _, ok := err.(errors.Unauthorized); ok {
			code = fiber.StatusUnauthorized
		}
		if _, ok := err.(errors.NotFound); ok {
			code = fiber.StatusNotFound
		}
		if _, ok := err.(errors.InternalServerError); ok {
			code = fiber.StatusInternalServerError
		}
		if _, ok := err.(errors.BadRequest); ok {
			code = fiber.StatusBadRequest
		}
		if _, ok := err.(errors.Conflict); ok {
			code = fiber.StatusConflict
		}

		return ctx.Status(code).JSON(fiber.Map{
			"message": msg,
			"errors":  ers,
		})
	}
}
