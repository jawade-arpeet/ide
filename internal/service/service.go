package service

import "ide/internal/repository"

type Service struct {
	Auth *AuthService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: newAuthService(repo.Auth),
	}
}
