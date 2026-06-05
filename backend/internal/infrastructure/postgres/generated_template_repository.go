package postgres

import (
	"context"

	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GeneratedTemplateRepository struct {
	db *pgxpool.Pool
}

func NewGeneratedTemplateRepository(
	db *pgxpool.Pool,
) *GeneratedTemplateRepository {
	return &GeneratedTemplateRepository{
		db: db,
	}
}

func (r *GeneratedTemplateRepository) Create(
	ctx context.Context,
	template *domain.GeneratedTemplate,
) error {

	query := `
		INSERT INTO public.generated_templates(
			id,
			generation_id,
			type,
			content
		)
		VALUES(
			$1,
			$2,
			$3,
			$4
		)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		template.ID,
		template.GenerationID,
		template.Type,
		template.Content,
	)

	return err
}

func (r *GeneratedTemplateRepository) GetByGenerationID(
	ctx context.Context,
	generationID uuid.UUID,
) ([]domain.GeneratedTemplate, error) {

	query := `
		SELECT
			id,
			generation_id,
			type,
			content,
			created_at
		FROM public.generated_templates
		WHERE generation_id = $1
		ORDER BY created_at ASC
	`

	rows, err := r.db.Query(
		ctx,
		query,
		generationID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var templates []domain.GeneratedTemplate

	for rows.Next() {

		var template domain.GeneratedTemplate

		err := rows.Scan(
			&template.ID,
			&template.GenerationID,
			&template.Type,
			&template.Content,
			&template.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		templates = append(
			templates,
			template,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return templates, nil
}
