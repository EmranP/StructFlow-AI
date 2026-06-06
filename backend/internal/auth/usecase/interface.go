package usecase

import (
	"context"

	"github.com/google/uuid"
)

type AuthUseCase interface {
	Register(
		ctx context.Context,
		email string,
		password string,
	) error

	Login(
		ctx context.Context,
		email string,
		password string,
	) (string, error)

	Me(
		ctx context.Context,
		id uuid.UUID,
	) (string, error)
}
