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
	"github.com/jmoiron/sqlx/types"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
)

type SubModuleUseCase struct {
	DB                  *sqlx.DB
	Config              model.DotEnvConfig
	Log                 *logrus.Logger
	Validate            *validator.Validate
	ModuleRepository    *course_repository.ModuleRepository
	SubModuleRepository *course_repository.SubModuleRepository
}

func NewSubModuleUseCase(db *sqlx.DB, dotenvcfg model.DotEnvConfig, logger *logrus.Logger, validate *validator.Validate, moduleRepository *course_repository.ModuleRepository, subModuleRepository *course_repository.SubModuleRepository) *SubModuleUseCase {
	return &SubModuleUseCase{
		DB:                  db,
		Config:              dotenvcfg,
		Log:                 logger,
		Validate:            validate,
		ModuleRepository:    moduleRepository,
		SubModuleRepository: subModuleRepository,
	}
}

func (c *SubModuleUseCase) List(ctx context.Context, request *course_model.ListSubModuleRequest) ([]*course_model.SubModuleResponse, *model.PageMetadata, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var submodules []*course_entity.SubModule
	pageInfo, err := c.SubModuleRepository.List(tx, &submodules, request)
	if err != nil {
		c.Log.Warnf("Failed to list module : %+v", err)
		return nil, nil, errors.NewInternalServerError()
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, nil, errors.NewInternalServerError()
	}

	return lop.Map(submodules, func(submodule *course_entity.SubModule, _ int) *course_model.SubModuleResponse {
		return course_converter.SubModuleToResponse(submodule)
	}), pageInfo, nil
}

func (c *SubModuleUseCase) Create(ctx context.Context, request *course_model.CreateSubModuleRequest) (*course_model.SubModuleResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	structure, err := util.TrimJsonRawMessage(request.Structure)
	if err != nil {
		c.Log.Warnf("Failed to trim json raw message : %+v", err)
		return nil, errors.BadRequest{Message: "Invalid json structure"}
	}

	module := &course_entity.SubModule{
		ModuleID:  request.ModuleID,
		Name:      request.Name,
		Structure: types.JSONText(structure),
	}

	if err := c.ModuleRepository.FindById(tx, &course_entity.Module{}, module.ModuleID); err != nil {
		c.Log.Warnf("Failed to find module by id : %+v", err)
		return nil, errors.NewNotFound()
	}
	if err := c.SubModuleRepository.Create(tx, module); err != nil {
		c.Log.Warnf("Failed to create submodule : %+v", err)
		return nil, errors.NewInternalServerError()
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	return course_converter.SubModuleToResponse(module), nil
}

func (c *SubModuleUseCase) Update(ctx context.Context, request *course_model.UpdateSubModuleRequest) (*course_model.SubModuleResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var entity course_entity.SubModule
	if err := c.SubModuleRepository.FindById(tx, &entity, request.ID); err != nil {
		c.Log.Warnf("Failed to find submodule : %+v", err)
		return nil, errors.NewNotFound()
	}

	if entity.ModuleID != request.ModuleID {
		c.Log.Warnf("Submodule not found in module : %+v", err)
		return nil, errors.NewNotFound()
	}

	entity.Name = request.Name
	if err := c.SubModuleRepository.Update(tx, &entity); err != nil {
		c.Log.Warnf("Failed to update submodule : %+v", err)
		return nil, errors.NewInternalServerError()
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	return course_converter.SubModuleToResponse(&entity), nil
}

func (c *SubModuleUseCase) Delete(ctx context.Context, request *course_model.DeleteSubModuleRequest) error {
	if err := c.Validate.Struct(request); err != nil {
		return err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var entity course_entity.SubModule
	if err := c.SubModuleRepository.FindById(tx, &entity, request.ID); err != nil {
		c.Log.Warnf("Failed to find submodule : %+v", err)
		return errors.NewNotFound()
	}

	if entity.ModuleID != request.ModuleID {
		c.Log.Warnf("Submodule not found in module : %+v", err)
		return errors.NewNotFound()
	}

	if err := c.SubModuleRepository.Delete(tx, &entity); err != nil {
		c.Log.Warnf("Failed to delete submodule : %+v", err)
		return errors.NewInternalServerError()
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return errors.NewInternalServerError()
	}

	return nil
}

func (c *SubModuleUseCase) FindById(ctx context.Context, request *course_model.FindSubModuleRequest) (*course_model.SubModuleResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		return nil, err
	}

	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var entity course_entity.SubModule
	if err := c.SubModuleRepository.FindById(tx, &entity, request.ID); err != nil {
		c.Log.Warnf("Failed to find submodule : %+v", err)
		return nil, errors.NewNotFound()
	}

	if entity.ModuleID != request.ModuleID {
		c.Log.Warnf("Submodule not found in module : %+v", err)
		return nil, errors.NewNotFound()
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	return course_converter.SubModuleToResponse(&entity), nil
}
