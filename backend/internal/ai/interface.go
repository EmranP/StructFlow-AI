package ai

import "context"

type Provider interface {
	GenerateStructure(
		ctx context.Context,
		prompt string,
	) (string, error)

	Name() string

	IsAvailable(
		ctx context.Context,
	) bool
}
