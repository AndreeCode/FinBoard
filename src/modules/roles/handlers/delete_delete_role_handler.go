package handlers

import (
	"finboard/src/modules/roles/domains"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *RoleHandler) DeleteRole(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response(
			c,
			fiber.StatusBadRequest,
			"invalid role id",
			nil,
		)
	}

	now := time.Now()

	role := &domains.Role{
		Id:        id,
		DeletedAt: &now,
	}

	err = h.service.DeleteRole(
		c.Context(),
		role,
	)
	if err != nil {
		return response(
			c,
			fiber.StatusInternalServerError,
			"error deleting role",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"role deleted successfully",
		nil,
	)
}
