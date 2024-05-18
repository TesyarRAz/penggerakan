package user_config

import (
	user_http "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http"
	user_middleware "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http/middleware"
	user_route "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http/route"
	user_repository "github.com/TesyarRAz/penggerak/internal/app/user/repository"
	user_usecase "github.com/TesyarRAz/penggerak/internal/app/user/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type BootstrapConfig struct {
	App      *fiber.App
	DB       *sqlx.DB
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   util.DotEnvConfig
}

func Bootstrap(config *BootstrapConfig) {
	userRepository := user_repository.NewUserRepository(config.Log, config.DB)
	permissionRepository := user_repository.NewPermissionRepository(config.Log, config.DB)

	userUseCase := user_usecase.NewUserUseCase(config.DB, config.Config, config.Log, config.Validate, userRepository, permissionRepository)

	userController := user_http.NewUserController(userUseCase, config.Log)

	// setup middleware
	authMiddleware := user_middleware.NewAuth(userUseCase)

	routeConfig := user_route.RouteConfig{
		App:            config.App,
		UserController: userController,
		AuthMiddleware: authMiddleware,
	}

	routeConfig.Setup()
}
