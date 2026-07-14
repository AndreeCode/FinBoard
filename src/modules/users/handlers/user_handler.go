package handlers

import (
	"finboard/src/modules/roles/repository"
	"finboard/src/modules/users/services"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	service    *services.UserService
	roleRepo   *repository.RoleRepository
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		service:  service,
		roleRepo: repository.NewRoleRepository(),
	}
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
