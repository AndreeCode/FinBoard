package handlers

import (
	"finboard/src/modules/permissions/domains"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *PermissionHandler) UpdatePermission(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid permission id",
			nil,
		)
	}

	var req domains.UpdatePermissionRequest

	if err := c.Bind().Body(&req); err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	now := time.Now()

	permission := &domains.Permission{
		Id:        id,
		UpdatedAt: &now,
	}

	if req.Name != nil {
		permission.Name = *req.Name
	}

	if req.Description != nil {
		permission.Description = *req.Description
	}

	result, err := h.service.UpdatePermission(
		c.Context(),
		permission,
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error updating permission",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"permission updated successfully",
		result,
	)
}
