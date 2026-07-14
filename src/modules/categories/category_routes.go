package routes

import (
	"finboard/src/core/middleware"
	"finboard/src/modules/categories/handlers"
	"finboard/src/modules/categories/repository"
	"finboard/src/modules/categories/services"

	"github.com/gofiber/fiber/v3"
)

func RegisterCategoryRoutes(app *fiber.App) {
	authMiddleware := middleware.NewAuthMiddleware()

	repo := repository.NewCategoryRepository()
	service := services.NewCategoryService(repo)
	handler := handlers.NewCategoryHandler(service)

	categories := app.Group("/categories")
	categories.Use(authMiddleware.RequireAuth())
	categories.Post("/", authMiddleware.RequirePermission("create_categories"), handler.CreateCategory)
	categories.Get("/:id", authMiddleware.RequirePermission("read_categories"), handler.GetCategory)
	categories.Get("/", authMiddleware.RequirePermission("read_categories"), handler.GetCategories)
	categories.Put("/:id", authMiddleware.RequirePermission("update_categories"), handler.UpdateCategory)
	categories.Delete("/:id", authMiddleware.RequirePermission("delete_categories"), handler.DeleteCategory)
}
