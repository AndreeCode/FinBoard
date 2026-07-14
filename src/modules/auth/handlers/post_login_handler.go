package handlers

import (
	"errors"
	"finboard/src/modules/auth/domains"
	"finboard/src/modules/auth/services"

	"github.com/gofiber/fiber/v3"
)

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req domains.LoginRequest

	if err := c.Bind().Body(&req); err != nil {
		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	if req.Email == "" || req.Password == "" {
		return response(
			c,
			fiber.StatusBadRequest,
			"email and password are required",
			nil,
		)
	}

	result, err := h.service.Login(c.Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrUserNotFound):
			return response(
				c,
				fiber.StatusUnauthorized,
				"user not found",
				nil,
			)
		case errors.Is(err, services.ErrInvalidCredentials):
			return response(
				c,
				fiber.StatusUnauthorized,
				"invalid credentials",
				nil,
			)
		default:
			return response(
				c,
				fiber.StatusInternalServerError,
				"error during login",
				nil,
			)
		}
	}

	SetAuthCookie(c, result.Token)

	return response(
		c,
		fiber.StatusOK,
		"login successful",
		fiber.Map{
			"user": result.User,
		},
	)
}
