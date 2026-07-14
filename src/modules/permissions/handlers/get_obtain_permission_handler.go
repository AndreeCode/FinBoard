package handlers

import (
	"errors"
	"finboard/src/modules/permissions/domains"
	"finboard/src/modules/permissions/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *PermissionHandler) GetPermission(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid permission id",
			nil,
		)
	}

	permission := &domains.Permission{
		Id: id,
	}

	result, err := h.service.ObtainPermission(
		c.Context(),
		permission,
	)
	if err != nil {

		switch {

		case errors.Is(err, repository.ErrPermissionNotFound):

			return response(
				c,
				fiber.StatusNotFound,
				"permission not found",
				nil,
			)

		default:

			return response(
				c,
				fiber.StatusInternalServerError,
				"error obtaining permission",
				nil,
			)
		}
	}

	return response(
		c,
		fiber.StatusOK,
		"permission obtained successfully",
		result,
	)
}
