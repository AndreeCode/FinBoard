package handlers

import "github.com/gofiber/fiber/v3"

func (h *UserHandler) GetUsers(c fiber.Ctx) error {

	result, err := h.service.ObtainUsers(
		c.Context(),
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error obtaining users",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"users obtained successfully",
		result,
	)
}
