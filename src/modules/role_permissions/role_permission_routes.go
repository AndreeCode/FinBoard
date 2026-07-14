package routes

import (
	"finboard/src/core/middleware"
	"finboard/src/modules/role_permissions/handlers"
	"finboard/src/modules/role_permissions/repository"
	"finboard/src/modules/role_permissions/services"

	"github.com/gofiber/fiber/v3"
)

func RegisterRolePermissionRoutes(app *fiber.App) {
	authMiddleware := middleware.NewAuthMiddleware()

	repo := repository.NewRolePermissionRepository()
	service := services.NewRolePermissionService(repo)
	handler := handlers.NewRolePermissionHandler(service)

	rolePermissions := app.Group("/role-permissions")
	rolePermissions.Use(authMiddleware.RequireAuth())
	rolePermissions.Post("/", authMiddleware.RequirePermission("create_role_permissions"), handler.CreateRolePermission)
	rolePermissions.Get("/:id", authMiddleware.RequirePermission("read_role_permissions"), handler.GetRolePermission)
	rolePermissions.Get("/", authMiddleware.RequirePermission("read_role_permissions"), handler.GetRolePermissions)
	rolePermissions.Delete("/:id", authMiddleware.RequirePermission("delete_role_permissions"), handler.DeleteRolePermission)
}
