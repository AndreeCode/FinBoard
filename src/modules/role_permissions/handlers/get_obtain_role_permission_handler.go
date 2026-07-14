package handlers

import (
	"errors"
	"finboard/src/modules/role_permissions/domains"
	"finboard/src/modules/role_permissions/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *RolePermissionHandler) GetRolePermission(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid role permission id",
			nil,
		)
	}

	rolePermission := &domains.RolePermission{
		Id: id,
	}

	result, err := h.service.ObtainRolePermission(
		c.Context(),
		rolePermission,
	)
	if err != nil {

		switch {

		case errors.Is(err, repository.ErrRolePermissionNotFound):

			return response(
				c,
				fiber.StatusNotFound,
				"role permission not found",
				nil,
			)

		default:

			return response(
				c,
				fiber.StatusInternalServerError,
				"error obtaining role permission",
				nil,
			)
		}
	}

	return response(
		c,
		fiber.StatusOK,
		"role permission obtained successfully",
		result,
	)
}
