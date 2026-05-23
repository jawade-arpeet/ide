package client

import (
	"context"
	"fmt"
	"ide/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresClient struct {
	pool *pgxpool.Pool
}

func newPostgresClient(ctx context.Context, cfg *config.PostgresConfig) (*PostgresClient, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		zap.L().Error(
			"failed to connect to postgres",
			zap.String("operation", "client.newPostgresClient"),
			zap.Error(err),
		)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		zap.L().Error(
			"failed to ping postgres",
			zap.String("operation", "client.newPostgresClient"),
			zap.Error(err),
		)
		return nil, err
	}

	zap.L().Info(
		"successfully connected to postgres",
		zap.String("operation", "client.newPostgresClient"),
	)

	return &PostgresClient{pool: pool}, nil
}
