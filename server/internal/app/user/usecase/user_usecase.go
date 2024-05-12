package user_usecase

import (
	"context"
	"time"

	user_entity "github.com/TesyarRAz/penggerak/internal/app/user/entity"
	user_repository "github.com/TesyarRAz/penggerak/internal/app/user/repository"
	"github.com/TesyarRAz/penggerak/internal/pkg/model"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserUseCase struct {
	DB             *sqlx.DB
	Config         util.DotEnvConfig
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *user_repository.UserRepository
}

func NewUserUseCase(db *sqlx.DB, dotenvcfg util.DotEnvConfig, logger *logrus.Logger, validate *validator.Validate,
	userRepository *user_repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		Config:         dotenvcfg,
		Log:            logger,
		Validate:       validate,
		UserRepository: userRepository,
	}
}

func (c *UserUseCase) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {
	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(util.StringOrDefault(c.Config["JWT_SECRET_KEY"], c.Config["APP_ID"])), nil
	}, jwt.WithExpirationRequired(), jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if err != nil {
		c.Log.Warnf("Failed to parse token : %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.Log.Warn("Failed to validate token")
		return nil, fiber.ErrUnauthorized
	}

	return &model.Auth{
		ID:   claims["id"].(string),
		Name: claims["sub"].(string),
	}, nil
}

func (c *UserUseCase) Login(ctx context.Context, request *model.LoginUserRequest) (*model.LoginUserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		return nil, fiber.ErrServiceUnavailable
	}
	defer tx.Rollback()

	var user user_entity.User
	if err := c.UserRepository.FindByEmail(tx, &user, request.Email); err != nil {
		c.Log.Warnf("Failed to find user by email : %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	if err := util.CheckPasswordHash(request.Password, user.Password); err != nil {
		c.Log.Warnf("Failed to check password hash : %+v", err)
		return nil, fiber.ErrUnauthorized
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	secretKey := []byte(c.Config.StringOrDefaultKey("JWT_SECRET", "APP_ID"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"sub": user.Name,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return &model.LoginUserResponse{
		Name:      user.Name,
		Email:     user.Email,
		Token:     tokenString,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (c *UserUseCase) User(ctx context.Context, request *model.GetUserRequest) (*model.UserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Failed to validate request : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	tx, err := c.DB.Beginx()
	if err != nil {
		return nil, fiber.ErrServiceUnavailable
	}
	defer tx.Rollback()

	var user user_entity.User
	if err := c.UserRepository.FindById(tx, &user, request.ID); err != nil {
		c.Log.Warnf("Failed to find user by id : %+v", err)
		return nil, fiber.ErrUnauthorized
	}
	if err := tx.Commit(); err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.UserResponse{
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
