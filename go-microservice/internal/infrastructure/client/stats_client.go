package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EnriqueEA/go-microservice/internal/domain"
)

type StatsClient struct {
	baseURL    string
	httpClient *http.Client
	jwtToken   string
}

func NewStatsClient(baseURL, jwtToken string) *StatsClient {
	return &StatsClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
		jwtToken:   jwtToken,
	}
}

type StatsRequest struct {
	Matrices map[string]domain.Matrix `json:"matrices"`
}

func (c *StatsClient) GetStatistics(Q, R, rotated domain.Matrix) (map[string]interface{}, error) {
	payload := StatsRequest{
		Matrices: map[string]domain.Matrix{
			"Q":       Q,
			"R":       R,
			"rotated": rotated,
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Crear request con token JWT
	req, err := http.NewRequest("POST", c.baseURL+"/stats", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.jwtToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.jwtToken)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to communicate with stats service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("stats service returned status %d", resp.StatusCode)
	}

	var stats map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("failed to parse stats response: %w", err)
	}

	return stats, nil
}
