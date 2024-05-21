package course_http

import (
	course_model "github.com/TesyarRAz/penggerak/internal/app/course/model"
	course_usecase "github.com/TesyarRAz/penggerak/internal/app/course/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ModuleController struct {
	Log     *logrus.Logger
	UseCase *course_usecase.ModuleUseCase
}

func NewModuleController(useCase *course_usecase.ModuleUseCase, log *logrus.Logger) *ModuleController {
	return &ModuleController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *ModuleController) List(ctx *fiber.Ctx) error {
	var request course_model.ListModuleRequest
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

	return ctx.JSON(model.PageResponse[*course_model.ModuleResponse]{Data: response, PageMetadata: *pageInfo})
}

func (c *ModuleController) Create(ctx *fiber.Ctx) error {
	var request course_model.CreateModuleRequest
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

func (c *ModuleController) Update(ctx *fiber.Ctx) error {
	var request course_model.UpdateModuleRequest
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

func (c *ModuleController) Delete(ctx *fiber.Ctx) error {
	var request course_model.DeleteModuleRequest
	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}

	err := c.UseCase.Delete(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse{
		Message: "Module deleted successfully",
	})
}

func (c *ModuleController) FindById(ctx *fiber.Ctx) error {
	var request course_model.FindModuleRequest
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
