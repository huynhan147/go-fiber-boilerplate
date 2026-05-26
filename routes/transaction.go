package routes

import (
	"myapp/app/container"
	"myapp/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func registerTransaction(
	api fiber.Router,
	c *container.Container,
) {

	protected := api.Group("",
		middleware.Auth(
			c.Config,
		),
	)

	tx := protected.Group("/transactions")

	tx.Get("/", c.Handlers.Transaction.Index)
	tx.Post("/", c.Handlers.Transaction.Store)
}
