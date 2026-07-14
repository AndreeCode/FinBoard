package handlers

import (
	"finboard/src/modules/transactions/domains"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"time"
)

func (h *TransactionHandler) CreateTransaction(c fiber.Ctx) error {

	var req domains.CreateTransactionRequest

	if err := c.Bind().Body(&req); err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	userIdStr := c.Locals("user_id").(string)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid user_id from token",
			nil,
		)
	}

	transactionDate, err := time.Parse(time.RFC3339, req.TransactionDate)
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid transaction_date format",
			nil,
		)
	}

	transaction := &domains.Transaction{
		UserId:          userId,
		Amount:          req.Amount,
		Type:            req.Type,
		TransactionDate: transactionDate,
		Description:     req.Description,
		CreatedBy:       userId,
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

	result, err := h.service.CreateTransaction(c.Context(), transaction)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error in transaction creation",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusCreated,
		"transaction created successfully",
		result,
	)
}
