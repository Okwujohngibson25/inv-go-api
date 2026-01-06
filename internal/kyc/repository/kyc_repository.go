package repository

import (
	"context"
	"database/sql"

	"github.com/coinserveringo/internal/db"
	"github.com/coinserveringo/internal/kyc/dto"
	"github.com/google/uuid"
)

type KycRepository interface {
	Savekyc(ctx context.Context, req *dto.KYCDTO) error
	GetAllKyc(ctx context.Context) ([]dto.UserKycDTO, error)
	GetAllPendingKyc(ctx context.Context) ([]dto.UserKycDTO, error)
	GetKycById(ctx context.Context, id uuid.UUID) (*dto.UserKycDTO, error)
	Approvekyc(ctx context.Context, id uuid.UUID) error
	Declinekyc(ctx context.Context, id uuid.UUID) error
}

type SqlcRepository struct {
	queries *db.Queries
}

func NewSqlcRepository(conn *sql.DB) *SqlcRepository {
	return &SqlcRepository{
		queries: db.New(conn),
	}
}

func (r *SqlcRepository) Savekyc(ctx context.Context, req *dto.KYCDTO) error {
	params := db.SavekycParams{
		UserID:       req.UserID,
		IDImageFront: req.IDImageFront,
		IDImageBack:  req.IDImageBack,
	}
	if err := r.queries.Savekyc(ctx, params); err != nil {
		return err
	}
	return nil
}
