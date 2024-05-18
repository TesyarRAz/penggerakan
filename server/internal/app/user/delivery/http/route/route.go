package user_route

import (
	user_http "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	Fiber *fiber.App

	UserController *user_http.UserController

	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.Fiber.Post("/auth/login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.Fiber.Delete("/auth/logout", c.AuthMiddleware, c.UserController.Logout)

	c.Fiber.Get("/auth/me", c.AuthMiddleware, c.UserController.Me)
}
