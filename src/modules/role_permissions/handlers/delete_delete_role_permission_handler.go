package handlers

import (
	"finboard/src/modules/role_permissions/domains"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *RolePermissionHandler) DeleteRolePermission(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response(
			c,
			fiber.StatusBadRequest,
			"invalid role permission id",
			nil,
		)
	}

	now := time.Now()

	rolePermission := &domains.RolePermission{
		Id:        id,
		DeletedAt: &now,
	}

	err = h.service.DeleteRolePermission(
		c.Context(),
		rolePermission,
	)
	if err != nil {
		return response(
			c,
			fiber.StatusInternalServerError,
			"error deleting role permission",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"role permission deleted successfully",
		nil,
	)
}
