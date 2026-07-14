package handlers

import (
	"finboard/src/core/middleware"

	"github.com/gofiber/fiber/v3"
)

func (h *InvestmentHandler) GetInvestments(c fiber.Ctx) error {
	userId := ""
	authMiddleware := middleware.NewAuthMiddleware()
	if !authMiddleware.IsAdmin(c) {
		userId = c.Locals("user_id").(string)
	}

	result, err := h.service.ObtainInvestments(
		c.Context(),
		userId,
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error obtaining investments",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"investments obtained successfully",
		result,
	)
}
