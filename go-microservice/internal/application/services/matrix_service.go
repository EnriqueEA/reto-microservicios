package services

import (
	"fmt"

	"github.com/EnriqueEA/go-microservice/internal/domain"
	customMath "github.com/EnriqueEA/go-microservice/pkg/math"
)

type StatsProvider interface {
	GetStatistics(Q, R, rotated domain.Matrix) (map[string]interface{}, error)
}

type MatrixService struct {
	statsProvider StatsProvider
}

func NewMatrixService(statsProvider StatsProvider) *MatrixService {
	return &MatrixService{
		statsProvider: statsProvider,
	}
}

func (s *MatrixService) ProcessMatrix(matrix domain.Matrix) (*domain.MatrixOperationResult, error) {
	if err := s.validateMatrix(matrix); err != nil {
		return nil, err
	}

	rotated := customMath.Rotate90Clockwise(matrix)
	qr := customMath.QRFactorization(rotated)

	stats, err := s.statsProvider.GetStatistics(qr.Q, qr.R, rotated)
	if err != nil {
		return nil, fmt.Errorf("failed to get statistics: %w", err)
	}

	result := &domain.MatrixOperationResult{
		Rotated:    rotated,
		Q:          qr.Q,
		R:          qr.R,
		Statistics: stats,
	}

	return result, nil
}

func (s *MatrixService) validateMatrix(matrix domain.Matrix) error {
	if len(matrix) == 0 {
		return fmt.Errorf("matrix cannot be empty")
	}
	if len(matrix[0]) == 0 {
		return fmt.Errorf("matrix rows cannot be empty")
	}
	return nil
}
