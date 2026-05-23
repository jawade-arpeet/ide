package repository

import (
	"context"
	"ide/internal/client"
	"ide/internal/dao"
	"ide/internal/errs"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type AuthRepository struct {
	pg *client.PostgresClient
}

func newAuthRepository(pg *client.PostgresClient) *AuthRepository {
	return &AuthRepository{pg: pg}
}

func (r *AuthRepository) CreateAccount(
	ctx context.Context,
	email string,
) (uuid.UUID, error) {
	query := `
		INSERT INTO accounts (email)
		VALUES (@email)
		RETURNING id
	`
	args := pgx.NamedArgs{"email": email}
	var accID uuid.UUID
	err := r.pg.QueryOne(ctx, query, args, &accID)
	if err != nil {
		zap.L().Error(
			"failed to create account",
			zap.Any("operation", "AuthRepository.CreateAccount"),
			zap.Error(err),
		)
		return uuid.Nil, err
	}

	zap.L().Info(
		"successfully created account",
		zap.String("operation", "AuthRepository.CreateAccount"),
	)

	return accID, nil
}

func (r *AuthRepository) CreateProfile(
	ctx context.Context,
	accID uuid.UUID,
	firstName string,
	lastName string,
) error {
	query := `
		INSERT INTO profiles (account_id, first_name, last_name)
		VALUES (@account_id, @first_name, @last_name)
	`
	args := pgx.NamedArgs{
		"account_id": accID,
		"first_name": firstName,
		"last_name":  lastName,
	}

	rowsAffected, err := r.pg.Exec(ctx, query, args)
	if err != nil {
		zap.L().Error(
			"failed to create profile",
			zap.Any("operation", "AuthRepository.CreateProfile"),
			zap.Error(err),
		)
		return err
	}

	if rowsAffected == 0 {
		return errs.ErrNoRowsAffected
	}

	zap.L().Info(
		"successfully created profile",
		zap.String("operation", "AuthRepository.CreateProfile"),
		zap.Int64("rows_affected", rowsAffected),
	)

	return nil
}

func (r *AuthRepository) GetProfile(
	ctx context.Context,
	accID uuid.UUID,
) (*dao.Profile, error) {
	query := `
		SELECT
			p.first_name,
			p.last_name,
			p.avatar_url,
			a.email
		FROM profiles p
		JOIN accounts a ON p.account_id = a.id
		WHERE p.account_id = @account_id
	`
	args := pgx.NamedArgs{"account_id": accID}

	var profile dao.Profile
	err := r.pg.QueryOne(ctx, query, args, &profile)
	if err != nil {
		zap.L().Error(
			"failed to get profile",
			zap.String("operation", "AuthRepository.GetProfile"),
			zap.Error(err),
		)
		return nil, err
	}

	zap.L().Info(
		"successfully retrieved profile",
		zap.String("operation", "AuthRepository.GetProfile"),
	)

	return &profile, nil
}
