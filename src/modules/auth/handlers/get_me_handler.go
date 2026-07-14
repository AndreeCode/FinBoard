package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func (h *AuthHandler) Logout(c fiber.Ctx) error {
	ClearAuthCookie(c)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"msg":    "logout successful",
		"data":   nil,
	})
}

func (h *AuthHandler) Me(c fiber.Ctx) error {
	token := GetAuthToken(c)
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": fiber.StatusUnauthorized,
			"msg":    "unauthorized",
			"data":   nil,
		})
	}

	claims, err := h.service.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": fiber.StatusUnauthorized,
			"msg":    "invalid token",
			"data":   nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"msg":    "user data retrieved",
		"data": map[string]interface{}{
			"user_id":   claims.UserID,
			"name":      claims.Name,
			"last_name": claims.LastName,
			"email":     claims.Email,
			"role_id":   claims.RoleID,
		},
	})
}
