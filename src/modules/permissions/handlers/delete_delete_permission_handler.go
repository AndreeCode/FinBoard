package handlers

import (
	"finboard/src/modules/permissions/domains"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *PermissionHandler) DeletePermission(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response(
			c,
			fiber.StatusBadRequest,
			"invalid permission id",
			nil,
		)
	}

	now := time.Now()

	permission := &domains.Permission{
		Id:        id,
		DeletedAt: &now,
	}

	err = h.service.DeletePermission(
		c.Context(),
		permission,
	)
	if err != nil {
		return response(
			c,
			fiber.StatusInternalServerError,
			"error deleting permission",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"permission deleted successfully",
		nil,
	)
}
