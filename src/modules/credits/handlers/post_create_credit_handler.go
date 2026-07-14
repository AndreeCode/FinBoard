package handlers

import (
	"finboard/src/modules/credits/domains"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *CreditHandler) CreateCredit(c fiber.Ctx) error {
	var req domains.CreateCreditRequest

	if err := c.Bind().Body(&req); err != nil {
		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	userId, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {
		return response(
			c,
			fiber.StatusBadRequest,
			"invalid user_id from token",
			nil,
		)
	}

	credit := &domains.Credit{
		UserId:       userId,
		PersonName:   req.PersonName,
		Amount:       req.Amount,
		InterestRate: req.InterestRate,
		IsCreditor:   req.IsCreditor,
		IsSecure:     req.IsSecure,
		Status:       req.Status,
		CreatedBy:    userId,
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
		credit.DueDate = &dueDate
	}

	result, err := h.service.CreateCredit(c.Context(), credit)
	if err != nil {
		return response(
			c,
			fiber.StatusInternalServerError,
			"error in credit creation",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusCreated,
		"credit created successfully",
		result,
	)
}
