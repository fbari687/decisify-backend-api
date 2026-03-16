package server

import (
	"decisify-backend-api/internal/handler"
	"decisify-backend-api/internal/repository"
	"decisify-backend-api/internal/routes"
	"decisify-backend-api/internal/service"
	"decisify-backend-api/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"log"
	"time"
)

func Execute() {
	cfg := config.LoadConfig()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.Cors,
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
	}))

	aiRepo := repository.NewAIRepository(cfg.OpenAIKey)
	aiService := service.NewAIService(aiRepo)
	aiHandler := handler.NewAIHandler(aiService)

	routes.SetupRoutes(app, aiHandler)

	err := app.Listen(":" + cfg.Port)
	if err != nil {
		log.Fatalf("Error starting server %v", err)
	}
}
