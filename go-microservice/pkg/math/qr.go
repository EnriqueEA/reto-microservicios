package math

import (
	"math"

	"github.com/EnriqueEA/go-microservice/internal/domain"
)

func QRFactorization(A domain.Matrix) domain.QRFactorization {
	if len(A) == 0 || len(A[0]) == 0 {
		return domain.QRFactorization{
			Q: domain.Matrix{},
			R: domain.Matrix{},
		}
	}

	m := len(A)
	n := len(A[0])

	cols := matrixToColumns(A, m, n)
	Qcols := make([][]float64, 0, n)
	R := createEmptyMatrix(n, n)

	for k := 0; k < n; k++ {
		v := make([]float64, m)
		copy(v, cols[k])

		for j := 0; j < len(Qcols); j++ {
			qj := Qcols[j]
			rjk := Dot(qj, cols[k])
			R[j][k] = rjk

			for i := 0; i < m; i++ {
				v[i] -= rjk * qj[i]
			}
		}
		norm := math.Sqrt(Dot(v, v))
		if norm == 0 {
			R[k][k] = 0
			Qcols = append(Qcols, make([]float64, m))
			continue
		}

		R[k][k] = norm
		for i := 0; i < m; i++ {
			v[i] /= norm
		}
		Qcols = append(Qcols, v)
	}

	Q := columnsToMatrix(Qcols, m, n)

	return domain.QRFactorization{
		Q: Q,
		R: R,
	}
}

func matrixToColumns(A domain.Matrix, m, n int) [][]float64 {
	cols := make([][]float64, n)
	for j := 0; j < n; j++ {
		col := make([]float64, m)
		for i := 0; i < m; i++ {
			col[i] = A[i][j]
		}
		cols[j] = col
	}
	return cols
}

func columnsToMatrix(Qcols [][]float64, m, n int) domain.Matrix {
	Q := make(domain.Matrix, m)
	for i := 0; i < m; i++ {
		Q[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			Q[i][j] = Qcols[j][i]
		}
	}
	return Q
}

func createEmptyMatrix(rows, cols int) domain.Matrix {
	matrix := make(domain.Matrix, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}
	return matrix
}
