package handlers

import (
	"errors"
	"finboard/src/modules/role_permissions/domains"
	"finboard/src/modules/role_permissions/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *RolePermissionHandler) CreateRolePermission(c fiber.Ctx) error {

	var req domains.CreateRolePermissionRequest

	if err := c.Bind().Body(&req); err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	roleId, err := uuid.Parse(req.RoleId)
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid role_id uuid",
			nil,
		)
	}

	permissionId, err := uuid.Parse(req.PermissionId)
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid permission_id uuid",
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

	rolePermission := &domains.RolePermission{
		RoleId:       roleId,
		PermissionId: permissionId,
		CreatedBy:    createdByUUID,
	}

	result, err := h.service.CreateRolePermission(c.Context(), rolePermission)
	if err != nil {

		switch {

		case errors.Is(err, repository.ErrRolePermissionAlreadyExists):

			return response(
				c,
				fiber.StatusConflict,
				"role permission already exists",
				nil,
			)

		case errors.Is(err, repository.ErrRoleNotFoundForPermission):

			return response(
				c,
				fiber.StatusBadRequest,
				"role not found",
				nil,
			)

		case errors.Is(err, repository.ErrPermissionNotFoundForRole):

			return response(
				c,
				fiber.StatusBadRequest,
				"permission not found",
				nil,
			)

		default:

			return response(
				c,
				fiber.StatusInternalServerError,
				"error in role permission creation",
				nil,
			)
		}
	}

	return response(
		c,
		fiber.StatusCreated,
		"role permission created successfully",
		result,
	)
}
