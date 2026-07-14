package handlers

import (
	"finboard/src/core/middleware"
	"finboard/src/modules/users/domains"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *UserHandler) DeleteUser(c fiber.Ctx) error {

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

	now := time.Now()

	user := &domains.User{
		Id:        id,
		DeletedAt: &now,
	}

	err = h.service.DeleteUser(
		c.Context(),
		user,
	)
	if err != nil {
		return response(
			c,
			fiber.StatusInternalServerError,
			"error deleting user",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"user deleted successfully",
		nil,
	)
}
