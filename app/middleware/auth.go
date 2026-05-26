package middleware

import (
	"myapp/app/http/responses"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Auth(cfg *viper.Viper) fiber.Handler {

	secret := cfg.GetString("API_SECRET_KEY")

	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")

		if authHeader == "" ||
			!strings.HasPrefix(authHeader, "Bearer ") {
			return responses.Unauthorized(c)
		}

		token := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		if token != secret {
			return responses.Unauthorized(c)
		}

		return c.Next()
	}
}
