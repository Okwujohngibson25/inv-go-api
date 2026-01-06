package repository

import (
	"context"
	"database/sql"

	"github.com/coinserveringo/internal/db"
	"github.com/coinserveringo/internal/withdrawal/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type WithdrawalRepository interface {
	CreateNewWithdrawal(ctx context.Context, req *dto.UserWithdrawalDTO) error
	ApproveWithdrawal(ctx context.Context, id uuid.UUID) error
	DeclineWithdrawal(ctx context.Context, id uuid.UUID) error
	GetAllWithdrawal(ctx context.Context) ([]dto.UserWithdrawalDTO, error)
	GetAllPendingWithdrawal(ctx context.Context) ([]dto.UserWithdrawalDTO, error)
	GetWithdrawalById(ctx context.Context, id uuid.UUID) (*dto.UserWithdrawalDTO, error)
	GetAllPendingWithdrawalByID(ctx context.Context, id uuid.UUID) ([]dto.UserWithdrawalDTO, error)
	SumAllWithdrawalAmountByUserID(ctx context.Context, id uuid.UUID) (decimal.Decimal, error)
	GetAllWithdrawalByUserId(ctx context.Context, id uuid.UUID) ([]dto.UserWithdrawalDTO, error)
	SumApprovedWithdrawals(ctx context.Context) (decimal.Decimal, error)
}

type SqlcRepository struct {
	Queries *db.Queries
}

func NewSqlcRepository(conn *sql.DB) *SqlcRepository {
	return &SqlcRepository{Queries: db.New(conn)}
}

func (r *SqlcRepository) CreateNewWithdrawal(ctx context.Context, req *dto.UserWithdrawalDTO) error {
	params := db.CreateNewWithdrawalParams{
		UserID:        req.UserID,
		Amount:        req.Amount,
		TransactionID: req.TransactionID,
		CoinType:      req.CoinType,
		WalletAddress: req.WalletAddress,
	}
	if err := r.Queries.CreateNewWithdrawal(ctx, params); err != nil {
		return err
	}
	return nil
}
