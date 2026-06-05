package service

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/ai"
	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/ai/parser"
	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/ai/prompt"
	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/ai/schema"
	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/domain"
	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/generation/status"
	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/repository"
)

type Generator struct {
	generationRepo repository.GenerationRepository
	templateRepo   repository.GeneratedTemplateRepository
	projectRepo    repository.ProjectRepository

	aiProvider ai.Provider
}

func New(
	generationRepo repository.GenerationRepository,
	templateRepo repository.GeneratedTemplateRepository,
	projectRepo repository.ProjectRepository,
	aiProvider ai.Provider,
) *Generator {

	return &Generator{
		generationRepo: generationRepo,
		templateRepo:   templateRepo,
		projectRepo:    projectRepo,
		aiProvider:     aiProvider,
	}
}

func (g *Generator) Process(
	ctx context.Context,
	generationID uuid.UUID,
	projectID uuid.UUID,
) {
	errGen := g.generationRepo.UpdateStatus(
		ctx,
		generationID,
		status.Processing,
		nil,
	)

	if errGen != nil {
		return
	}

	project, err := g.projectRepo.GetByID(
		ctx,
		projectID,
	)

	if err != nil {
		return
	}

	if project == nil {
		return
	}

	promptText := prompt.BuildProjectPrompt(
		project,
	)

	response, err := g.aiProvider.GenerateStructure(
		ctx,
		promptText,
	)

	if err != nil {
		g.fail(
			ctx,
			generationID,
			err.Error(),
		)

		return
	}

	cleaned := parser.CleanJSON(
		response,
	)

	var structures schema.ProjectStructures

	err = json.Unmarshal(
		[]byte(cleaned),
		&structures,
	)

	if err != nil {
		g.fail(
			ctx,
			generationID,
			err.Error(),
		)

		return
	}

	simpleContent, errSimple := json.Marshal(structures.Simple)
	mediumContent, errMedium := json.Marshal(structures.Medium)
	enterpriseContent, errEnterprise := json.Marshal(structures.Enterprise)

	if errSimple != nil {
		g.fail(
			ctx,
			generationID,
			errSimple.Error(),
		)

		return
	} else if errMedium != nil {
		g.fail(
			ctx,
			generationID,
			errMedium.Error(),
		)

		return
	} else if errEnterprise != nil {
		g.fail(
			ctx,
			generationID,
			errEnterprise.Error(),
		)

		return
	}

	errTemp := g.addTemp(ctx, generationID, "simple", simpleContent)
	errTemp = g.addTemp(ctx, generationID, "medium", mediumContent)
	errTemp = g.addTemp(ctx, generationID, "enterprise", enterpriseContent)

	if errTemp != nil {
		g.fail(
			ctx,
			generationID,
			errTemp.Error(),
		)

		return
	}

	err = g.generationRepo.UpdateStatus(
		ctx,
		generationID,
		status.Completed,
		nil,
	)

	if err != nil {
		g.fail(
			ctx,
			generationID,
			err.Error(),
		)

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

func (g *Generator) fail(
	ctx context.Context,
	generationID uuid.UUID,
	message string,
) {
	_ = g.generationRepo.UpdateStatus(
		ctx,
		generationID,
		status.Failed,
		&message,
	)
}
