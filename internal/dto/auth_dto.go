package dto

type SignInPayload struct {
	Email string `json:"email" binding:"required,email"`
}

type CreateProfilePayload struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type GetProfileResponse struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	AvatarURL *string `json:"avatar_url"`
	Email     string  `json:"email"`
}
