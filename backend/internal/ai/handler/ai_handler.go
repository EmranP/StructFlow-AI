package handler

import (
	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/ai"
	"github.com/gofiber/fiber/v2"
)

type AIHandler struct {
	manager *ai.Manager
}

func NewAI(
	manager *ai.Manager,
) *AIHandler {

	return &AIHandler{
		manager: manager,
	}
}

func (h *AIHandler) Models(
	c *fiber.Ctx,
) error {

	return c.JSON(
		h.manager.Models(),
	)
}
