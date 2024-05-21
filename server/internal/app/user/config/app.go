package user_config

import (
	user_http "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http"
	user_middleware "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http/middleware"
	user_route "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http/route"
	user_provider "github.com/TesyarRAz/penggerak/internal/app/user/delivery/provider"
	user_repository "github.com/TesyarRAz/penggerak/internal/app/user/repository"
	user_usecase "github.com/TesyarRAz/penggerak/internal/app/user/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/config"
	"github.com/TesyarRAz/penggerak/internal/pkg/repository"
	"github.com/TesyarRAz/penggerak/internal/pkg/service"
)

type App struct {
	cfg *config.BootstrapConfig

	redisRepository *repository.RedisRepository

	userRepository *user_repository.UserRepository
	userUseCase    *user_usecase.UserUseCase

	permissionRepository *user_repository.PermissionRepository

	teacherRepository *user_repository.TeacherRepository
	teacherUseCase    *user_usecase.TeacherUseCase
}

var _ config.App = &App{}

func NewApp(cfg *config.BootstrapConfig) *App {
	redisRepository := repository.NewRedisRepository(cfg.Redis)
	userRepository := user_repository.NewUserRepository(cfg.Log, cfg.DB)
	permissionRepository := user_repository.NewPermissionRepository(cfg.Log, cfg.DB)
	teacherRepository := user_repository.NewTeacherRepository(cfg.Log, cfg.DB)

	userUseCase := user_usecase.NewUserUseCase(cfg.DB, cfg.Env, cfg.Log, cfg.Validate, userRepository, permissionRepository, redisRepository)
	teacherUseCase := user_usecase.NewTeacherUseCase(cfg.DB, cfg.Env, cfg.Log, cfg.Validate, userRepository, teacherRepository)

	return &App{
		cfg: cfg,

		redisRepository: redisRepository,

		userRepository: userRepository,
		userUseCase:    userUseCase,

		permissionRepository: permissionRepository,

		teacherRepository: teacherRepository,
		teacherUseCase:    teacherUseCase,
	}
}

func (a *App) Provider() config.Provider {
	return config.Provider{
		service.AUTH_SERVICE:    user_provider.NewAuthProvider(a.userUseCase),
		service.TEACHER_SERVICE: user_provider.NewTeacherProvider(a.teacherUseCase),
	}
}

func (a *App) Service(_ config.Provider) {
	userController := user_http.NewUserController(a.userUseCase, a.cfg.Log)
	teacherController := user_http.NewTeacherController(a.teacherUseCase, a.cfg.Log)

	authMiddleware := user_middleware.NewAuth(a.userUseCase)

	routeConfig := user_route.RouteConfig{
		Fiber:             a.cfg.Fiber,
		UserController:    userController,
		TeacherController: teacherController,
		AuthMiddleware:    authMiddleware,
	}

	routeConfig.Setup()
}
