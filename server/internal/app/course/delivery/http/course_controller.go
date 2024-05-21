package course_http

import (
	"net/http"

	course_model "github.com/TesyarRAz/penggerak/internal/app/course/model"
	course_usecase "github.com/TesyarRAz/penggerak/internal/app/course/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	shared_model "github.com/TesyarRAz/penggerak/internal/pkg/model/shared"
	"github.com/TesyarRAz/penggerak/internal/pkg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

type CourseController struct {
	Log            *logrus.Logger
	UseCase        *course_usecase.CourseUseCase
	TeacherService *service.TeacherService
}

func NewCourseController(useCase *course_usecase.CourseUseCase, log *logrus.Logger, teacherService *service.TeacherService) *CourseController {
	return &CourseController{
		Log:            log,
		UseCase:        useCase,
		TeacherService: teacherService,
	}
}

func (c *CourseController) List(ctx *fiber.Ctx) error {
	var request model.PageRequest
	if err := ctx.QueryParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}

	request.GenerateDefault()

	response, pageInfo, err := c.UseCase.List(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	// Add Teacher to Course
	if len(response) > 0 {
		teacherIds := lop.Map(response, func(course *course_model.CourseResponse, _ int) string {
			return course.TeacherID
		})
		teachers, err := (*c.TeacherService).FindByIds(ctx.UserContext(), teacherIds...)
		if err != nil {
			c.Log.Warnf("Failed to find teachers : %+v", err)
			return fiber.ErrInternalServerError
		}
		response = lop.Map(response, func(course *course_model.CourseResponse, _ int) *course_model.CourseResponse {
			teacher, _ := lo.Find(teachers, func(teacher *shared_model.TeacherResponse) bool {
				return teacher.ID == course.TeacherID
			})
			course.Teacher = teacher

			return course
		})
	}

	return ctx.JSON(model.PageResponse[*course_model.CourseResponse]{Data: response, PageMetadata: *pageInfo})
}

func (c *CourseController) Create(ctx *fiber.Ctx) error {
	var request course_model.CreateCourseRequest
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(response)
}

func (c *CourseController) Update(ctx *fiber.Ctx) error {
	var request course_model.UpdateCourseRequest
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

func (c *CourseController) Delete(ctx *fiber.Ctx) error {
	var request course_model.DeleteCourseRequest

	if err := ctx.ParamsParser(&request); err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}

	err := c.UseCase.Delete(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse{
		Message: "Course deleted successfully",
	})
}

func (c *CourseController) FindById(ctx *fiber.Ctx) error {
	var request course_model.FindCourseRequest
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
