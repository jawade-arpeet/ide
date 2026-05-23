package service

import "ide/internal/repository"

type Service struct{}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
