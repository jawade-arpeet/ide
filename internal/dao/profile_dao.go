package dao

import (
	"ide/internal/dto"

	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	AvatarURL *string   `db:"avatar_url"`
	Email     string    `db:"email"`
}

func (p *Profile) ToGetProfileResponse() *dto.GetProfileResponse {
	return &dto.GetProfileResponse{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		AvatarURL: p.AvatarURL,
		Email:     p.Email,
	}
}
