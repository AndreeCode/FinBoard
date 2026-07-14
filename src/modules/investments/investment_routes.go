package routes

import (
	"finboard/src/core/middleware"
	"finboard/src/modules/investments/handlers"
	"finboard/src/modules/investments/repository"
	"finboard/src/modules/investments/services"

	"github.com/gofiber/fiber/v3"
)

func RegisterInvestmentRoutes(app *fiber.App) {
	authMiddleware := middleware.NewAuthMiddleware()

	repo := repository.NewInvestmentRepository()
	service := services.NewInvestmentService(repo)
	handler := handlers.NewInvestmentHandler(service)

	investments := app.Group("/investments")
	investments.Use(authMiddleware.RequireAuth())
	investments.Post("/", authMiddleware.RequirePermission("create_investments"), handler.CreateInvestment)
	investments.Get("/:id", authMiddleware.RequirePermission("read_investments"), handler.GetInvestment)
	investments.Get("/", authMiddleware.RequirePermission("read_investments"), handler.GetInvestments)
	investments.Put("/:id", authMiddleware.RequirePermission("update_investments"), handler.UpdateInvestment)
	investments.Delete("/:id", authMiddleware.RequirePermission("delete_investments"), handler.DeleteInvestment)
}
