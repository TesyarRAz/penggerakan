package course_usecase

import (
	"context"

	course_entity "github.com/TesyarRAz/penggerak/internal/app/course/entity"
	course_model "github.com/TesyarRAz/penggerak/internal/app/course/model"
	course_converter "github.com/TesyarRAz/penggerak/internal/app/course/model/converter"
	course_repository "github.com/TesyarRAz/penggerak/internal/app/course/repository"
	"github.com/TesyarRAz/penggerak/internal/pkg/errors"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
)

type CourseUseCase struct {
	DB               *sqlx.DB
	Config           util.DotEnvConfig
	Log              *logrus.Logger
	Validate         *validator.Validate
	CourseRepository *course_repository.CourseRepository
}

func NewCourseUseCase(db *sqlx.DB, dotenvcfg util.DotEnvConfig, logger *logrus.Logger, validate *validator.Validate, courseRepository *course_repository.CourseRepository) *CourseUseCase {
	return &CourseUseCase{
		DB:               db,
		Config:           dotenvcfg,
		Log:              logger,
		Validate:         validate,
		CourseRepository: courseRepository,
	}
}

func (c *CourseUseCase) List(ctx context.Context, request *model.PageRequest) ([]*course_model.CourseResponse, *model.PageMetadata, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, nil, errors.InternalServerError{}
	}
	defer tx.Rollback()

	var courses []*course_entity.Course
	pageInfo, err := c.CourseRepository.List(tx, &courses, request)
	if err != nil {
		c.Log.Warnf("Failed to list course : %+v", err)
		return nil, nil, errors.InternalServerError{}
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, nil, errors.InternalServerError{}
	}

	return lop.Map(courses, func(course *course_entity.Course, _ int) *course_model.CourseResponse {
		return course_converter.CourseToResponse(course)
	}), pageInfo, nil
}

func (c *CourseUseCase) Create(ctx context.Context, request *course_model.CreateCourseRequest) (*course_model.CourseResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}
	defer tx.Rollback()

	course := course_entity.Course{
		Name:  request.Name,
		Image: request.Image,
	}

	if err = c.CourseRepository.Create(tx, &course); err != nil {
		c.Log.Warnf("Failed to create course : %+v", err)
		return nil, errors.InternalServerError{}
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}

	return course_converter.CourseToResponse(&course), nil
}

func (c *CourseUseCase) Update(ctx context.Context, request *course_model.UpdateCourseRequest) (*course_model.CourseResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}
	defer tx.Rollback()

	var course course_entity.Course
	if err := c.CourseRepository.FindById(tx, &course, request.ID); err != nil {
		c.Log.Warnf("Failed to find course : %+v", err)
		return nil, errors.NotFound{}
	}

	course.Name = request.Name
	course.Image = request.Image

	if err = c.CourseRepository.Update(tx, &course); err != nil {
		c.Log.Warnf("Failed to update course : %+v", err)
		return nil, errors.InternalServerError{}
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}

	return course_converter.CourseToResponse(&course), nil
}

func (c *CourseUseCase) Delete(ctx context.Context, request *course_model.DeleteCourseRequest) error {
	if err := c.Validate.Struct(request); err != nil {
		return err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return errors.InternalServerError{}
	}
	defer tx.Rollback()

	var course course_entity.Course
	if err := c.CourseRepository.FindById(tx, &course, request.ID); err != nil {
		c.Log.Warnf("Failed to find course : %+v", err)
		return errors.NotFound{}
	}

	if err = c.CourseRepository.Delete(tx, &course); err != nil {
		c.Log.Warnf("Failed to delete course : %+v", err)
		return errors.InternalServerError{}
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return errors.InternalServerError{}
	}

	return nil
}

func (c *CourseUseCase) FindById(ctx context.Context, request *course_model.FindCourseRequest) (*course_model.CourseResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}
	defer tx.Rollback()

	var course course_entity.Course
	if err := c.CourseRepository.FindById(tx, &course, request.ID); err != nil {
		c.Log.Warnf("Failed to find course : %+v", err)
		return nil, errors.NotFound{}
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}

	return course_converter.CourseToResponse(&course), nil
}
