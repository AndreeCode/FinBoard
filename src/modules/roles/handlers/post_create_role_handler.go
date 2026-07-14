package handlers

import (
	"errors"
	"finboard/src/modules/roles/domains"
	"finboard/src/modules/roles/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *RoleHandler) CreateRole(c fiber.Ctx) error {

	var req domains.CreateRoleRequest

	if err := c.Bind().Body(&req); err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	createdByUUID, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid user_id from token",
			nil,
		)
	}

	role := &domains.Role{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   createdByUUID,
	}

	result, err := h.service.CreateRole(c.Context(), role)
	if err != nil {

		switch {

		case errors.Is(err, repository.ErrRoleAlreadyExists):

			return response(
				c,
				fiber.StatusConflict,
				"role already exists",
				nil,
			)

		case errors.Is(err, repository.ErrRoleDescriptionExists):

			return response(
				c,
				fiber.StatusConflict,
				"description already exists",
				nil,
			)

		default:

			return response(
				c,
				fiber.StatusInternalServerError,
				"error in role creation",
				nil,
			)
		}
	}

	return response(
		c,
		fiber.StatusCreated,
		"role created successfully",
		result,
	)
}
