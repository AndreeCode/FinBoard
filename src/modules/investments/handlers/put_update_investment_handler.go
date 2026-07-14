package handlers

import (
	"errors"
	"finboard/src/core/middleware"
	"finboard/src/modules/investments/domains"
	"finboard/src/modules/investments/repository"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *InvestmentHandler) UpdateInvestment(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid investment id",
			nil,
		)
	}

	existingInvestment := &domains.Investment{Id: id}
	existingInvestment, err = h.service.ObtainInvestment(c.Context(), existingInvestment)
	if err != nil {
		if errors.Is(err, repository.ErrInvestmentNotFound) {
			return response(c, fiber.StatusNotFound, "investment not found", nil)
		}
		return response(c, fiber.StatusInternalServerError, "error obtaining investment", nil)
	}

	authMiddleware := middleware.NewAuthMiddleware()
	if !authMiddleware.IsAdmin(c) {
		ownershipErr := h.service.CheckOwnership(c.Context(), existingInvestment.TransactionId, c.Locals("user_id").(string))
		if ownershipErr != nil {
			return response(c, fiber.StatusForbidden, "you don't have permission to update this investment", nil)
		}
	}

	var req domains.UpdateInvestmentRequest

	if err := c.Bind().Body(&req); err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	now := time.Now()

	investment := &domains.Investment{
		Id:        id,
		UpdatedAt: &now,
	}

	if req.ExpectedGain != nil {
		investment.ExpectedGain = *req.ExpectedGain
	}

	if req.RiskLevel != nil {
		investment.RiskLevel = *req.RiskLevel
	}

	if req.Status != nil {
		investment.Status = *req.Status
	}

	result, err := h.service.UpdateInvestment(
		c.Context(),
		investment,
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error updating investment",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"investment updated successfully",
		result,
	)
}
