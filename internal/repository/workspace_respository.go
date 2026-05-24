package repository

import (
	"context"
	"ide/internal/client"
	"ide/internal/errs"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type WorkspaceRepository struct {
	pg *client.PostgresClient
}

func newWorkspaceRepository(pg *client.PostgresClient) *WorkspaceRepository {
	return &WorkspaceRepository{pg: pg}
}

func (r *WorkspaceRepository) CreateWorkspace(
	ctx context.Context,
	accountID uuid.UUID,
	name string,
	slug string,
) error {
	query := `
		INSERT INTO workspaces (created_by, name, slug)
		VALUES (@account_id, @name, @slug)
	`
	args := pgx.NamedArgs{
		"account_id": accountID,
		"name":       name,
		"slug":       slug,
	}

	rowsAff, err := r.pg.Exec(ctx, query, args)
	if err != nil {
		zap.L().Error(
			"failed to create workspace",
			zap.String("operation", "workspace.CreateWorkspace"),
			zap.Any("account_id", accountID),
			zap.String("name", name),
			zap.String("slug", slug),
			zap.Error(err),
		)
		return err
	}

	if rowsAff != 1 {
		zap.L().Error(
			"expected 1 row affected, got",
			zap.String("operation", "workspace.CreateWorkspace"),
			zap.Any("account_id", accountID),
			zap.String("name", name),
			zap.String("slug", slug),
		)
		return errs.ErrNoRowsAffected
	}

	return nil
}
