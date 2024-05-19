package course_middleware

import (
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	shared_model "github.com/TesyarRAz/penggerak/internal/pkg/model/shared"
	"github.com/TesyarRAz/penggerak/internal/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func NewAuth(authService *service.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request, err := shared_model.NewVerifyUserRequestFromAuthorizationHeader(ctx.Get("Authorization"))
		if err != nil {
			return err
		}

		auth, err := (*authService).Verify(ctx.UserContext(), request)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		if err := auth.ParseRoleAndPermission(); err != nil {
			return fiber.ErrUnauthorized
		}

		ctx.Locals("auth", auth)

		return ctx.Next()
	}
}

func GetAuth(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
