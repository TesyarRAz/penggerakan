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
	cfg *config.BootstrapConfig

	courseRepository *course_repository.CourseRepository
	courseUseCase    *course_usecase.CourseUseCase

	moduleRepository *course_repository.ModuleRepository
	moduleUseCase    *course_usecase.ModuleUseCase

	subModuleRepository *course_repository.SubModuleRepository
	subModuleUseCase    *course_usecase.SubModuleUseCase
}

var _ config.App = &App{}

func NewApp(cfg *config.BootstrapConfig) *App {
	courseRepository := course_repository.NewCourseRepository(cfg.Log, cfg.DB)
	courseUseCase := course_usecase.NewCourseUseCase(cfg.DB, cfg.Env, cfg.Log, cfg.Validate, courseRepository)

	moduleRepository := course_repository.NewModuleRepository(cfg.Log, cfg.DB)
	moduleUseCase := course_usecase.NewModuleUseCase(cfg.DB, cfg.Env, cfg.Log, cfg.Validate, courseRepository, moduleRepository)

	subModuleRepository := course_repository.NewSubModuleRepository(cfg.Log, cfg.DB)
	subModuleUseCase := course_usecase.NewSubModuleUseCase(cfg.DB, cfg.Env, cfg.Log, cfg.Validate, moduleRepository, subModuleRepository)

	return &App{
		cfg: cfg,

		courseRepository: courseRepository,
		courseUseCase:    courseUseCase,

		moduleRepository: moduleRepository,
		moduleUseCase:    moduleUseCase,

		subModuleRepository: subModuleRepository,
		subModuleUseCase:    subModuleUseCase,
	}
}

func (a *App) Provider() config.Provider {
	return nil
}

func (a *App) Service(providers config.Provider) {
	authService := providers[service.AUTH_SERVICE].(service.AuthService)
	teacherService := providers[service.TEACHER_SERVICE].(service.TeacherService)

	courseController := course_http.NewCourseController(a.courseUseCase, a.cfg.Log, &teacherService)
	moduleController := course_http.NewModuleController(a.moduleUseCase, a.cfg.Log)
	subModuleController := course_http.NewSubModuleController(a.subModuleUseCase, a.cfg.Log)

	authMiddleware := course_middleware.NewAuth(&authService)

	routeConfig := course_route.RouteConfig{
		Fiber:               a.cfg.Fiber,
		AuthMiddleware:      authMiddleware,
		CourseController:    courseController,
		ModuleController:    moduleController,
		SubModuleController: subModuleController,
	}
	routeConfig.Setup()
}
