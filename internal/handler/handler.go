package handler

import "ide/internal/service"

type Handler struct {
	Health    *HealthHandler
	Auth      *AuthHandler
	Workspace *WorkspaceHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Health:    newHealthHandler(),
		Auth:      newAuthHandler(service.Auth),
		Workspace: newWorkspaceHandler(service.Workspace),
	}
}
