package routes

import (
	"myapp/app/container"
	"myapp/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func registerUser(
	api fiber.Router,
	c *container.Container,
) {
	protected := api.Group("",
		middleware.Auth(
			c.Config,
		),
	)

	users := protected.Group("/users")

	users.Get("/", c.Handlers.User.Index)
	users.Post("/", c.Handlers.User.Store)
	users.Get("/:id", c.Handlers.User.Show)
	users.Put("/:id", c.Handlers.User.Update)
	users.Delete("/:id", c.Handlers.User.Destroy)
}
