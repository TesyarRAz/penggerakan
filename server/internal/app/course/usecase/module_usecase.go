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

type ModuleUseCase struct {
	DB               *sqlx.DB
	Config           util.DotEnvConfig
	Log              *logrus.Logger
	Validate         *validator.Validate
	CourseRepository *course_repository.CourseRepository
	ModuleRepository *course_repository.ModuleRepository
}

func NewModuleUseCase(db *sqlx.DB, dotenvcfg util.DotEnvConfig, logger *logrus.Logger, validate *validator.Validate, courseRepository *course_repository.CourseRepository, moduleRepository *course_repository.ModuleRepository) *ModuleUseCase {
	return &ModuleUseCase{
		DB:               db,
		Config:           dotenvcfg,
		Log:              logger,
		Validate:         validate,
		CourseRepository: courseRepository,
		ModuleRepository: moduleRepository,
	}
}

func (c *ModuleUseCase) List(ctx context.Context, request *course_model.ListModuleRequest) ([]*course_model.ModuleResponse, *model.PageMetadata, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, nil, errors.InternalServerError{}
	}
	defer tx.Rollback()

	var modules []*course_entity.Module
	pageInfo, err := c.ModuleRepository.List(tx, &modules, request)
	if err != nil {
		c.Log.Warnf("Failed to list module : %+v", err)
		return nil, nil, errors.InternalServerError{}
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, nil, errors.InternalServerError{}
	}

	return lop.Map(modules, func(module *course_entity.Module, _ int) *course_model.ModuleResponse {
		return course_converter.ModuleToResponse(module)
	}), pageInfo, nil
}

func (c *ModuleUseCase) Create(ctx context.Context, request *course_model.CreateModuleRequest) (*course_model.ModuleResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}
	defer tx.Rollback()

	module := &course_entity.Module{
		CourseID: request.CourseID,
		Name:     request.Name,
	}

	if err := c.CourseRepository.FindById(tx, &course_entity.Course{}, module.CourseID); err != nil {
		c.Log.Warnf("Failed to find course by id : %+v", err)
		return nil, errors.NotFound{}
	}
	if err := c.ModuleRepository.Create(tx, module); err != nil {
		c.Log.Warnf("Failed to create module : %+v", err)
		return nil, errors.InternalServerError{}
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}

	return course_converter.ModuleToResponse(module), nil
}

func (c *ModuleUseCase) Update(ctx context.Context, request *course_model.UpdateModuleRequest) (*course_model.ModuleResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}
	defer tx.Rollback()

	var module course_entity.Module
	if err := c.ModuleRepository.FindById(tx, &module, request.ID); err != nil {
		c.Log.Warnf("Failed to find module : %+v", err)
		return nil, errors.NotFound{}
	}

	module.Name = request.Name
	if err := c.ModuleRepository.Update(tx, &module); err != nil {
		c.Log.Warnf("Failed to update module : %+v", err)
		return nil, errors.InternalServerError{}
	}
	if module.CourseID != request.CourseID {
		c.Log.Warnf("Module not found in course : %+v", err)
		return nil, errors.NotFound{}
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}

	return course_converter.ModuleToResponse(&module), nil
}

func (c *ModuleUseCase) Delete(ctx context.Context, request *course_model.DeleteModuleRequest) error {
	if err := c.Validate.Struct(request); err != nil {
		return err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return errors.InternalServerError{}
	}
	defer tx.Rollback()

	var module course_entity.Module
	if err := c.ModuleRepository.FindById(tx, &module, request.ID); err != nil {
		c.Log.Warnf("Failed to find module : %+v", err)
		return errors.NotFound{}
	}

	if module.CourseID != request.CourseID {
		c.Log.Warnf("Module not found in course : %+v", err)
		return errors.NotFound{}
	}

	if err := c.ModuleRepository.Delete(tx, &module); err != nil {
		c.Log.Warnf("Failed to delete module : %+v", err)
		return errors.InternalServerError{}
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return errors.InternalServerError{}
	}

	return nil
}

func (c *ModuleUseCase) FindById(ctx context.Context, request *course_model.FindModuleRequest) (*course_model.ModuleResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}
	defer tx.Rollback()

	var module course_entity.Module
	if err := c.ModuleRepository.FindById(tx, &module, request.ID); err != nil {
		c.Log.Warnf("Failed to find module : %+v", err)
		return nil, errors.NotFound{}
	}

	if module.CourseID != request.CourseID {
		c.Log.Warnf("Module not found in course : %+v", err)
		return nil, errors.NotFound{}
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}

	return course_converter.ModuleToResponse(&module), nil
}
