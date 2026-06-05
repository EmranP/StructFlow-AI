package ai

import "context"

type Provider interface {
	GenerateStructure(
		ctx context.Context,
		prompt string,
	) (string, error)
}
