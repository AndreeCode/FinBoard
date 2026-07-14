package handlers

import (
	"finboard/src/modules/categories/domains"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *CategoryHandler) CreateCategory(c fiber.Ctx) error {

	var req domains.CreateCategoryRequest

	if err := c.Bind().Body(&req); err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	createdByUUID, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid user_id from token",
			nil,
		)
	}

	category := &domains.Category{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   createdByUUID,
	}

	if req.ParentId != nil {
		parentId, err := uuid.Parse(*req.ParentId)
		if err != nil {

			return response(
				c,
				fiber.StatusBadRequest,
				"invalid parent_id uuid",
				nil,
			)
		}
		category.ParentId = &parentId
	}

	userId := c.Locals("user_id").(string)
	userIdUUID, err := uuid.Parse(userId)
	if err != nil {

		return response(
			c,
			fiber.StatusBadRequest,
			"invalid user_id from token",
			nil,
		)
	}
	category.UserId = &userIdUUID

	result, err := h.service.CreateCategory(c.Context(), category)
	if err != nil {

		return response(
			c,
			fiber.StatusInternalServerError,
			"error in category creation",
			nil,
		)
	}

	return response(
		c,
		fiber.StatusCreated,
		"category created successfully",
		result,
	)
}
