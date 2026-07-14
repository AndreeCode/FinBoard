package handlers

import (
	"finboard/src/core/middleware"

	"github.com/gofiber/fiber/v3"
)

func (h *CategoryHandler) GetCategories(c fiber.Ctx) error {
	userId := ""
	authMiddleware := middleware.NewAuthMiddleware()
	if !authMiddleware.IsAdmin(c) {
		userId = c.Locals("user_id").(string)
	}

	result, err := h.service.ObtainCategories(
		c.Context(),
		userId,
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error obtaining categories",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"categories obtained successfully",
		result,
	)
}
