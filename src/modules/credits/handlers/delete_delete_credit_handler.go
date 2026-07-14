package handlers

import (
	"finboard/src/core/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *CreditHandler) DeleteCredit(c fiber.Ctx) error {
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

	existingCredit, err := h.service.ObtainCredit(c.Context(), id)
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
	if existingCredit.UserId != uuid.MustParse(currentUserID) && !authMiddleware.IsAdmin(c) {
		return response(c, fiber.StatusForbidden, "you don't have permission to delete this credit", nil)
	}

	if err := h.service.DeleteCredit(c.Context(), id); err != nil {
		return response(
			c,
			fiber.StatusInternalServerError,
			"error deleting credit",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusNoContent,
		"credit deleted successfully",
		nil,
	)
}
