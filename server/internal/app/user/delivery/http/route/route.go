package user_route

import (
	user_http "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http"
	user_middleware "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http/middleware"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	Fiber *fiber.App

	UserController    *user_http.UserController
	TeacherController *user_http.TeacherController

	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	admin := user_middleware.Role("admin")

	auth := c.Fiber.Group("/auth")
	auth.Post("/login", c.UserController.Login)
	auth.Delete("/logout", c.AuthMiddleware, c.UserController.Logout)
	auth.Get("/me", c.AuthMiddleware, c.UserController.Me)
	auth.Post("/refresh", c.UserController.RefreshToken)

	users := c.Fiber.Group("/users", c.AuthMiddleware, admin)
	users.Post("/", c.UserController.Create)
	users.Put("/:id", c.UserController.Update)
	users.Get("/", c.UserController.List)
	users.Get("/:id", c.UserController.FindById)
	users.Delete("/:id", c.UserController.Delete)
	users.Post("/:id/roles", c.UserController.AttachRole)
	users.Delete("/:id/roles", c.UserController.DetachRole)

	teachers := c.Fiber.Group("/teachers", c.AuthMiddleware, admin)
	teachers.Post("/", c.TeacherController.Create)
	teachers.Put("/:id", c.TeacherController.Update)
	teachers.Get("/", c.TeacherController.List)
	teachers.Get("/:id", c.TeacherController.FindById)
	teachers.Delete("/:id", c.TeacherController.Delete)
}
