package handlers

import (
	"errors"
	"finboard/src/core/middleware"
	"finboard/src/modules/investments/domains"
	"finboard/src/modules/investments/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *InvestmentHandler) GetInvestment(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid investment id",
			nil,
		)
	}

	investment := &domains.Investment{
		Id: id,
	}

	result, err := h.service.ObtainInvestment(
		c.Context(),
		investment,
	)
	if err != nil {

		switch {

		case errors.Is(err, repository.ErrInvestmentNotFound):

			return response(
				c,
				fiber.StatusNotFound,
				"investment not found",
				nil,
			)

		default:

			return response(
				c,
				fiber.StatusInternalServerError,
				"error obtaining investment",
				nil,
			)
		}
	}

	authMiddleware := middleware.NewAuthMiddleware()
	if !authMiddleware.IsAdmin(c) {
		ownershipErr := h.service.CheckOwnership(c.Context(), result.TransactionId, c.Locals("user_id").(string))
		if ownershipErr != nil {
			return response(
				c,
				fiber.StatusForbidden,
				"you don't have permission to view this investment",
				nil,
			)
		}
	}

	return response(
		c,
		fiber.StatusOK,
		"investment obtained successfully",
		result,
	)
}
