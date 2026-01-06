package repository

import (
	"context"

	"github.com/coinserveringo/internal/investment/dto"
	"github.com/google/uuid"
)

func (r *SqlcRepository) GetActiveInvByUserId(ctx context.Context, id uuid.UUID) ([]dto.InvestmentResponseDTO, error) {
	var invData []dto.InvestmentResponseDTO
	activeInv, err := r.queries.GetActiveInvByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, value := range activeInv {
		invData = append(invData, dto.InvestmentResponseDTO{
			ID:             value.ID,
			PlanID:         value.PlanID,
			Amount:         value.Amount,
			StartDate:      value.StartDate,
			EndDate:        value.EndDate,
			ExpectedProfit: value.ExpectedProfit,
			Status:         value.Status,
			CreatedAt:      value.CreatedAt,
		})
	}

	return invData, nil
}

func (r *SqlcRepository) GetAllInvByUserId(ctx context.Context, id uuid.UUID) ([]dto.InvestmentResponseDTO, error) {
	var invData []dto.InvestmentResponseDTO
	allInv, err := r.queries.GetAllInvByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, value := range allInv {
		invData = append(invData, dto.InvestmentResponseDTO{
			ID:             value.ID,
			PlanID:         value.PlanID,
			Amount:         value.Amount,
			StartDate:      value.StartDate,
			EndDate:        value.EndDate,
			ExpectedProfit: value.ExpectedProfit,
			Status:         value.Status,
			CreatedAt:      value.CreatedAt,
		})
	}

	return invData, nil
}
