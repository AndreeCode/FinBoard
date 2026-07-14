package handlers

import "github.com/gofiber/fiber/v3"

func (h *RolePermissionHandler) GetRolePermissions(c fiber.Ctx) error {

	result, err := h.service.ObtainRolePermissions(
		c.Context(),
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error obtaining role permissions",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"role permissions obtained successfully",
		result,
	)
}
