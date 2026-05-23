package client

import (
	"context"
	"ide/internal/config"
)

type Client struct {
	Postgres *PostgresClient
}

func CreateClient(
	ctx context.Context,
	pgCfg *config.PostgresConfig,
) (*Client, error) {
	pg, err := newPostgresClient(ctx, pgCfg)
	if err != nil {
		return nil, err
	}

	return &Client{
		Postgres: pg,
	}, nil
}
