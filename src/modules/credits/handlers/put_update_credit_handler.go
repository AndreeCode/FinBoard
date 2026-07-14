package handlers

import (
	"finboard/src/core/middleware"
	"finboard/src/modules/credits/domains"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *CreditHandler) UpdateCredit(c fiber.Ctx) error {
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
		return response(c, fiber.StatusForbidden, "you don't have permission to update this credit", nil)
	}

	var req domains.UpdateCreditRequest

	if err := c.Bind().Body(&req); err != nil {
		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	if req.PersonName != nil {
		existingCredit.PersonName = *req.PersonName
	}
	if req.Amount != nil {
		existingCredit.Amount = *req.Amount
	}
	if req.InterestRate != nil {
		existingCredit.InterestRate = *req.InterestRate
	}
	if req.IsCreditor != nil {
		existingCredit.IsCreditor = *req.IsCreditor
	}
	if req.IsSecure != nil {
		existingCredit.IsSecure = *req.IsSecure
	}
	if req.DueDate != nil {
		t, err := time.Parse(time.RFC3339, *req.DueDate)
		if err == nil {
			existingCredit.DueDate = &t
		}
	}
	if req.Status != nil {
		existingCredit.Status = *req.Status
	}

	updatedCredit, err := h.service.UpdateCredit(c.Context(), id, &existingCredit)
	if err != nil {
		return response(
			c,
			fiber.StatusInternalServerError,
			"error updating credit",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"credit updated successfully",
		updatedCredit,
	)
}
