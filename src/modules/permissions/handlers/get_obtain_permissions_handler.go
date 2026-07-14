package handlers

import "github.com/gofiber/fiber/v3"

func (h *PermissionHandler) GetPermissions(c fiber.Ctx) error {

	result, err := h.service.ObtainPermissions(
		c.Context(),
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error obtaining permissions",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"permissions obtained successfully",
		result,
	)
}
