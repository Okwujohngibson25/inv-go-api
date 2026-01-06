package repository

import (
	"context"
	"database/sql"

	"github.com/coinserveringo/internal/db"
	"github.com/coinserveringo/internal/plans/dto"
	"github.com/google/uuid"
)

type PlanRepository interface {
	CreatePlan(ctx context.Context, req *dto.PlanDTO) error
	GetPlanByID(ctx context.Context, id uuid.UUID) (*dto.PlanResponseDTO, error)
	DeletePlanById(ctx context.Context, id uuid.UUID) error
	FetchAllPlans(ctx context.Context) ([]dto.PlanResponseDTO, error)
}

type SqlcRepository struct {
	queries *db.Queries
}

func NewSqlcRepository(conn *sql.DB) *SqlcRepository {
	return &SqlcRepository{queries: db.New(conn)}
}

func (r *SqlcRepository) GetPlanByID(ctx context.Context, id uuid.UUID) (*dto.PlanResponseDTO, error) {
	plan, err := r.queries.GetPlanByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.PlanResponseDTO{
		ID:        plan.ID,
		Name:      plan.Name,
		MinAmount: plan.MinAmount,
		MaxAmount: plan.MaxAmount,
		Duration:  int(plan.DurationDays),
		Roi:       plan.Roi,
		CreatedAt: plan.CreatedAt,
	}, nil
}
