package routes

import (
	"finboard/src/core/middleware"
	"finboard/src/modules/transactions/handlers"
	"finboard/src/modules/transactions/repository"
	"finboard/src/modules/transactions/services"

	"github.com/gofiber/fiber/v3"
)

func RegisterTransactionRoutes(app *fiber.App) {
	authMiddleware := middleware.NewAuthMiddleware()

	repo := repository.NewTransactionRepository()
	service := services.NewTransactionService(repo)
	handler := handlers.NewTransactionHandler(service)

	transactions := app.Group("/transactions")
	transactions.Use(authMiddleware.RequireAuth())
	transactions.Post("/", authMiddleware.RequirePermission("create_transactions"), handler.CreateTransaction)
	transactions.Get("/:id", authMiddleware.RequirePermission("read_transactions"), handler.GetTransaction)
	transactions.Get("/", authMiddleware.RequirePermission("read_transactions"), handler.GetTransactions)
	transactions.Put("/:id", authMiddleware.RequirePermission("update_transactions"), handler.UpdateTransaction)
	transactions.Delete("/:id", authMiddleware.RequirePermission("delete_transactions"), handler.DeleteTransaction)
}
