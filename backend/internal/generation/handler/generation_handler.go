package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/generation/dto"
	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/generation/usecase"
)

type GenerationHandler struct {
	generationUC usecase.GenerationUseCase
}

func New(
	generationUC usecase.GenerationUseCase,
) *GenerationHandler {

	return &GenerationHandler{
		generationUC: generationUC,
	}
}

func (h *GenerationHandler) Create(
	c *fiber.Ctx,
) error {
	projectID, err := uuid.Parse(
		c.Params("id"),
	)

	if err != nil {
		return fiber.ErrBadRequest
	}

	var req dto.GenerationRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(
			fiber.StatusBadRequest,
		).JSON(
			fiber.Map{
				"error": "invalid-body model is not empty",
			},
		)
	}

	switch req.Model {

	case "gpt":
	case "claude":
	case "gemini":

	default:
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid body models: gpt, claude, gemini",
		})
	}

	generation, err := h.generationUC.
		Add(
			c.Context(),
			projectID,
			req.Model,
		)

	if err != nil {
		return err
	}

	return c.Status(
		fiber.StatusCreated,
	).JSON(
		dto.GenerationResponse{
			ID: generation.ID.String(),

			Status: generation.Status,
		},
	)
}

func (h *GenerationHandler) GetAll(
	c *fiber.Ctx,
) error {
	projectId, err := uuid.Parse(c.Params("id"))

	pageParam := c.Query("page", "1")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	limitParam := c.Query("limit", "10")
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	gens, totalCount, err := h.generationUC.FindByProjectID(
		c.Context(),
		projectId,
		limit,
		page,
	)
	if err != nil {
		return err
	}

	return c.Status(
		fiber.StatusOK,
	).JSON(
		fiber.Map{
			"generations": gens,
			"page":        page,
			"limit":       limit,
			"totalCount":  totalCount,
		},
	)
}

func (h *GenerationHandler) GetByID(
	c *fiber.Ctx,
) error {
	generationID, err := uuid.Parse(
		c.Params("id"),
	)

	if err != nil {
		return fiber.ErrBadRequest
	}

	generation, err := h.generationUC.
		FindByID(
			c.Context(),
			generationID,
		)

	if err != nil {
		return err
	}

	return c.JSON(
		dto.GenerationResponse{
			ID: generation.ID.String(),

			Status: generation.Status,
		},
	)
}

func (h *GenerationHandler) GetTemplates(c *fiber.Ctx) error {
	generationID, err := uuid.Parse(
		c.Params("id"),
	)

	if err != nil {
		return fiber.ErrBadRequest
	}

	genTemps, err := h.generationUC.FindTemplates(c.Context(), generationID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(genTemps)
}

func (h *GenerationHandler) Download(
	c *fiber.Ctx,
) error {
	id, err := uuid.Parse(
		c.Params("id"),
	)
	if err != nil {
		return err
	}

	archive, errArchive := h.generationUC.Download(
		c.Context(),
		id,
	)
	if errArchive != nil {
		return errArchive
	}

	c.Set(
		"Content-Type",
		"application/zip",
	)

	c.Set(
		"Content-Disposition",
		"attachment; filename=project.zip",
	)

	return c.Send(
		archive,
	)
}
