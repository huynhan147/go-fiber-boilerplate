package handlers

import (
	"myapp/app/http/requests"
	"myapp/app/http/responses"
	"myapp/app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Index GET /api/users
func (h *UserHandler) Index(c *fiber.Ctx) error {
	page, _  := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "15"))

	users, total, err := h.userService.GetAll(page, limit)
	if err != nil {
		return responses.ServerError(c, err)
	}

	return responses.Paginated(c, responses.NewUserCollection(users), total, page, limit)
}

// Show GET /api/users/:id
func (h *UserHandler) Show(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return responses.BadRequest(c, "invalid id")
	}

	user, err := h.userService.GetByID(uint(id))
	if err != nil {
		return responses.NotFound(c, "User")
	}

	return responses.OK(c, responses.NewUserResponse(user))
}

// Store POST /api/users
func (h *UserHandler) Store(c *fiber.Ctx) error {
	req := new(requests.CreateUserRequest)
	if errs := requests.ParseAndValidate(c, req); errs != nil {
		return responses.BadRequest(c, errs)
	}

	user, err := h.userService.Create(req.Name, req.Email, req.Password)
	if err != nil {
		return responses.BadRequest(c, err.Error())
	}

	return responses.Created(c, responses.NewUserResponse(user))
}

// Update PUT /api/users/:id
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return responses.BadRequest(c, "invalid id")
	}

	req := new(requests.UpdateUserRequest)
	if errs := requests.ParseAndValidate(c, req); errs != nil {
		return responses.BadRequest(c, errs)
	}

	user, err := h.userService.Update(uint(id), req.Name, req.Email)
	if err != nil {
		return responses.NotFound(c, "User")
	}

	return responses.OK(c, responses.NewUserResponse(user))
}

// Destroy DELETE /api/users/:id
func (h *UserHandler) Destroy(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return responses.BadRequest(c, "invalid id")
	}

	if err := h.userService.Delete(uint(id)); err != nil {
		return responses.NotFound(c, "User")
	}

	return responses.OK(c, nil)
}
