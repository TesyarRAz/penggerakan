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
	c.Fiber.Post("/courses", c.AuthMiddleware, c.UserController.Create)
}
