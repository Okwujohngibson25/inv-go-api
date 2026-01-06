package repository

import (
	"context"
	"database/sql"

	"github.com/coinserveringo/internal/db"
	"github.com/coinserveringo/internal/deposit/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type DepositRepository interface {
	// User
	CreateNewDeposit(ctx context.Context, req *dto.DepositDTO) error

	// Admin
	GetAllDeposits(ctx context.Context) ([]dto.UserDepositDTO, error)
	GetAllPendingDeposit(ctx context.Context) ([]dto.UserDepositDTO, error)
	GetDepositByID(ctx context.Context, id uuid.UUID) (*dto.UserDepositDTO, error)
	ApproveDeposits(ctx context.Context, id uuid.UUID) error
	DeclineDeposit(ctx context.Context, id uuid.UUID) error

	// Shared
	GetAllDepositsByUserId(ctx context.Context, id uuid.UUID) ([]dto.UserDepositDTO, error)
	GetAllPendingDepositByUserID(ctx context.Context, id uuid.UUID) ([]dto.UserDepositDTO, error)
	SumAllDepositAmountByUserID(ctx context.Context, id uuid.UUID) (decimal.Decimal, error)
	SumApprovedDeposits(ctx context.Context) (decimal.Decimal, error)
}

type SqlcRepository struct {
	queries *db.Queries
}

func NewSqlcRepository(conn *sql.DB) *SqlcRepository {
	return &SqlcRepository{
		queries: db.New(conn),
	}
}

func (r *SqlcRepository) CreateNewDeposit(ctx context.Context, req *dto.DepositDTO) error {
	params := db.CreateNewDepositParams{
		UserID:        req.UserID,
		Amount:        req.Amount,
		TransactionID: req.TransactionID,
		DepositProof:  req.DepositProof,
	}
	if err := r.queries.CreateNewDeposit(ctx, params); err != nil {
		return err
	}

	return nil
}
