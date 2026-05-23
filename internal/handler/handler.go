package handler

import "ide/internal/service"

type Handler struct {
	Health *HealthHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Health: newHealthHandler(),
	}
}
