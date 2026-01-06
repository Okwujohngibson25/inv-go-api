package repository

import (
	"context"
	"database/sql"

	"github.com/coinserveringo/internal/db"
	"github.com/coinserveringo/internal/investment/dto"
	"github.com/google/uuid"
)

type InvRepository interface {
	CreateInv(ctx context.Context, req *dto.InvestmentDTO) error
	GetAllActiveInv(ctx context.Context) ([]dto.InvestmentResponseDTO, error)
	GetActiveInvByUserId(ctx context.Context, id uuid.UUID) ([]dto.InvestmentResponseDTO, error)
	GetAllInvByUserId(ctx context.Context, id uuid.UUID) ([]dto.InvestmentResponseDTO, error)
}

type SqlcRepository struct {
	queries *db.Queries
}

func NewSqlcRepository(conn *sql.DB) *SqlcRepository {
	return &SqlcRepository{
		queries: db.New(conn),
	}
}

func (r *SqlcRepository) CreateInv(ctx context.Context, req *dto.InvestmentDTO) error {
	params := db.CreateInvParams{
		UserID:         req.UserID,
		PlanID:         req.PlanID,
		Amount:         req.Amount,
		EndDate:        req.EndDate,
		ExpectedProfit: req.ExpectedProfit,
	}
	if err := r.queries.CreateInv(ctx, params); err != nil {
		return err
	}

	return nil
}
