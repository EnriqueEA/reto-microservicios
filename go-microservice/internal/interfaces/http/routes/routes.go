package routes

import (
	"github.com/EnriqueEA/go-microservice/internal/interfaces/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	healthHandler *handlers.HealthHandler,
	matrixHandler *handlers.MatrixHandler,
) {
	app.Get("/health", healthHandler.Check)
	app.Post("/factorize", matrixHandler.Factorize)
}
