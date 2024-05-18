package course_usecase

import (
	"context"

	course_entity "github.com/TesyarRAz/penggerak/internal/app/course/entity"
	course_converter "github.com/TesyarRAz/penggerak/internal/app/course/model/converter"
	course_repository "github.com/TesyarRAz/penggerak/internal/app/course/repository"
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

func (c *CourseUseCase) List(ctx context.Context, request *model.PageRequest) ([]*model.CourseResponse, *model.PageMetadata, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, nil, err
	}
	defer tx.Rollback()

	var courses []*course_entity.Course
	pageInfo, err := c.CourseRepository.List(tx, &courses, request)
	if err != nil {
		c.Log.Warnf("Failed to list course : %+v", err)
		return nil, nil, err
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, nil, err
	}

	return lop.Map(courses, func(course *course_entity.Course, _ int) *model.CourseResponse {
		return course_converter.CourseToResponse(course)
	}), pageInfo, nil
}

func (c *CourseUseCase) Create(ctx context.Context, request *model.CreateCourseRequest) (*model.CourseResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, err
	}
	defer tx.Rollback()

	course := course_entity.Course{
		Name:  request.Name,
		Image: request.Image,
	}

	if err = c.CourseRepository.Create(tx, &course); err != nil {
		c.Log.Warnf("Failed to create course : %+v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, err
	}

	return course_converter.CourseToResponse(&course), nil
}
