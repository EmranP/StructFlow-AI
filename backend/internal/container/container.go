package container

import (
	"github.com/EmranP/Design-Struct-Project-AI/backend/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Container struct {
	DB     *pgxpool.Pool
	Logger *zap.Logger

	UserRepository    repository.UserRepository
	ProjectRepository repository.ProjectRepository

	GenerationRepository        repository.GenerationRepository
	GeneratedTemplateRepository repository.GeneratedTemplateRepository
}
