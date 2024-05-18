package user_config

import (
	user_http "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http"
	user_middleware "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http/middleware"
	user_route "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http/route"
	user_provider "github.com/TesyarRAz/penggerak/internal/app/user/delivery/provider"
	user_repository "github.com/TesyarRAz/penggerak/internal/app/user/repository"
	user_usecase "github.com/TesyarRAz/penggerak/internal/app/user/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/config"
)

type App struct {
	cfg *config.BootstrapConfig

	userRepository       *user_repository.UserRepository
	permissionRepository *user_repository.PermissionRepository
	userUseCase          *user_usecase.UserUseCase
}

var _ config.App = &App{}

func NewApp(cfg *config.BootstrapConfig) *App {
	userRepository := user_repository.NewUserRepository(cfg.Log, cfg.DB)
	permissionRepository := user_repository.NewPermissionRepository(cfg.Log, cfg.DB)

	userUseCase := user_usecase.NewUserUseCase(cfg.DB, cfg.Env, cfg.Log, cfg.Validate, userRepository, permissionRepository)

	return &App{
		cfg:                  cfg,
		userRepository:       userRepository,
		permissionRepository: permissionRepository,
		userUseCase:          userUseCase,
	}
}

func (a *App) Provider() config.Provider {
	return config.Provider{
		"auth": user_provider.NewAuthProvider(a.userUseCase),
	}
}

func (a *App) Service(_ config.Provider) {
	userController := user_http.NewUserController(a.userUseCase, a.cfg.Log)

	authMiddleware := user_middleware.NewAuth(a.userUseCase)

	routeConfig := user_route.RouteConfig{
		Fiber:          a.cfg.Fiber,
		UserController: userController,
		AuthMiddleware: authMiddleware,
	}

	routeConfig.Setup()
}
