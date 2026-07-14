package routes

import (
	"finboard/src/core/middleware"
	"finboard/src/modules/roles/handlers"
	"finboard/src/modules/roles/repository"
	"finboard/src/modules/roles/services"
	"sync"

	"github.com/gofiber/fiber/v3"
)

var sessionStore sync.Map

func RegisterRoleRoutes(app *fiber.App) {
	authMiddleware := middleware.NewAuthMiddleware()

	repo := repository.NewRoleRepository()
	service := services.NewRoleService(repo)
	handler := handlers.NewRoleHandler(service)

	roles := app.Group("/roles")
	roles.Use(authMiddleware.RequireAuth())
	roles.Post("/", authMiddleware.RequirePermission("create_roles"), handler.CreateRole)
	roles.Get("/:id", authMiddleware.RequirePermission("read_roles"), handler.GetRole)
	roles.Get("/", authMiddleware.RequirePermission("read_roles"), handler.GetRoles)
	roles.Put("/:id", authMiddleware.RequirePermission("update_roles"), handler.UpdateRole)
	roles.Delete("/:id", authMiddleware.RequirePermission("delete_roles"), handler.DeleteRole)
}
