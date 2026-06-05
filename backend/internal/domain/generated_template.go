package domain

import (
	"time"

	"github.com/google/uuid"
)

type GeneratedTemplate struct {
	ID uuid.UUID

	GenerationID uuid.UUID

	Type string

	Content []byte

	CreatedAt time.Time
}
