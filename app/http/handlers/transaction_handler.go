package handlers

import (
	"myapp/app/services"
	"myapp/models"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(
	service services.TransactionService,
) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

func (h *TransactionHandler) Index(
	c *fiber.Ctx,
) error {

	data, err := h.service.GetAll()

	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (h *TransactionHandler) Store(
	c *fiber.Ctx,
) error {

	var tx models.Transaction

	if err := c.BodyParser(&tx); err != nil {
		return err
	}

	err := h.service.Create(&tx)

	if err != nil {
		return err
	}

	return c.JSON(tx)
}
