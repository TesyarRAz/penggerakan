package user_usecase

import (
	"context"
	"database/sql"

	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	user_model "github.com/TesyarRAz/penggerak/internal/app/user/model"
	user_converter "github.com/TesyarRAz/penggerak/internal/app/user/model/converter"
	user_repository "github.com/TesyarRAz/penggerak/internal/app/user/repository"
	"github.com/TesyarRAz/penggerak/internal/pkg/errors"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	lop "github.com/samber/lo/parallel"
)

type TeacherUseCase struct {
	DB       *sqlx.DB
	Config   model.DotEnvConfig
	Log      *logrus.Logger
	Validate *validator.Validate

	UserRepository    *user_repository.UserRepository
	TeacherRepository *user_repository.TeacherRepository
}

func NewTeacherUseCase(db *sqlx.DB, dotenvcfg model.DotEnvConfig, logger *logrus.Logger, validate *validator.Validate,
	userRepository *user_repository.UserRepository, teacherRepository *user_repository.TeacherRepository) *TeacherUseCase {
	return &TeacherUseCase{
		DB:                db,
		Config:            dotenvcfg,
		Log:               logger,
		Validate:          validate,
		UserRepository:    userRepository,
		TeacherRepository: teacherRepository,
	}
}

func (c *TeacherUseCase) Create(ctx context.Context, request *user_model.CreateTeacherRequest) (*user_model.TeacherResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var user user_entity.User
	if err := c.UserRepository.FindById(tx, &user, request.UserID); err != nil {
		c.Log.Warnf("Failed to find user : %+v", err)
		return nil, errors.NewNotFound()
	}

	// Check if user has been registered as teacher
	var teacher user_entity.Teacher
	if err := c.TeacherRepository.FindByUserId(tx, &teacher, user.ID); err == nil {
		c.Log.Warnf("User has been registered as teacher")
		return nil, errors.Conflict{
			Message: "User has been registered as teacher",
		}
	}

	teacher = user_entity.Teacher{
		UserID: user.ID,
		Name:   request.Name,
	}
	if request.ProfileImage != "" {
		teacher.ProfileImage = sql.NullString{
			String: request.ProfileImage,
			Valid:  true,
		}
	}

	if err := c.TeacherRepository.Create(tx, &teacher); err != nil {
		c.Log.Warnf("Failed to create teacher : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	return user_converter.TeacherToResponse(&teacher), nil
}

func (c *TeacherUseCase) List(ctx context.Context, request *model.PageRequest) ([]*user_model.TeacherResponse, *model.PageMetadata, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, nil, err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var teachers []*user_entity.Teacher
	pageInfo, err := c.TeacherRepository.List(tx, &teachers, request)
	if err != nil {
		c.Log.Warnf("Failed to list teacher : %+v", err)
		return nil, nil, errors.NewInternalServerError()
	}

	return lop.Map(teachers, func(teacher *user_entity.Teacher, _ int) *user_model.TeacherResponse {
		return user_converter.TeacherToResponse(teacher)
	}), pageInfo, nil
}

func (c *TeacherUseCase) Update(ctx context.Context, request *user_model.UpdateTeacherRequest) (*user_model.TeacherResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var teacher user_entity.Teacher
	if err := c.TeacherRepository.FindById(tx, &teacher, request.ID); err != nil {
		c.Log.Warnf("Failed to find teacher : %+v", err)
		return nil, errors.NewNotFound()
	}

	teacher.Name = request.Name
	if request.ProfileImage != "" {
		teacher.ProfileImage = sql.NullString{
			String: request.ProfileImage,
			Valid:  true,
		}
	}

	if err := c.TeacherRepository.Update(tx, &teacher); err != nil {
		c.Log.Warnf("Failed to update teacher : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	return user_converter.TeacherToResponse(&teacher), nil
}

func (c *TeacherUseCase) Delete(ctx context.Context, request *user_model.DeleteTeacherRequest) error {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var teacher user_entity.Teacher
	if err := c.TeacherRepository.FindById(tx, &teacher, request.ID); err != nil {
		c.Log.Warnf("Failed to find teacher : %+v", err)
		return errors.NewNotFound()
	}

	if err := c.TeacherRepository.Delete(tx, &teacher); err != nil {
		c.Log.Warnf("Failed to delete teacher : %+v", err)
		return errors.NewInternalServerError()
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return errors.NewInternalServerError()
	}

	return nil
}

func (c *TeacherUseCase) FindById(ctx context.Context, request *user_model.FindTeacherRequest) (*user_model.TeacherResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var teacher user_entity.Teacher
	if err := c.TeacherRepository.FindById(tx, &teacher, request.ID); err != nil {
		c.Log.Warnf("Failed to find teacher : %+v", err)
		return nil, errors.NewNotFound()
	}

	return user_converter.TeacherToResponse(&teacher), nil
}

func (c *TeacherUseCase) FindByIds(ctx context.Context, request *user_model.FindTeachersRequest) ([]*user_model.TeacherResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var teachers []*user_entity.Teacher
	if err := c.TeacherRepository.FindByIds(tx, &teachers, request.IDs); err != nil {
		c.Log.Warnf("Failed to find teachers : %+v", err)
		return nil, errors.NewNotFound()
	}

	return lop.Map(teachers, func(teacher *user_entity.Teacher, _ int) *user_model.TeacherResponse {
		return user_converter.TeacherToResponse(teacher)
	}), nil
}

func (c *TeacherUseCase) FindByUserID(ctx context.Context, request *user_model.FindTeacherByUserIdRequest) (*user_model.TeacherResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		c.Log.Warnf("Failed to begin transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var teacher user_entity.Teacher
	if err := c.TeacherRepository.FindByUserId(tx, &teacher, request.UserID); err != nil {
		c.Log.Warnf("Failed to find teacher : %+v", err)
		return nil, errors.NewNotFound()
	}

	return user_converter.TeacherToResponse(&teacher), nil
}
