package repository

import (
	"context"

	"github.com/coinserveringo/internal/db"
	"github.com/coinserveringo/internal/plans/dto"
	"github.com/google/uuid"
)

func (r *SqlcRepository) CreatePlan(ctx context.Context, req *dto.PlanDTO) error {
	params := db.CreatePlanParams{
		Name:         req.Name,
		MinAmount:    req.MinAmount,
		MaxAmount:    req.MaxAmount,
		DurationDays: int32(req.Duration),
		Roi:          req.Roi,
	}
	if err := r.queries.CreatePlan(ctx, params); err != nil {
		return err
	}
	return nil
}

func (r *SqlcRepository) DeletePlanById(ctx context.Context, id uuid.UUID) error {
	if err := r.queries.DeletePlanById(ctx, id); err != nil {
		return err
	}
	return nil
}
