package responses

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type PaginatedResponse struct {
	Success bool         `json:"success"`
	Data    interface{}  `json:"data"`
	Meta    PaginateMeta `json:"meta"`
}

type PaginateMeta struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Pages int64 `json:"pages"`
}

func OK(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{Success: true, Data: data})
}

func Created(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{Success: true, Data: data})
}

func Paginated(c *fiber.Ctx, data interface{}, total int64, page, limit int) error {
	pages := total / int64(limit)
	if total%int64(limit) > 0 {
		pages++
	}
	return c.Status(fiber.StatusOK).JSON(PaginatedResponse{
		Success: true,
		Data:    data,
		Meta:    PaginateMeta{Total: total, Page: page, Limit: limit, Pages: pages},
	})
}

func BadRequest(c *fiber.Ctx, errors interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{Success: false, Errors: errors})
}

func Unauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{Success: false, Message: "Unauthorized"})
}

func Forbidden(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(Response{Success: false, Message: "Forbidden"})
}

func NotFound(c *fiber.Ctx, resource string) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{Success: false, Message: resource + " not found"})
}

func ServerError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Response{Success: false, Message: err.Error()})
}
