package routes

import (
	"finboard/src/core/middleware"
	"finboard/src/modules/users/handlers"
	"finboard/src/modules/users/repository"
	"finboard/src/modules/users/services"

	"github.com/gofiber/fiber/v3"
)

func RegisterUserRoutes(app *fiber.App) {
	authMiddleware := middleware.NewAuthMiddleware()

	repo := repository.NewUserRepository()
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	users := app.Group("/users")
	users.Post("/", handler.CreateUser)

	usersAuth := users.Group("")
	usersAuth.Use(authMiddleware.RequireAuth())
	usersAuth.Get("/:id", authMiddleware.RequirePermission("read_users"), handler.GetUser)
	usersAuth.Get("/", authMiddleware.RequirePermission("read_users"), handler.GetUsers)
	usersAuth.Put("/:id", authMiddleware.RequirePermission("update_users"), handler.UpdateUser)
	usersAuth.Delete("/:id", authMiddleware.RequirePermission("delete_users"), handler.DeleteUser)
}
