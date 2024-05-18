package user_middleware

import (
	"strings"

	user_usecase "github.com/TesyarRAz/penggerak/internal/app/user/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func NewAuth(userUserCase *user_usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		if token == "" {
			return fiber.ErrUnauthorized
		}

		splitToken := strings.Split(token, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}

		parsedToken := splitToken[1]

		request := &model.VerifyUserRequest{Token: parsedToken}

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
