package domain

import (
	"time"

	"github.com/google/uuid"
)

type Generation struct {
	ID uuid.UUID

	ProjectID uuid.UUID

	Status string

	ErrorMessage *string

	CreatedAt time.Time

	UpdatedAt time.Time
}
