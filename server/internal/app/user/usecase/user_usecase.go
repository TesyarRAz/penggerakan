package user_usecase

import (
	"context"
	"database/sql"
	"time"

	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	user_model "github.com/TesyarRAz/penggerak/internal/app/user/model"
	user_converter "github.com/TesyarRAz/penggerak/internal/app/user/model/converter"
	user_repository "github.com/TesyarRAz/penggerak/internal/app/user/repository"
	"github.com/TesyarRAz/penggerak/internal/pkg/errors"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	shared_model "github.com/TesyarRAz/penggerak/internal/pkg/model/shared"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"github.com/sirupsen/logrus"
)

type UserUseCase struct {
	DB                   *sqlx.DB
	Config               util.DotEnvConfig
	Log                  *logrus.Logger
	Validate             *validator.Validate
	UserRepository       *user_repository.UserRepository
	PermissionRepository *user_repository.PermissionRepository
}

func NewUserUseCase(db *sqlx.DB, dotenvcfg util.DotEnvConfig, logger *logrus.Logger, validate *validator.Validate,
	userRepository *user_repository.UserRepository, permissionRepository *user_repository.PermissionRepository) *UserUseCase {
	return &UserUseCase{
		DB:                   db,
		Config:               dotenvcfg,
		Log:                  logger,
		Validate:             validate,
		UserRepository:       userRepository,
		PermissionRepository: permissionRepository,
	}
}

func (c *UserUseCase) Verify(ctx context.Context, request *shared_model.VerifyUserRequest) (*model.Auth, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, err
	}

	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(util.StringOrDefault(c.Config["JWT_SECRET_KEY"], c.Config["APP_ID"])), nil
	}, jwt.WithExpirationRequired(), jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if err != nil {
		c.Log.Warnf("Failed to parse token : %+v", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.Log.Warn("Failed to validate token")
		return nil, err
	}

	return model.NewAuth(claims["id"].(string), claims["sub"].(string), claims), nil
}

