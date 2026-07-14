package handlers

import (
	"errors"
	"finboard/src/core/middleware"
	"finboard/src/modules/categories/domains"
	"finboard/src/modules/categories/repository"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *CategoryHandler) DeleteCategory(c fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response(
			c,
			fiber.StatusBadRequest,
			"invalid category id",
			nil,
		)
	}

	existingCategory := &domains.Category{Id: id}
	existingCategory, err = h.service.ObtainCategory(c.Context(), existingCategory)
	if err != nil {
		if errors.Is(err, repository.ErrCategoryNotFound) {
			return response(c, fiber.StatusNotFound, "category not found", nil)
		}
		return response(c, fiber.StatusInternalServerError, "error obtaining category", nil)
	}

	authMiddleware := middleware.NewAuthMiddleware()
	currentUserID := c.Locals("user_id").(string)
	if existingCategory.UserId != nil && *existingCategory.UserId != uuid.MustParse(currentUserID) && !authMiddleware.IsAdmin(c) {
		return response(c, fiber.StatusForbidden, "you don't have permission to delete this category", nil)
	}

	now := time.Now()

	category := &domains.Category{
		Id:        id,
		DeletedAt: &now,
	}

	err = h.service.DeleteCategory(
		c.Context(),
		category,
	)
	if err != nil {
		return response(
			c,
			fiber.StatusInternalServerError,
			"error deleting category",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"category deleted successfully",
		nil,
	)
}
