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

func (h *CategoryHandler) UpdateCategory(c fiber.Ctx) error {

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
		return response(c, fiber.StatusForbidden, "you don't have permission to update this category", nil)
	}

	var req domains.UpdateCategoryRequest

	if err := c.Bind().Body(&req); err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	now := time.Now()

	category := &domains.Category{
		Id:        id,
		UpdatedAt: &now,
	}

	if req.Name != nil {
		category.Name = *req.Name
	}

	if req.Description != nil {
		category.Description = *req.Description
	}

	result, err := h.service.UpdateCategory(
		c.Context(),
		category,
	)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error updating category",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusOK,
		"category updated successfully",
		result,
	)
}
