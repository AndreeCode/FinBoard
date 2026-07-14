package handlers

import (
	"errors"
	"finboard/src/core/middleware"
	"finboard/src/modules/users/domains"
	"finboard/src/modules/users/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *UserHandler) GetUser(c fiber.Ctx) error {

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

	user := &domains.User{
		Id: id,
	}

	result, err := h.service.ObtainUser(
		c.Context(),
		user,
	)
	if err != nil {

		switch {

		case errors.Is(err, repository.ErrUserNotFound):

			return response(
				c,
				fiber.StatusNotFound,
				"user not found",
				nil,
			)

		default:

			return response(
				c,
				fiber.StatusInternalServerError,
				"error obtaining user",
				nil,
			)
		}
	}

	return response(
		c,
		fiber.StatusOK,
		"user obtained successfully",
		result,
	)
}
