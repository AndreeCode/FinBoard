package handlers

import (
	"finboard/src/core/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *CreditHandler) GetCredit(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return response(
			c,
			fiber.StatusBadRequest,
			"invalid credit id",
			nil,
		)
	}

	credit, err := h.service.ObtainCredit(c.Context(), id)
	if err != nil {
		return response(
			c,
			fiber.StatusNotFound,
			"credit not found",
			nil,
		)
	}

	authMiddleware := middleware.NewAuthMiddleware()
	currentUserID := c.Locals("user_id").(string)
	if credit.UserId != uuid.MustParse(currentUserID) && !authMiddleware.IsAdmin(c) {
		return response(
			c,
			fiber.StatusForbidden,
			"you don't have permission to view this credit",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"credit obtained successfully",
		credit,
	)
}
