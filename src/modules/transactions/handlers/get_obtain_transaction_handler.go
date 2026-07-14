package handlers

import (
	"errors"
	"finboard/src/core/middleware"
	"finboard/src/modules/transactions/domains"
	"finboard/src/modules/transactions/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *TransactionHandler) GetTransaction(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid transaction id",
			nil,
		)
	}

	transaction := &domains.Transaction{
		Id: id,
	}

	result, err := h.service.ObtainTransaction(
		c.Context(),
		transaction,
	)
	if err != nil {

		switch {

		case errors.Is(err, repository.ErrTransactionNotFound):

			return response(
				c,
				fiber.StatusNotFound,
				"transaction not found",
				nil,
			)

		default:

			return response(
				c,
				fiber.StatusInternalServerError,
				"error obtaining transaction",
				nil,
			)
		}
	}

	authMiddleware := middleware.NewAuthMiddleware()
	currentUserID := c.Locals("user_id").(string)
	if result.UserId != uuid.MustParse(currentUserID) && !authMiddleware.IsAdmin(c) {
		return response(
			c,
			fiber.StatusForbidden,
			"you don't have permission to view this transaction",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"transaction obtained successfully",
		result,
	)
}
