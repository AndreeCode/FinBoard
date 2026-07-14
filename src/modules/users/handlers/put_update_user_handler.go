package handlers

import (
	"finboard/src/core/middleware"
	"finboard/src/modules/users/domains"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *UserHandler) UpdateUser(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid user id",
			nil,
		)
	}

	authMiddleware := middleware.NewAuthMiddleware()
	if !authMiddleware.CanAccessResource(c, id.String()) {
		return response(
			c,
			fiber.StatusForbidden,
			"access denied",
			nil,
		)
	}

	var req domains.UpdateUserRequest

	if err := c.Bind().Body(&req); err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	now := time.Now()

	user := &domains.User{
		Id:        id,
		UpdatedAt: &now,
	}

	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.LastName != nil {
		user.LastName = *req.LastName
	}

	result, err := h.service.UpdateUser(
		c.Context(),
		user,
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error updating user",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"user updated successfully",
		result,
	)
}
