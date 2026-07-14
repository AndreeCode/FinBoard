package handlers

import (
	"errors"
	"finboard/src/modules/permissions/domains"
	"finboard/src/modules/permissions/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *PermissionHandler) CreatePermission(c fiber.Ctx) error {

	var req domains.CreatePermissionRequest

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

	permission := &domains.Permission{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   createdByUUID,
	}

	result, err := h.service.CreatePermission(c.Context(), permission)
	if err != nil {

		switch {

		case errors.Is(err, repository.ErrPermissionAlreadyExists):

			return response(
				c,
				fiber.StatusConflict,
				"permission already exists",
				nil,
			)

		default:

			return response(
				c,
				fiber.StatusInternalServerError,
				"error in permission creation",
				nil,
			)
		}
	}

	return response(
		c,
		fiber.StatusCreated,
		"permission created successfully",
		result,
	)
}
