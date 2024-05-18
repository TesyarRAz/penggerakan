package course_route

import (
	course_http "github.com/TesyarRAz/penggerak/internal/app/course/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	Fiber          *fiber.App
	AuthMiddleware fiber.Handler
	UserController *course_http.CourseController
}

func (c *RouteConfig) Setup() {
	c.Fiber.Get("/courses", c.AuthMiddleware, c.UserController.List)
	c.Fiber.Get("/courses/:id", c.AuthMiddleware, c.UserController.FindById)
	c.Fiber.Post("/courses", c.AuthMiddleware, c.UserController.Create)
	c.Fiber.Put("/courses/:id", c.AuthMiddleware, c.UserController.Update)
	c.Fiber.Delete("/courses/:id", c.AuthMiddleware, c.UserController.Delete)
}
