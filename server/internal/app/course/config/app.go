package course_config

import (
	course_http "github.com/TesyarRAz/penggerak/internal/app/course/delivery/http"
	course_route "github.com/TesyarRAz/penggerak/internal/app/course/delivery/http/route"
	course_repository "github.com/TesyarRAz/penggerak/internal/app/course/repository"
	course_usecase "github.com/TesyarRAz/penggerak/internal/app/course/usecase"
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
	courseRepository := course_repository.NewCourseRepository(config.Log, config.DB)

	userUseCase := course_usecase.NewCourseUseCase(config.DB, config.Config, config.Log, config.Validate, courseRepository)

	userController := course_http.NewCourseController(userUseCase, config.Log)

	routeConfig := course_route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}
	routeConfig.Setup()
}
