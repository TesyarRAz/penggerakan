package course_config

import (
	course_http "github.com/TesyarRAz/penggerak/internal/app/course/delivery/http"
	course_middleware "github.com/TesyarRAz/penggerak/internal/app/course/delivery/http/middleware"
	course_route "github.com/TesyarRAz/penggerak/internal/app/course/delivery/http/route"
	course_repository "github.com/TesyarRAz/penggerak/internal/app/course/repository"
	course_usecase "github.com/TesyarRAz/penggerak/internal/app/course/usecase"
	"github.com/TesyarRAz/penggerak/internal/pkg/config"
	"github.com/TesyarRAz/penggerak/internal/pkg/service"
)

type App struct {
	cfg              *config.BootstrapConfig
	courseRepository *course_repository.CourseRepository
	courseUseCase    *course_usecase.CourseUseCase
}

var _ config.App = &App{}

func NewApp(cfg *config.BootstrapConfig) *App {
	courseRepository := course_repository.NewCourseRepository(cfg.Log, cfg.DB)

	courseUseCase := course_usecase.NewCourseUseCase(cfg.DB, cfg.Env, cfg.Log, cfg.Validate, courseRepository)

	return &App{
		cfg:              cfg,
		courseRepository: courseRepository,
		courseUseCase:    courseUseCase,
	}
}

func (a *App) Provider() config.Provider {
	return nil
}

func (a *App) Service(providers config.Provider) {
	userController := course_http.NewCourseController(a.courseUseCase, a.cfg.Log)

	authService := providers["auth"].(service.AuthService)

	authMiddleware := course_middleware.NewAuth(&authService)

	routeConfig := course_route.RouteConfig{
		Fiber:          a.cfg.Fiber,
		AuthMiddleware: authMiddleware,
		UserController: userController,
	}
	routeConfig.Setup()
}
