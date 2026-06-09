package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/EmranP/Design-Struct-Project-AI/backend/configs"
)

func NewPostgres(cfg *configs.Config) (*pgxpool.Pool, error) {
	var dbUrl string
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	if cfg.DBUrl == "" || cfg.ENV == "development" {
		dbUrl = dsn
	} else {
		dbUrl = cfg.DBUrl
	}

	pool, err := pgxpool.New(
		context.Background(),
		dbUrl,
	)

	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
