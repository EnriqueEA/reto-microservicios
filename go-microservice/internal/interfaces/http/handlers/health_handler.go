package handlers

import (
	"github.com/EnriqueEA/go-microservice/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	serviceName string
}

func NewHealthHandler(serviceName string) *HealthHandler {
	return &HealthHandler{
		serviceName: serviceName,
	}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	response := domain.HealthResponse{
		Status:  "ok",
		Service: h.serviceName,
	}
	return c.JSON(response)
}
