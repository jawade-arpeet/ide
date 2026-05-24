package service

import (
	"context"
	"ide/internal/dto"
	"ide/internal/repository"
	"strings"

	"github.com/google/uuid"
)

type WorkspaceService struct {
	workspaceRepository *repository.WorkspaceRepository
}

func newWorkspaceService(repo *repository.WorkspaceRepository) *WorkspaceService {
	return &WorkspaceService{workspaceRepository: repo}
}

func (s *WorkspaceService) CreateWorkspace(ctx context.Context, accID uuid.UUID, params *dto.CreateWorkspacePayload) error {
	params.Name = strings.TrimSpace(params.Name)
	slug := strings.ToLower(params.Name)
	slug = strings.ReplaceAll(slug, " ", "-")

	if err := s.workspaceRepository.CreateWorkspace(
		ctx,
		accID,
		params.Name,
		slug,
	); err != nil {
		return err
	}

	return nil
}
