package repository

import (
	"context"

	"github.com/coinserveringo/internal/plans/dto"
)

func (r *SqlcRepository) FetchAllPlans(ctx context.Context) ([]dto.PlanResponseDTO, error) {
	var planData []dto.PlanResponseDTO
	plans, err := r.queries.FetchAllPlans(ctx)
	if err != nil {
		return nil, err
	}

	for _, value := range plans {
		planData = append(planData, dto.PlanResponseDTO{
			ID:        value.ID,
			Name:      value.Name,
			MinAmount: value.MinAmount,
			MaxAmount: value.MaxAmount,
			Duration:  int(value.DurationDays),
			Roi:       value.Roi,
			CreatedAt: value.CreatedAt,
		})
	}

	return planData, nil
}
