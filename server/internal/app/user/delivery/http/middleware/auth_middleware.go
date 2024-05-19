package user_middleware

import (
	user_usecase "github.com/TesyarRAz/penggerak/internal/app/user/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	shared_model "github.com/TesyarRAz/penggerak/internal/pkg/model/shared"
	"github.com/gofiber/fiber/v2"
)

func NewAuth(userUserCase *user_usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request, err := shared_model.NewVerifyUserRequestFromAuthorizationHeader(ctx.Get("Authorization"))
		if err != nil {
			userUserCase.Log.Warnf("Failed parse token : %+v", err)
			return err
		}

		auth, err := userUserCase.Verify(ctx.UserContext(), request)
		if err != nil {
			userUserCase.Log.Warnf("Failed find user by token : %+v", err)
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
