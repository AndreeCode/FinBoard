package main

import (
	"context"
	"finboard/src/core/config"
	"finboard/src/core/db"
	routes "finboard/src/modules"
	"finboard/src/modules/seeder"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{

			"http://localhost:3000",
		},
		AllowCredentials: true,
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"PATCH",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
	}))

	config.LoadConfig()
	db.InitDB()

	s := seeder.NewSeeder()
	if err := s.Seed(context.Background()); err != nil {
		log.Printf("Seeder warning: %v", err)
	}

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	routes.RegisterRoutes(app)

	port := config.Config.Port
	if port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen(":" + port))

}
