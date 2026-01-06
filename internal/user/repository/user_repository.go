package repository

import (
	"context"
	"database/sql"

	"github.com/coinserveringo/internal/db"
	"github.com/coinserveringo/internal/user/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type UserRepository interface {
	Create(ctx context.Context, data *db.CreateUserParams) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*dto.UserDataDTO, error)
	UpdateUserBalancebyId(ctx context.Context, id uuid.UUID, newBalance decimal.Decimal) error
	GetUserByEmail(ctx context.Context, email string) (*dto.UserDataDTO, error)
	CreateResetToken(ctx context.Context, req *dto.UserPasswordReset) error
	GetTokenById(ctx context.Context, token string) (*dto.UserPasswordReset, error)
	DeleteToken(ctx context.Context, id uuid.UUID) error
	UpdatePassword(ctx context.Context, id uuid.UUID, newpassword string) error
	AdminExists(ctx context.Context) (bool, error)
	GetAllUsers(ctx context.Context) ([]dto.UserDataDTO, error)
	CountUsers(ctx context.Context) (int64, error)
	UpdateAccountStatus(ctx context.Context, id uuid.UUID, newStatus string) error
	DeleteUserById(ctx context.Context, id uuid.UUID) error
	UpdateUserAcc(ctx context.Context, id uuid.UUID, User *dto.UserDataDTO) error
	UpdateKycStatus(ctx context.Context, id uuid.UUID) error
}

type SqlcRepository struct {
	queries *db.Queries
}

func NewSqlcRepository(conn *sql.DB) *SqlcRepository {
	return &SqlcRepository{
		queries: db.New(conn),
	}
}

func (r *SqlcRepository) Create(ctx context.Context, data *db.CreateUserParams) error {
	params := db.CreateUserParams{
		Email:       data.Email,
		Role:        data.Role,
		Fullname:    data.Fullname,
		Username:    data.Username,
		Gender:      data.Gender,
		Country:     data.Country,
		PhoneNumber: data.PhoneNumber,
		ReferalLink: data.ReferalLink,
		Password:    data.Password,
	}
	if err := r.queries.CreateUser(ctx, params); err != nil {
		return err
	}

	return nil
}

func (r *SqlcRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*dto.UserDataDTO, error) {
	user, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.UserDataDTO{
		ID:             user.ID,
		Role:           user.Role,
		Email:          user.Email,
		Fullname:       user.Fullname,
		Username:       user.Username,
		Gender:         user.Gender,
		Country:        user.Country,
		PhoneNumber:    user.PhoneNumber,
		WalletBalance:  user.WalletBalance,
		ProfitBalance:  user.ProfitBalance,
		InvestedAmount: user.InvestedAmount,
		ReferalBonus:   user.ReferalBonus,
		ReferalLink:    user.ReferalLink,
		Verified:       user.Verified,
		Status:         user.Status,
		CreatedAt:      user.CreatedAt,
	}, nil
}

func (r *SqlcRepository) UpdateUserBalancebyId(ctx context.Context, id uuid.UUID, newBalance decimal.Decimal) error {
	params := db.UpdateUserBalanceByIDParams{
		WalletBalance: newBalance,
		ID:            id,
	}
	if err := r.queries.UpdateUserBalanceByID(ctx, params); err != nil {
		return err
	}

	return nil
}

func (r *SqlcRepository) GetUserByEmail(ctx context.Context, email string) (*dto.UserDataDTO, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &dto.UserDataDTO{
		ID:             user.ID,
		Role:           user.Role,
		Email:          user.Email,
		Fullname:       user.Fullname,
		Username:       user.Username,
		Gender:         user.Gender,
		Country:        user.Country,
		PhoneNumber:    user.PhoneNumber,
		WalletBalance:  user.WalletBalance,
		ProfitBalance:  user.ProfitBalance,
		InvestedAmount: user.InvestedAmount,
		ReferalBonus:   user.ReferalBonus,
		ReferalLink:    user.ReferalLink,
		Verified:       user.Verified,
		Status:         user.Status,
		Password:       user.Password,
		CreatedAt:      user.CreatedAt,
	}, nil
}

func (r *SqlcRepository) CreateResetToken(ctx context.Context, req *dto.UserPasswordReset) error {
	params := db.CreateResetTokenParams{
		UserID:    req.UserID,
		Token:     req.Token,
		ExpiresAt: req.ExpiresAt,
	}
	if err := r.queries.CreateResetToken(ctx, params); err != nil {
		return err
	}
	return nil
}

func (r *SqlcRepository) GetTokenById(ctx context.Context, token string) (*dto.UserPasswordReset, error) {
	tokenData, err := r.queries.GetTokenById(ctx, token)
	if err != nil {
		return nil, err
	}

	return &dto.UserPasswordReset{
		UserID:    tokenData.UserID,
		Token:     tokenData.Token,
		TokenID:   tokenData.ID,
		ExpiresAt: tokenData.ExpiresAt,
	}, nil
}

func (r *SqlcRepository) DeleteToken(ctx context.Context, id uuid.UUID) error {
	if err := r.queries.DeleteTokenByID(ctx, id); err != nil {
		return err
	}

	return nil
}

func (r *SqlcRepository) UpdatePassword(ctx context.Context, id uuid.UUID, newpassword string) error {
	params := db.UpdateUserPasswordParams{
		Password: newpassword,
		ID:       id,
	}
	if err := r.queries.UpdateUserPassword(ctx, params); err != nil {
		return err
	}
	return nil
}
