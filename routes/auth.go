package routes

import (
	"myapp/app/container"
	"myapp/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func registerAuth(
	api fiber.Router,
	c *container.Container,
) {
	auth := api.Group("/auth")

	auth.Post("/login", c.Handlers.Auth.Login)

	protected := api.Group("",
		middleware.Auth(
			c.Config,
		),
	)

	protected.Get("/auth/me", c.Handlers.Auth.Me)
	protected.Post("/auth/logout", c.Handlers.Auth.Logout)
}
