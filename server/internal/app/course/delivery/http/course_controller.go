package course_http

import (
	"net/http"

	course_usecase "github.com/TesyarRAz/penggerak/internal/app/course/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CourseController struct {
	Log     *logrus.Logger
	UseCase *course_usecase.CourseUseCase
}

func NewCourseController(useCase *course_usecase.CourseUseCase, log *logrus.Logger) *CourseController {
	return &CourseController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *CourseController) List(ctx *fiber.Ctx) error {
	var request model.PageRequest
	if err := ctx.QueryParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}
	if request.Order == "" {
		request.Order = "desc"
	}
	if request.PerPage == 0 {
		request.PerPage = 10
	}
	response, pageInfo, err := c.UseCase.List(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.PageResponse[*model.CourseResponse]{Data: response, PageMetadata: *pageInfo})
}

func (c *CourseController) Create(ctx *fiber.Ctx) error {
	var request model.CreateCourseRequest
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(model.WebResponse[*model.CourseResponse]{Data: response})
}

func (c *CourseController) Update(ctx *fiber.Ctx) error {
	var request model.UpdateCourseRequest
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("id")

	response, err := c.UseCase.Update(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CourseResponse]{Data: response})
}

func (c *CourseController) Delete(ctx *fiber.Ctx) error {
	request := model.DeleteCourseRequest{
		ID: ctx.Params("id"),
	}

	err := c.UseCase.Delete(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[any]{
		Message: "Course deleted successfully",
	})
}

func (c *CourseController) FindById(ctx *fiber.Ctx) error {
	request := model.FindCourseRequest{
		ID: ctx.Params("id"),
	}

	response, err := c.UseCase.FindById(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CourseResponse]{Data: response})
}
