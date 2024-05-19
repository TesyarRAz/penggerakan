package course_middleware

import "github.com/gofiber/fiber/v2"

func Role(role ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := GetAuth(ctx)

		if !auth.HasRole(role...) {
			return fiber.ErrForbidden
		}

		return ctx.Next()
	}
}

func Permission(permission ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := GetAuth(ctx)

		if !auth.HasPermission(permission...) {
			return fiber.ErrForbidden
		}

		return ctx.Next()
	}
}
