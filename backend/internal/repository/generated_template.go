package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/domain"
)

type GeneratedTemplateRepository interface {
	Create(
		ctx context.Context,
		template *domain.GeneratedTemplate,
	) error

	GetByGenerationID(
		ctx context.Context,
		generationID uuid.UUID,
	) ([]domain.GeneratedTemplate, error)
}