func (c *UserUseCase) Login(ctx context.Context, request *user_model.LoginUserRequest) (*user_model.LoginUserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var user user_entity.User
	if err := c.UserRepository.FindByEmail(tx, &user, request.Email); err != nil {
		c.Log.Warnf("Failed to find user by email : %+v", err)
		return nil, errors.Unauthorized{
			Message: "Email or password is incorrect",
		}
	}

	if err := util.CheckPasswordHash(request.Password, user.Password); err != nil {
		c.Log.Warnf("Failed to check password hash : %+v", err)
		return nil, errors.Unauthorized{
			Message: "Email or password is incorrect",
		}
	}

	if err = c.getUserAcl(tx, &user); err != nil {
		c.Log.Warnf("Failed to get user acl : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	secretKey := []byte(c.Config.StringOrDefaultKey("JWT_SECRET", "APP_ID"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          user.ID,
		"sub":         user.Name,
		"roles":       user.Roles,
		"permissions": user.Permissions,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.Log.Warnf("Failed to sign token : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	return &user_model.LoginUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Token:     tokenString,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (c *UserUseCase) FindUserById(ctx context.Context, request *user_model.FindUserRequest) (*user_model.UserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		return nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var user user_entity.User
	if err := c.UserRepository.FindById(tx, &user, request.ID); err != nil {
		c.Log.Warnf("Failed to find user by id : %+v", err)
		return nil, errors.NewNotFound()
	}
	if request.IsDetailed {
		if err = c.getUserAcl(tx, &user); err != nil {
			c.Log.Warnf("Failed to get user acl : %+v", err)
			return nil, errors.NewInternalServerError()
		}
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	return user_converter.UserToResponse(&user, request.IsDetailed), nil
}

func (c *UserUseCase) List(ctx context.Context, request *model.PageRequest) ([]*user_model.UserResponse, *model.PageMetadata, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, nil, err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		return nil, nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var users []*user_entity.User
	pageInfo, err := c.UserRepository.List(tx, &users, request)
	if err != nil {
		c.Log.Warnf("Failed to list users : %+v", err)
		return nil, nil, errors.NewInternalServerError()
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, nil, errors.NewInternalServerError()
	}

	return lop.Map(users, func(user *user_entity.User, _ int) *user_model.UserResponse {
		return user_converter.UserToResponse(user, false)
	}), pageInfo, nil
}

func (c *UserUseCase) Create(ctx context.Context, request *user_model.CreateUserRequest) (*user_model.UserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		return nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	user := user_entity.User{
		Email:    request.Email,
		Password: util.HashPassword(request.Password),
		Name:     request.Name,
	}

	if request.ProfileImage != "" {
		user.ProfileImage = sql.NullString{
			String: request.ProfileImage,
			Valid:  true,
		}
	}

	if err = c.UserRepository.Create(tx, &user); err != nil {
		c.Log.Warnf("Failed to create user : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	return user_converter.UserToResponse(&user, true), nil
}

func (c *UserUseCase) Update(ctx context.Context, request *user_model.UpdateUserRequest) (*user_model.UserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		return nil, errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var user user_entity.User
	if err := c.UserRepository.FindById(tx, &user, request.ID); err != nil {
		c.Log.Warnf("Failed to find user : %+v", err)
		return nil, errors.NewNotFound()
	}

	user.Name = request.Name
	user.Email = request.Email
	if request.Password != "" {
		user.Password = util.HashPassword(request.Password)
	}
	if request.ProfileImage != "" {
		user.ProfileImage = sql.NullString{
			String: request.ProfileImage,
			Valid:  true,
		}
	}

	if err := c.UserRepository.Update(tx, &user); err != nil {
		c.Log.Warnf("Failed to update user : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.NewInternalServerError()
	}

	return user_converter.UserToResponse(&user, true), nil
}

func (c *UserUseCase) Delete(ctx context.Context, request *user_model.DeleteUserRequest) error {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		return errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var user user_entity.User
	if err := c.UserRepository.FindById(tx, &user, request.ID); err != nil {
		c.Log.Warnf("Failed to find user : %+v", err)
		return errors.NewNotFound()
	}

	if err := c.UserRepository.Delete(tx, &user); err != nil {
		c.Log.Warnf("Failed to delete user : %+v", err)
		return errors.NewInternalServerError()
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return errors.NewInternalServerError()
	}

	return nil
}

func (c *UserUseCase) AttachRole(ctx context.Context, request *user_model.AttachRoleToUserRequest) error {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		return errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var user user_entity.User
	if err := c.UserRepository.FindById(tx, &user, request.ID); err != nil {
		c.Log.Warnf("Failed to find user : %+v", err)
		return errors.NewNotFound()
	}

	var role user_entity.Role
	if err := c.PermissionRepository.FindRoleByName(tx, &role, request.Role); err != nil {
		c.Log.Warnf("Failed to find role : %+v", err)
		return errors.NewNotFound()
	}

	ok, err := c.PermissionRepository.UserHasRoles(tx, user.ID, role.ID)
	if err != nil {
		c.Log.Warnf("Failed to check user has roles : %+v", err)
		return errors.NewInternalServerError()
	}
	if ok {
		return errors.Conflict{
			Message: "User already has the role",
		}
	}

	if err := c.PermissionRepository.AttachRoleToUser(tx, role.ID, user.ID); err != nil {
		c.Log.Warnf("Failed to attach role to user : %+v", err)
		return errors.NewInternalServerError()
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return errors.NewInternalServerError()
	}

	return nil
}

func (c *UserUseCase) DetachRole(ctx context.Context, request *user_model.DetachRoleFromUserRequest) error {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		return errors.NewInternalServerError()
	}
	defer tx.Rollback()

	var user user_entity.User
	if err := c.UserRepository.FindById(tx, &user, request.ID); err != nil {
		c.Log.Warnf("Failed to find user : %+v", err)
		return errors.NewNotFound()
	}

	var role user_entity.Role
	if err := c.PermissionRepository.FindRoleByName(tx, &role, request.Role); err != nil {
		c.Log.Warnf("Failed to find role : %+v", err)
		return errors.NewNotFound()
	}

	if err := c.PermissionRepository.DetachRoleFromUser(tx, role.ID, user.ID); err != nil {
		c.Log.Warnf("Failed to detach role from user : %+v", err)
		return errors.NewInternalServerError()
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return errors.NewInternalServerError()
	}

	return nil
}

func (c *UserUseCase) getUserAcl(tx *sqlx.Tx, users ...*user_entity.User) error {
	if err := c.PermissionRepository.RolesByUser(tx, users...); err != nil {
		c.Log.Warnf("Failed to find roles by user : %+v", err)
		return err
	}

	roles := lo.Flatten(lop.Map(users, func(user *user_entity.User, _ int) []*user_entity.Role {
		return user.Roles
	}))
	if err := c.PermissionRepository.PermissionsByRoles(tx, roles...); err != nil {
		c.Log.Warnf("Failed to find permissions by roles : %+v", err)
		return err
	}

	if err := c.PermissionRepository.PermissionsByUsers(tx, users...); err != nil {
		c.Log.Warnf("Failed to find permissions by user : %+v", err)
		return err
	}

	return nil
}
