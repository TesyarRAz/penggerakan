package user_http

import (
	user_middleware "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http/middleware"
	user_usecase "github.com/TesyarRAz/penggerak/internal/app/user/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *user_usecase.UserUseCase
}

func NewUserController(useCase *user_usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var request model.LoginUserRequest
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Login(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.LoginUserResponse]{Data: response})
}

func (c *UserController) Me(ctx *fiber.Ctx) error {
	auth := user_middleware.GetUser(ctx)

	request := &model.GetUserRequest{
		ID:         auth.ID,
		IsDetailed: true,
	}

	response, err := c.UseCase.GetUser(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Logout(ctx *fiber.Ctx) error {
	return nil
}
