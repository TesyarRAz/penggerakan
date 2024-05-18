package user_middleware

import (
	user_usecase "github.com/TesyarRAz/penggerak/internal/app/user/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func NewAuth(userUserCase *user_usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request, err := model.NewVerifyUserRequestFromAuthorizationHeader(ctx.Get("Authorization"))
		if err != nil {
			userUserCase.Log.Warnf("Failed parse token : %+v", err)
			return err
		}

		auth, err := userUserCase.Verify(ctx.UserContext(), request)
		if err != nil {
			userUserCase.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
