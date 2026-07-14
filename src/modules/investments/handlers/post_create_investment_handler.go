package handlers

import (
	"finboard/src/modules/investments/domains"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *InvestmentHandler) CreateInvestment(c fiber.Ctx) error {

	var req domains.CreateInvestmentRequest

	if err := c.Bind().Body(&req); err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	transactionId, err := uuid.Parse(req.TransactionId)
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid transaction_id uuid",
			nil,
		)
	}

	createdByUUID, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid user_id from token",
			nil,
		)
	}

	investment := &domains.Investment{
		TransactionId: transactionId,
		ExpectedGain:  req.ExpectedGain,
		RiskLevel:     req.RiskLevel,
		Status:        req.Status,
		CreatedBy:     createdByUUID,
	}

	result, err := h.service.CreateInvestment(c.Context(), investment)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error in investment creation",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusCreated,
		"investment created successfully",
		result,
	)
}
