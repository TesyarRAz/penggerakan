package user_usecase

import (
	"context"
	"time"

	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	user_converter "github.com/TesyarRAz/penggerak/internal/app/user/model/converter"
	user_repository "github.com/TesyarRAz/penggerak/internal/app/user/repository"
	"github.com/TesyarRAz/penggerak/internal/pkg/errors"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
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

func (c *UserUseCase) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {
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

	return &model.Auth{
		ID:   claims["id"].(string),
		Name: claims["sub"].(string),
	}, nil
}

func (c *UserUseCase) Login(ctx context.Context, request *model.LoginUserRequest) (*model.LoginUserResponse, error) {
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
			Field:   request.Email,
			Message: "Email or password is incorrect",
		}
	}

	if err := util.CheckPasswordHash(request.Password, user.Password); err != nil {
		c.Log.Warnf("Failed to check password hash : %+v", err)
		return nil, errors.Unauthorized{
			Field:   request.Email,
			Message: "Email or password is incorrect",
		}
	}

	if err = c.getUserAcl(tx, &user); err != nil {
		c.Log.Warnf("Failed to get user acl : %+v", err)
		return nil, errors.InternalServerError{}
	}

	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.InternalServerError{}
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
		return nil, errors.InternalServerError{}
	}

	return &model.LoginUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Token:     tokenString,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (c *UserUseCase) GetUser(ctx context.Context, request *model.GetUserRequest) (*model.UserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, err
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		return nil, errors.InternalServerError{}
	}
	defer tx.Rollback()

	var user user_entity.User
	if err := c.UserRepository.FindById(tx, &user, request.ID); err != nil {
		c.Log.Warnf("Failed to find user by id : %+v", err)
		return nil, errors.NotFound{}
	}
	if request.IsDetailed {
		if err = c.getUserAcl(tx, &user); err != nil {
			c.Log.Warnf("Failed to get user acl : %+v", err)
			return nil, errors.InternalServerError{}
		}
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, errors.InternalServerError{}
	}

	return user_converter.UserToResponse(&user, request.IsDetailed), nil
}

func (c *UserUseCase) getUserAcl(tx *sqlx.Tx, user *user_entity.User) error {
	if err := c.PermissionRepository.RolesByUser(tx, user); err != nil {
		c.Log.Warnf("Failed to find roles by user : %+v", err)
		return err
	}

	if err := c.PermissionRepository.PermissionsByRoles(tx, user.Roles...); err != nil {
		c.Log.Warnf("Failed to find permissions by roles : %+v", err)
		return err
	}

	if err := c.PermissionRepository.PermissionsByUsers(tx, user); err != nil {
		c.Log.Warnf("Failed to find permissions by user : %+v", err)
		return err
	}

	return nil
}
