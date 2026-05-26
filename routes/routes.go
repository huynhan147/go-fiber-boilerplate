package routes

import (
	"myapp/app/container"

	"github.com/gofiber/fiber/v2"
)

func Register(
	app *fiber.App,
	c *container.Container,
) {
	api := app.Group("/api")

	registerHealth(api)
	registerAuth(api, c)
	registerUser(api, c)
	registerTransaction(api, c)
}
