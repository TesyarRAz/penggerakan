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

		// Parse the token
		// Example: "Bearer <token>"
		splitToken := strings.Split(token, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}

		parsedToken := splitToken[1]

		// Continue with the rest of the code
		request := &model.VerifyUserRequest{Token: parsedToken}
		userUserCase.Log.Debugf("Authorization : %s", request.Token)

		auth, err := userUserCase.Verify(ctx.UserContext(), request)
		if err != nil {
			userUserCase.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		userUserCase.Log.Debugf("User : %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
