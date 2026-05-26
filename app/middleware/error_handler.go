package middleware

import (
	"myapp/app/http/responses"
	"myapp/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	if code >= 500 {
		// Lấy request_id từ context để log cùng
		requestID, _ := c.Locals("request_id").(string)
		logger.Error("Internal server error",
			zap.String("request_id", requestID),
			zap.Int("status", code),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("error", err.Error()),
		)
	}

	// Dùng responses.ServerError để có request_id trong body
	if code == fiber.StatusInternalServerError {
		return responses.ServerError(c, err)
	}

	// Các lỗi fiber khác (404, 405...) trả về cùng format
	return c.Status(code).JSON(fiber.Map{
		"request_id": c.Locals("request_id"),
		"success":    false,
		"message":    message,
	})
}
