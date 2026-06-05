package dto

import (
	"encoding/json"
	"time"
)

type GenerationResponse struct {
	ID string `json:"id"`

	Status string `json:"status"`
}

type GenerationAllResponse struct {
	ID           string    `json:"id"`
	ProjectID    string    `json:"projectId"`
	Status       string    `json:"status"`
	ErrorMessage string    `json:"errorMessage"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type GenerationTempResponse struct {
	ID    string `json:"id"`
	GenID string `json:"generationId"`

	Type    string          `json:"type"`
	Content json.RawMessage `json:"content"`

	CreatedAt time.Time `json:"createdAt"`
}
