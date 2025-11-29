package handlers

import (
	"github.com/EnriqueEA/go-microservice/internal/application/services"
	"github.com/EnriqueEA/go-microservice/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type MatrixHandler struct {
	matrixService *services.MatrixService
}

func NewMatrixHandler(matrixService *services.MatrixService) *MatrixHandler {
	return &MatrixHandler{
		matrixService: matrixService,
	}
}

func (h *MatrixHandler) Factorize(c *fiber.Ctx) error {
	var req domain.MatrixRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	result, err := h.matrixService.ProcessMatrix(req.Matrix)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}
