package handlers

import (
	"finboard/src/core/middleware"

	"github.com/gofiber/fiber/v3"
)

func (h *TransactionHandler) GetTransactions(c fiber.Ctx) error {
	categoryId := c.Query("category_id")

	userId := ""
	authMiddleware := middleware.NewAuthMiddleware()
	if !authMiddleware.IsAdmin(c) {
		userId = c.Locals("user_id").(string)
	}

	result, err := h.service.ObtainTransactions(
		c.Context(),
		categoryId,
		userId,
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error obtaining transactions",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"transactions obtained successfully",
		result,
	)
}
