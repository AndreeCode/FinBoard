package handlers

import (
	"errors"
	"finboard/src/core/middleware"
	"finboard/src/modules/categories/domains"
	"finboard/src/modules/categories/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *CategoryHandler) GetCategory(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid category id",
			nil,
		)
	}

	category := &domains.Category{
		Id: id,
	}

	result, err := h.service.ObtainCategory(
		c.Context(),
		category,
	)
	if err != nil {

		switch {

		case errors.Is(err, repository.ErrCategoryNotFound):

			return response(
				c,
				fiber.StatusNotFound,
				"category not found",
				nil,
			)

		default:

			return response(
				c,
				fiber.StatusInternalServerError,
				"error obtaining category",
				nil,
			)
		}
	}

	authMiddleware := middleware.NewAuthMiddleware()
	currentUserID := c.Locals("user_id").(string)
	if result.UserId != nil && *result.UserId != uuid.MustParse(currentUserID) && !authMiddleware.IsAdmin(c) {
		return response(
			c,
			fiber.StatusForbidden,
			"you don't have permission to view this category",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"category obtained successfully",
		result,
	)
}
