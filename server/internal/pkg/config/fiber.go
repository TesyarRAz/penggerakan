package config

import (
	"strconv"

	"github.com/TesyarRAz/penggerak/internal/pkg/util"
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
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
