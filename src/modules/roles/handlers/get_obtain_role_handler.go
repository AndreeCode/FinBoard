package handlers

import (
	"errors"
	"finboard/src/modules/roles/domains"
	"finboard/src/modules/roles/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *RoleHandler) GetRole(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid role id",
			nil,
		)
	}

	role := &domains.Role{
		Id: id,
	}

	result, err := h.service.ObtainRole(
		c.Context(),
		role,
	)
	if err != nil {

		switch {

		case errors.Is(err, repository.ErrRoleAlreadyExists):

			return response(
				c,
				fiber.StatusNotFound,
				"role not found",
				nil,
			)

		default:

			return response(
				c,
				fiber.StatusInternalServerError,
				"error obtaining role",
				nil,
			)
		}
	}

	return response(
		c,
		fiber.StatusOK,
		"role obtained successfully",
		result,
	)
}
