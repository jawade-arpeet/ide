package main

import (
	"context"
	"ide/internal/client"
	"ide/internal/config"
	"ide/internal/logger"
	"ide/internal/server"

	"go.uber.org/zap"
)

func main() {
	logger := logger.NewLogger()
	zap.ReplaceGlobals(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		zap.L().Fatal(
			"failed to load config",
			zap.String("action", "shutdown"),
			zap.Error(err),
		)
	}

	client, err := client.CreateClient(
		context.Background(),
		cfg.Postgres,
	)
	if err != nil {
		zap.L().Fatal(
			"failed to create client",
			zap.String("action", "shutdown"),
			zap.Error(err),
		)
	}

	if err := server.NewServer(
		cfg.Server,
		client,
	).Start(); err != nil {
		zap.L().Fatal(
			"failed to start the server",
			zap.String("action", "shutdown"),
			zap.Error(err),
		)
	}
}
