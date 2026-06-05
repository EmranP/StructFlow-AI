package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/domain"
)

type GenerationRepository interface {
	Create(
		ctx context.Context,
		generation *domain.Generation,
	) error

	GetAllByProjectID(
		ctx context.Context,
		projectID uuid.UUID,
		limit,
		offset int,
	) ([]domain.Generation, error)

	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*domain.Generation, error)

	GetCount(
		ctx context.Context,
		projectID uuid.UUID,
	) (int64, error)

	UpdateStatus(
		ctx context.Context,
		id uuid.UUID,
		status string,
		errorMessage *string,
	) error
}
