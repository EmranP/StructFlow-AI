package dto

import (
	"time"

	"github.com/google/uuid"
)

type ProjectResponse struct {
	ID uuid.UUID `json:"id"`

	Title string `json:"title"`

	ProjectType string `json:"projectType"`

	Stack string `json:"stack"`

	Architecture string `json:"architecture"`

	Features string `json:"features"`

	AdditionalInfo string `json:"additionalInfo"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
