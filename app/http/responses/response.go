package responses

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	RequestID string      `json:"request_id,omitempty"`
	Success   bool        `json:"success"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
}

type PaginatedResponse struct {
	RequestID string       `json:"request_id,omitempty"`
	Success   bool         `json:"success"`
	Data      interface{}  `json:"data"`
	Meta      PaginateMeta `json:"meta"`
}

type PaginateMeta struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Pages int64 `json:"pages"`
}

// requestID lấy request_id đã được inject bởi RequestID middleware
func requestID(c *fiber.Ctx) string {
	id, _ := c.Locals("request_id").(string)
	return id
}

func OK(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		RequestID: requestID(c),
		Success:   true,
		Data:      data,
	})
}

func Created(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		RequestID: requestID(c),
		Success:   true,
		Data:      data,
	})
}

func Paginated(c *fiber.Ctx, data interface{}, total int64, page, limit int) error {
	pages := total / int64(limit)
	if total%int64(limit) > 0 {
		pages++
	}
	return c.Status(fiber.StatusOK).JSON(PaginatedResponse{
		RequestID: requestID(c),
		Success:   true,
		Data:      data,
		Meta:      PaginateMeta{Total: total, Page: page, Limit: limit, Pages: pages},
	})
}

func BadRequest(c *fiber.Ctx, errors interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		RequestID: requestID(c),
		Success:   false,
		Errors:    errors,
	})
}

func Unauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		RequestID: requestID(c),
		Success:   false,
		Message:   "Unauthorized",
	})
}

func Forbidden(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(Response{
		RequestID: requestID(c),
		Success:   false,
		Message:   "Forbidden",
	})
}

func NotFound(c *fiber.Ctx, resource string) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		RequestID: requestID(c),
		Success:   false,
		Message:   resource + " not found",
	})
}

func ServerError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		RequestID: requestID(c),
		Success:   false,
		Message:   err.Error(),
	})
}
