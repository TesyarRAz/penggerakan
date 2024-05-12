package user_route

import (
	user_http "github.com/TesyarRAz/penggerak/internal/app/user/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App *fiber.App

	UserController *user_http.UserController

	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/v1/auth/login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Delete("/v1/auth/logout", c.UserController.Logout)

	c.App.Get("/v1/auth/me", c.UserController.Me)
}
