package routes

import (
	"finboard/src/core/middleware"
	"finboard/src/modules/permissions/handlers"
	"finboard/src/modules/permissions/repository"
	"finboard/src/modules/permissions/services"

	"github.com/gofiber/fiber/v3"
)

func RegisterPermissionRoutes(app *fiber.App) {
	authMiddleware := middleware.NewAuthMiddleware()

	repo := repository.NewPermissionRepository()
	service := services.NewPermissionService(repo)
	handler := handlers.NewPermissionHandler(service)

	permissions := app.Group("/permissions")
	permissions.Use(authMiddleware.RequireAuth())
	permissions.Post("/", authMiddleware.RequirePermission("create_permissions"), handler.CreatePermission)
	permissions.Get("/:id", authMiddleware.RequirePermission("read_permissions"), handler.GetPermission)
	permissions.Get("/", authMiddleware.RequirePermission("read_permissions"), handler.GetPermissions)
	permissions.Put("/:id", authMiddleware.RequirePermission("update_permissions"), handler.UpdatePermission)
	permissions.Delete("/:id", authMiddleware.RequirePermission("delete_permissions"), handler.DeletePermission)
}
