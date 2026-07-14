package routes

import (
	"finboard/src/modules/auth/handlers"
	"finboard/src/modules/auth/services"
	"finboard/src/modules/users/repository"

	"github.com/gofiber/fiber/v3"
)

func RegisterAuthRoutes(app *fiber.App) {
	userRepo := repository.NewUserRepository()
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	auth := app.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/logout", authHandler.Logout)
	auth.Get("/me", authHandler.Me)
}
