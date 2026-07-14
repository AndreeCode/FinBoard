package routes

import (
	"finboard/src/core/middleware"
	"finboard/src/modules/credits/handlers"
	"finboard/src/modules/credits/repository"
	"finboard/src/modules/credits/services"

	"github.com/gofiber/fiber/v3"
)

func RegisterCreditRoutes(app *fiber.App) {
	authMiddleware := middleware.NewAuthMiddleware()

	repo := repository.NewCreditRepository()
	service := services.NewCreditService(repo)
	handler := handlers.NewCreditHandler(service)

	credits := app.Group("/credits")
	credits.Use(authMiddleware.RequireAuth())
	credits.Post("/", authMiddleware.RequirePermission("create_credits"), handler.CreateCredit)
	credits.Get("/:id", authMiddleware.RequirePermission("read_credits"), handler.GetCredit)
	credits.Get("/", authMiddleware.RequirePermission("read_credits"), handler.GetCredits)
	credits.Put("/:id", authMiddleware.RequirePermission("update_credits"), handler.UpdateCredit)
	credits.Delete("/:id", authMiddleware.RequirePermission("delete_credits"), handler.DeleteCredit)
}
