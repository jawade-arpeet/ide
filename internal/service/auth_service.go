package service

import (
	"context"
	"ide/internal/dto"
	"ide/internal/repository"
	"strings"

	"github.com/google/uuid"
)

type AuthService struct {
	authRepository *repository.AuthRepository
}

func newAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: repo,
	}
}

func (s *AuthService) SignIn(
	ctx context.Context,
	params *dto.SignInPayload,
) (string, error) {
	params.Email = strings.TrimSpace(params.Email)
	params.Email = strings.ToLower(params.Email)

	accID, err := s.authRepository.CreateAccount(ctx, params.Email)
	if err != nil {
		return "", err
	}

	return accID.String(), nil
}

func (s *AuthService) CreateProfile(
	ctx context.Context,
	accID uuid.UUID,
	params *dto.CreateProfilePayload,
) error {
	err := s.authRepository.CreateProfile(
		ctx,
		accID,
		params.FirstName,
		params.LastName,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GetProfile(
	ctx context.Context,
	accID uuid.UUID,
) (*dto.GetProfileResponse, error) {
	profile, err := s.authRepository.GetProfile(ctx, accID)
	if err != nil {
		return nil, err
	}

	profileDTO := profile.ToGetProfileResponse()

	return profileDTO, nil
}
