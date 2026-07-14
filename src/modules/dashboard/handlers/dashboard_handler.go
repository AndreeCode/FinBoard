package handlers

import (
	"finboard/src/modules/dashboard/services"

	"github.com/gofiber/fiber/v3"
)

type DashboardHandler struct {
	service *services.DashboardService
}

func NewDashboardHandler(service *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: service}
}

func (h *DashboardHandler) GetDashboard(c fiber.Ctx) error {
	userId := c.Locals("user_id")
	if userId == nil {
		return response(c, fiber.StatusUnauthorized, "unauthorized", nil)
	}

	period := c.Query("period", "monthly")
	if period == "" {
		period = "monthly"
	}

	result, err := h.service.GetDashboard(c.Context(), userId.(string), period)
	if err != nil {
		return response(c, fiber.StatusInternalServerError, "error obtaining dashboard", nil)
	}

	return response(c, fiber.StatusOK, "dashboard obtained successfully", result)
}

func response(c fiber.Ctx, status int, msg string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status": status,
		"msg":    msg,
		"data":   data,
	})
}
