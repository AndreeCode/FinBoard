package routes

import (
	"finboard/src/core/middleware"
	"finboard/src/modules/dashboard/handlers"
	"finboard/src/modules/dashboard/repository"
	"finboard/src/modules/dashboard/services"

	"github.com/gofiber/fiber/v3"
)

func RegisterDashboardRoutes(app *fiber.App) {
	authMiddleware := middleware.NewAuthMiddleware()

	repo := repository.NewDashboardRepository()
	service := services.NewDashboardService(repo)
	handler := handlers.NewDashboardHandler(service)

	dashboard := app.Group("/dashboard")
	dashboard.Use(authMiddleware.RequireAuth())
	dashboard.Get("/", authMiddleware.RequirePermission("read_dashboard"), handler.GetDashboard)
}
