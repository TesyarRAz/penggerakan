package course_route

import (
	course_http "github.com/TesyarRAz/penggerak/internal/app/course/delivery/http"
	course_middleware "github.com/TesyarRAz/penggerak/internal/app/course/delivery/http/middleware"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	Fiber               *fiber.App
	AuthMiddleware      fiber.Handler
	CourseController    *course_http.CourseController
	ModuleController    *course_http.ModuleController
	SubModuleController *course_http.SubModuleController
}

func (c *RouteConfig) Setup() {
	roleAdmin := course_middleware.Role("admin")

	courses := c.Fiber.Group("/courses", c.AuthMiddleware)

	courses.Get("/", c.CourseController.List)
	courses.Get("/:id", c.CourseController.FindById)
	courses.Post("/", roleAdmin, c.CourseController.Create)
	courses.Put("/:id", roleAdmin, c.CourseController.Update)
	courses.Delete("/:id", roleAdmin, c.CourseController.Delete)

	modules := c.Fiber.Group("/modules/:course_id", c.AuthMiddleware)
	modules.Get("/", c.ModuleController.List)
	modules.Get("/:id", c.ModuleController.FindById)
	modules.Post("/", roleAdmin, c.ModuleController.Create)
	modules.Put("/:id", roleAdmin, c.ModuleController.Update)
	modules.Delete("/:id", roleAdmin, c.ModuleController.Delete)

	subModules := c.Fiber.Group("/submodules/:module_id", c.AuthMiddleware)
	subModules.Get("/", c.SubModuleController.List)
	subModules.Get("/:id", c.SubModuleController.FindById)
	subModules.Post("/", roleAdmin, c.SubModuleController.Create)
	subModules.Put("/:id", roleAdmin, c.SubModuleController.Update)
	subModules.Delete("/:id", roleAdmin, c.SubModuleController.Delete)
}
