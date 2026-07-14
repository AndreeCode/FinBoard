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

func (h *TransactionHandler) UpdateTransaction(c fiber.Ctx) error {

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
		return response(c, fiber.StatusForbidden, "you don't have permission to update this transaction", nil)
	}

	var req domains.UpdateTransactionRequest

	if err := c.Bind().Body(&req); err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	now := time.Now()

	transaction := &domains.Transaction{
		Id:        id,
		UpdatedAt: &now,
	}

	if req.CategoryId != nil {
		categoryId, err := uuid.Parse(*req.CategoryId)
		if err != nil {

			return response(
				c,
				fiber.StatusBadRequest,
				"invalid category_id uuid",
				nil,
			)
		}
		transaction.CategoryId = &categoryId
	}

	if req.Amount != nil {
		transaction.Amount = *req.Amount
	}

	if req.Type != nil {
		transaction.Type = *req.Type
	}

	if req.TransactionDate != nil {
		transactionDate, err := time.Parse(time.RFC3339, *req.TransactionDate)
		if err != nil {

			return response(
				c,
				fiber.StatusBadRequest,
				"invalid transaction_date format",
				nil,
			)
		}
		transaction.TransactionDate = transactionDate
	}

	if req.ReceivedDate != nil {
		receivedDate, err := time.Parse(time.RFC3339, *req.ReceivedDate)
		if err != nil {

			return response(
				c,
				fiber.StatusBadRequest,
				"invalid received_date format",
				nil,
			)
		}
		transaction.ReceivedDate = &receivedDate
	}

	if req.DueDate != nil {
		dueDate, err := time.Parse(time.RFC3339, *req.DueDate)
		if err != nil {

			return response(
				c,
				fiber.StatusBadRequest,
				"invalid due_date format",
				nil,
			)
		}
		transaction.DueDate = &dueDate
	}

	if req.Canceled != nil {
		transaction.Canceled = *req.Canceled
	}

	if req.Description != nil {
		transaction.Description = *req.Description
	}

	result, err := h.service.UpdateTransaction(
		c.Context(),
		transaction,
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error updating transaction",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"transaction updated successfully",
		result,
	)
}
