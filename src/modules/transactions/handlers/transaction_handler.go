package handlers

import (
	"finboard/src/modules/transactions/services"

	"github.com/gofiber/fiber/v3"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func response(
	c fiber.Ctx,
	status int,
	msg string,
	data interface{},
) error {
	return c.Status(status).JSON(fiber.Map{
		"status": status,
		"msg":    msg,
		"data":   data,
	})
}
