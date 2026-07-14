package handlers

import (
	"finboard/src/modules/credits/repository"
	"finboard/src/modules/credits/services"

	"github.com/gofiber/fiber/v3"
)

type CreditHandler struct {
	service *services.CreditService
}

func NewCreditHandler(service *services.CreditService) *CreditHandler {
	return &CreditHandler{service: service}
}

func NewCreditRepository() *repository.CreditRepository {
	return repository.NewCreditRepository()
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
