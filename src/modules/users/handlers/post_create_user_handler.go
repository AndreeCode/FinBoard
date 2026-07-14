package handlers

import (
	"errors"
	"finboard/src/modules/users/domains"
	"finboard/src/modules/users/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *UserHandler) CreateUser(c fiber.Ctx) error {

	var req domains.CreateUserRequest

	if err := c.Bind().Body(&req); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"msg":    "invalid request body",
			"data":   nil,
		})
	}

	var roleId uuid.UUID
	if req.RoleId == "" {
		role, err := h.roleRepo.GetRoleByName(c.Context(), "public")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"msg":    "public role not found",
				"data":   nil,
			})
		}
		roleId = role.Id
	} else {
		var err error
		roleId, err = uuid.Parse(req.RoleId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"msg":    "invalid role_id uuid",
				"data":   nil,
			})
		}
	}

	var createdByUUID uuid.UUID
	if userID := c.Locals("user_id"); userID != nil {
		parsedUUID, err := uuid.Parse(userID.(string))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"msg":    "invalid user_id from token",
				"data":   nil,
			})
		}
		createdByUUID = parsedUUID
	}

	user := &domains.User{
		Name:      req.Name,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		RoleId:    roleId,
		CreatedBy: createdByUUID,
	}

	result, err := h.service.CreateUser(c.Context(), user)
	if err != nil {

		switch {

		case errors.Is(err, repository.ErrUserAlreadyExists):

			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status": fiber.StatusConflict,
				"msg":    "user email already exists",
				"data":   nil,
			})

		case errors.Is(err, repository.ErrRoleNotFound):

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"msg":    "role not found",
				"data":   nil,
			})

		default:

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": fiber.StatusInternalServerError,
				"msg":    "error in user creation",
				"data":   nil,
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": fiber.StatusCreated,
		"msg":    "user created successfully",
		"data":   result,
	})
}
