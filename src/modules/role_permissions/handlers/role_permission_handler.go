package handlers

import (
	"finboard/src/modules/role_permissions/services"

	"github.com/gofiber/fiber/v3"
)

type RolePermissionHandler struct {
	service *services.RolePermissionService
}

func NewRolePermissionHandler(service *services.RolePermissionService) *RolePermissionHandler {
	return &RolePermissionHandler{service: service}
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
