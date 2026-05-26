package handlers

import (
	"myapp/app/http/requests"
	"myapp/app/http/responses"
	"myapp/app/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login POST /api/auth/login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := new(requests.LoginRequest)
	if errs := requests.ParseAndValidate(c, req); errs != nil {
		return responses.BadRequest(c, errs)
	}

	token, user, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		return responses.Unauthorized(c)
	}

	return responses.OK(c, responses.NewAuthResponse(token, user))
}

// Logout POST /api/auth/logout
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	tokenStr, ok := c.Locals("token").(string)
	if !ok || tokenStr == "" {
		return responses.Unauthorized(c)
	}

	if err := h.authService.Logout(tokenStr); err != nil {
		return responses.ServerError(c, err)
	}

	return responses.OK(c, fiber.Map{"message": "logged out successfully"})
}

// Me GET /api/auth/me
func (h *AuthHandler) Me(c *fiber.Ctx) error {
	return responses.OK(c, fiber.Map{
		"id":    c.Locals("userID"),
		"email": c.Locals("email"),
	})
}
