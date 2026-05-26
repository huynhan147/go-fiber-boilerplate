package routes

import "github.com/gofiber/fiber/v2"

func registerHealth(api fiber.Router) {

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
}
