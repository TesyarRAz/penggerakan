package user_http

import (
	user_middleware "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http/middleware"
	user_policy "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http/policy"
	user_model "github.com/TesyarRAz/penggerak/internal/app/user/model"
	user_usecase "github.com/TesyarRAz/penggerak/internal/app/user/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	shared_model "github.com/TesyarRAz/penggerak/internal/pkg/model/shared"
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
	var request user_model.LoginUserRequest
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Login(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *UserController) Me(ctx *fiber.Ctx) error {
	auth := user_middleware.GetAuth(ctx)

	request := &user_model.FindUserRequest{
		ID:         auth.ID,
		IsDetailed: user_policy.AllowDetailedUser(auth, auth.ID),
	}

	response, err := c.UseCase.FindUserById(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *UserController) Logout(ctx *fiber.Ctx) error {
	return nil
}

func (c *UserController) List(ctx *fiber.Ctx) error {
	var request model.PageRequest
	if err := ctx.QueryParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request query : %+v", err)
		return fiber.ErrBadRequest
	}

	request.GenerateDefault()

	response, pageInfo, err := c.UseCase.List(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.PageResponse[*user_model.UserResponse]{Data: response, PageMetadata: *pageInfo})
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	var request user_model.CreateUserRequest
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	auth := user_middleware.GetAuth(ctx)

	var request user_model.UpdateUserRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	if !user_policy.AllowUpdateUser(auth, request.ID) {
		return fiber.ErrUnauthorized
	}

	response, err := c.UseCase.Update(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *UserController) Delete(ctx *fiber.Ctx) error {
	auth := user_middleware.GetAuth(ctx)
	var request user_model.DeleteUserRequest

	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}

	if !user_policy.AllowDeleteUser(auth, request.ID) {
		return fiber.ErrUnauthorized
	}

	err := c.UseCase.Delete(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse{
		Message: "User deleted successfully",
	})
}

func (c *UserController) FindById(ctx *fiber.Ctx) error {
	auth := user_middleware.GetAuth(ctx)

	var request user_model.FindUserRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}

	request.IsDetailed = user_policy.AllowDetailedUser(auth, request.ID)

	response, err := c.UseCase.FindUserById(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *UserController) AttachRole(ctx *fiber.Ctx) error {
	var request user_model.AttachRoleToUserRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	err := c.UseCase.AttachRole(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse{
		Message: "Role attached successfully",
	})
}

func (c *UserController) DetachRole(ctx *fiber.Ctx) error {
	var request user_model.DetachRoleFromUserRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	err := c.UseCase.DetachRole(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse{
		Message: "Role detached successfully",
	})
}

func (c *UserController) RefreshToken(ctx *fiber.Ctx) error {
	var request shared_model.RefreshTokenRequest
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.RefreshToken(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}
