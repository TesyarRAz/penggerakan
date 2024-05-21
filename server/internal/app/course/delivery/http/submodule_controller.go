package course_http

import (
	course_model "github.com/TesyarRAz/penggerak/internal/app/course/model"
	course_usecase "github.com/TesyarRAz/penggerak/internal/app/course/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SubModuleController struct {
	Log     *logrus.Logger
	UseCase *course_usecase.SubModuleUseCase
}

func NewSubModuleController(useCase *course_usecase.SubModuleUseCase, log *logrus.Logger) *SubModuleController {
	return &SubModuleController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *SubModuleController) List(ctx *fiber.Ctx) error {
	var request course_model.ListSubModuleRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}
	if err := ctx.QueryParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}

	request.GenerateDefault()

	response, pageInfo, err := c.UseCase.List(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.PageResponse[*course_model.SubModuleResponse]{Data: response, PageMetadata: *pageInfo})
}

func (c *SubModuleController) Create(ctx *fiber.Ctx) error {
	var request course_model.CreateSubModuleRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubModuleController) Update(ctx *fiber.Ctx) error {
	var request course_model.UpdateSubModuleRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
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

	return ctx.JSON(response)
}

func (c *SubModuleController) Delete(ctx *fiber.Ctx) error {
	var request course_model.DeleteSubModuleRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}

	err := c.UseCase.Delete(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse{
		Message: "SubModule deleted successfully",
	})
}

func (c *SubModuleController) FindById(ctx *fiber.Ctx) error {
	var request course_model.FindSubModuleRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.FindById(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}
