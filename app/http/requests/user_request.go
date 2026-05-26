package requests

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type CreateUserRequest struct {
	Name     string `json:"name"     validate:"required,min=2,max=100"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"  validate:"required,min=2,max=100"`
	Email string `json:"email" validate:"required,email"`
}

type LoginRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// ParseAndValidate parses body và trả về map lỗi nếu có
func ParseAndValidate(c *fiber.Ctx, req interface{}) map[string]string {
	if err := c.BodyParser(req); err != nil {
		return map[string]string{"body": "invalid request body"}
	}

	errs := validate.Struct(req)
	if errs == nil {
		return nil
	}

	fieldErrors := make(map[string]string)
	for _, e := range errs.(validator.ValidationErrors) {
		fieldErrors[e.Field()] = e.Tag()
	}
	return fieldErrors
}
