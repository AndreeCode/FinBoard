package handlers

import "github.com/gofiber/fiber/v3"

func (h *RoleHandler) GetRoles(c fiber.Ctx) error {

	result, err := h.service.ObtainRoles(
		c.Context(),
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error obtaining roles",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"roles obtained successfully",
		result,
	)
}
