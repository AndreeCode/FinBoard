package handlers

import (
	"finboard/src/modules/roles/services"

	"github.com/gofiber/fiber/v3"
)

type RoleHandler struct {
	service *services.RoleService
}

func NewRoleHandler(service *services.RoleService) *RoleHandler {
	return &RoleHandler{service: service}
}

func response(
	c fiber.Ctx,
	status int,
	msg string,
	data interface{},
) error {
	return c.Status(status).JSON(fiber.Map{
		"status": status,
		"msg":    msg,
		"data":   data,
	})
}
