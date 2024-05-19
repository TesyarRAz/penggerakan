package user_http

import (
	"net/http"

	user_model "github.com/TesyarRAz/penggerak/internal/app/user/model"
	user_usecase "github.com/TesyarRAz/penggerak/internal/app/user/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TeacherController struct {
	Log     *logrus.Logger
	UseCase *user_usecase.TeacherUseCase
}

func NewTeacherController(useCase *user_usecase.TeacherUseCase, logger *logrus.Logger) *TeacherController {
	return &TeacherController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *TeacherController) Create(ctx *fiber.Ctx) error {
	var request user_model.CreateTeacherRequest
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(model.WebResponse[*user_model.TeacherResponse]{Data: response})
}

func (c *TeacherController) List(ctx *fiber.Ctx) error {
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

	return ctx.JSON(model.PageResponse[*user_model.TeacherResponse]{Data: response, PageMetadata: *pageInfo})
}

func (c *TeacherController) FindById(ctx *fiber.Ctx) error {
	var request user_model.FindTeacherRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request query : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.FindById(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*user_model.TeacherResponse]{Data: response})
}

func (c *TeacherController) Update(ctx *fiber.Ctx) error {
	var request user_model.UpdateTeacherRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request query : %+v", err)
		return fiber.ErrBadRequest
	}
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Update(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*user_model.TeacherResponse]{Data: response})
}

func (c *TeacherController) Delete(ctx *fiber.Ctx) error {
	var request user_model.DeleteTeacherRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request query : %+v", err)
		return fiber.ErrBadRequest
	}

	if err := c.UseCase.Delete(ctx.UserContext(), &request); err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[any]{
		Message: "Teacher deleted",
	})
}
