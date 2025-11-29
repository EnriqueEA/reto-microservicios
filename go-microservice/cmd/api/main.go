package main

import (
	"log"

	"github.com/EnriqueEA/go-microservice/internal/application/services"
	"github.com/EnriqueEA/go-microservice/internal/infrastructure/client"
	"github.com/EnriqueEA/go-microservice/internal/infrastructure/config"
	"github.com/EnriqueEA/go-microservice/internal/interfaces/http/handlers"
	"github.com/EnriqueEA/go-microservice/internal/interfaces/http/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg := config.LoadConfig()

	statsClient := client.NewStatsClient(cfg.NodeServiceURL, cfg.JWTToken)
	matrixService := services.NewMatrixService(statsClient)

	healthHandler := handlers.NewHealthHandler(cfg.AppName)
	matrixHandler := handlers.NewMatrixHandler(matrixService)

	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	app.Use(cors.New())
	app.Use(logger.New())

	routes.SetupRoutes(app, healthHandler, matrixHandler)

	log.Printf("%s starting on port %s", cfg.AppName, cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
