package handlers

import (
	"finboard/src/modules/investments/services"

	"github.com/gofiber/fiber/v3"
)

type InvestmentHandler struct {
	service *services.InvestmentService
}

func NewInvestmentHandler(service *services.InvestmentService) *InvestmentHandler {
	return &InvestmentHandler{service: service}
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
