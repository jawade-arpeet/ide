package repository

import "ide/internal/client"

type Repository struct{}

func NewRepository(client *client.Client) *Repository {
	return &Repository{}
}
