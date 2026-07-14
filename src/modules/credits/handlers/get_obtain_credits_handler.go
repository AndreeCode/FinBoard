package handlers

import (
	"finboard/src/core/middleware"

	"github.com/gofiber/fiber/v3"
)

func (h *CreditHandler) GetCredits(c fiber.Ctx) error {
	userId := ""
	authMiddleware := middleware.NewAuthMiddleware()
	if !authMiddleware.IsAdmin(c) {
		userId = c.Locals("user_id").(string)
	}

	credits, err := h.service.ObtainCredits(c.Context(), userId)
	if err != nil {
		return response(
			c,
			fiber.StatusInternalServerError,
			"error obtaining credits",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"credits obtained successfully",
		credits,
	)
}
