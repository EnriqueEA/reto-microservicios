package math

import "github.com/EnriqueEA/go-microservice/internal/domain"

func Rotate90Clockwise(m domain.Matrix) domain.Matrix {
	if len(m) == 0 || len(m[0]) == 0 {
		return domain.Matrix{}
	}

	rows := len(m)
	cols := len(m[0])
	rotated := make(domain.Matrix, cols)

	for i := 0; i < cols; i++ {
		rotated[i] = make([]float64, rows)
		for j := 0; j < rows; j++ {
			rotated[i][j] = m[rows-1-j][i]
		}
	}

	return rotated
}

func Dot(a, b []float64) float64 {
	if len(a) != len(b) {
		return 0
	}

	sum := 0.0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}
