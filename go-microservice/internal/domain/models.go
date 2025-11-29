package domain

type Matrix [][]float64

type QRFactorization struct {
	Q Matrix
	R Matrix
}

type MatrixOperationResult struct {
	Rotated    Matrix                 `json:"rotated"`
	Q          Matrix                 `json:"Q"`
	R          Matrix                 `json:"R"`
	Statistics map[string]interface{} `json:"statistics,omitempty"`
}

type MatrixRequest struct {
	Matrix Matrix `json:"matrix"`
}

type Statistics struct {
	Max      float64         `json:"max"`
	Min      float64         `json:"min"`
	Sum      float64         `json:"sum"`
	Average  float64         `json:"avg"`
	Diagonal map[string]bool `json:"diagonal"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}
