package repository

import "ide/internal/client"

type Repository struct {
	Auth *AuthRepository
}

func NewRepository(client *client.Client) *Repository {
	return &Repository{
		Auth: newAuthRepository(client.Postgres),
	}
}
