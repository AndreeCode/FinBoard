package handlers

import (
	"finboard/src/modules/permissions/services"

	"github.com/gofiber/fiber/v3"
)

type PermissionHandler struct {
	service *services.PermissionService
}

func NewPermissionHandler(service *services.PermissionService) *PermissionHandler {
	return &PermissionHandler{service: service}
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
