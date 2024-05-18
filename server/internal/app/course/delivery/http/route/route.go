package course_route

import (
	course_http "github.com/TesyarRAz/penggerak/internal/app/course/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *course_http.CourseController
}

func (c *RouteConfig) Setup() {
	c.App.Get("/courses", c.UserController.List)
	c.App.Post("/courses", c.UserController.Create)
}
