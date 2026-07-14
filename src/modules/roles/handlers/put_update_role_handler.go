package handlers

import (
	"finboard/src/modules/roles/domains"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *RoleHandler) UpdateRole(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid role id",
			nil,
		)
	}

	var req domains.UpdateRoleRequest

	if err := c.Bind().Body(&req); err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	now := time.Now()

	role := &domains.Role{
		Id:        id,
		UpdatedAt: &now,
	}

	if req.Name != nil {
		role.Name = *req.Name
	}

	if req.Description != nil {
		role.Description = *req.Description
	}

	result, err := h.service.UpdateRole(
		c.Context(),
		role,
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error updating role",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"role updated successfully",
		result,
	)
}
