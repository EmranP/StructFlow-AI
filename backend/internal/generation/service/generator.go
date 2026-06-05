package service

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/domain"
	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/generation/status"
	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/repository"
)

type Generator struct {
	generationRepo repository.GenerationRepository

	templateRepo repository.GeneratedTemplateRepository
}

func New(
	generationRepo repository.GenerationRepository,
	templateRepo repository.GeneratedTemplateRepository,
) *Generator {

	return &Generator{
		generationRepo: generationRepo,

		templateRepo: templateRepo,
	}
}

func (g *Generator) Process(
	ctx context.Context,
	generationID uuid.UUID,
) {
	err := g.generationRepo.UpdateStatus(
		ctx,
		generationID,
		status.Processing,
		nil,
	)

	if err != nil {
		return
	}

	simple := map[string]any{
		"name": "simple",

		"directories": []string{
			"cmd",
			"internal",
			"pkg",
		},

		"files": []string{
			"main.go",
			".env",
		},
	}
	medium := map[string]any{
		"name": "medium",

		"directories": []string{
			"cmd",
			"internal",
			"pkg",
			"configs",
			"migrations",
		},
	}
	enterprise := map[string]any{
		"name": "enterprise",

		"directories": []string{
			"cmd",
			"internal",
			"pkg",
			"configs",
			"deployments",
			"api",
			"scripts",
		},
	}

	simpleContent, errSimple := json.Marshal(simple)
	mediumContent, errMedium := json.Marshal(medium)
	enterpriseContent, errEnterprise := json.Marshal(enterprise)

	if errSimple != nil || errMedium != nil || errEnterprise != nil {
		return
	}

	errTemp := g.addTemp(ctx, generationID, "simple", simpleContent)
	errTemp = g.addTemp(ctx, generationID, "medium", mediumContent)
	errTemp = g.addTemp(ctx, generationID, "enterprise", enterpriseContent)

	if errTemp != nil {
		return
	}

	err = g.generationRepo.UpdateStatus(
		ctx,
		generationID,
		status.Completed,
		nil,
	)

	if err != nil {
		return
	}
}

func (g *Generator) addTemp(ctx context.Context, genID uuid.UUID, genType string, content []byte) error {
	err := g.templateRepo.Create(
		ctx,
		&domain.GeneratedTemplate{
			ID: uuid.New(),

			GenerationID: genID,

			Type: genType,

			Content: content,
		},
	)

	return err
}
