package handlers

import (
	"errors"
	"finboard/src/core/middleware"
	"finboard/src/modules/transactions/domains"
	"finboard/src/modules/transactions/repository"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *TransactionHandler) DeleteTransaction(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response(
			c,
			fiber.StatusBadRequest,
			"invalid transaction id",
			nil,
		)
	}

	existingTransaction := &domains.Transaction{Id: id}
	existingTransaction, err = h.service.ObtainTransaction(c.Context(), existingTransaction)
	if err != nil {
		if errors.Is(err, repository.ErrTransactionNotFound) {
			return response(c, fiber.StatusNotFound, "transaction not found", nil)
		}
		return response(c, fiber.StatusInternalServerError, "error obtaining transaction", nil)
	}

	authMiddleware := middleware.NewAuthMiddleware()
	currentUserID := c.Locals("user_id").(string)
	if existingTransaction.UserId != uuid.MustParse(currentUserID) && !authMiddleware.IsAdmin(c) {
		return response(c, fiber.StatusForbidden, "you don't have permission to delete this transaction", nil)
	}

	now := time.Now()

	transaction := &domains.Transaction{
		Id:        id,
		DeletedAt: &now,
	}

	err = h.service.DeleteTransaction(
		c.Context(),
		transaction,
	)
	if err != nil {
		return response(
			c,
			fiber.StatusInternalServerError,
			"error deleting transaction",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"transaction deleted successfully",
		nil,
	)
}
