package repository

import (
	"context"

	"github.com/coinserveringo/internal/investment/dto"
	userDataDTO "github.com/coinserveringo/internal/user/dto"
)

func (r *SqlcRepository) GetAllActiveInv(ctx context.Context) ([]dto.InvestmentResponseDTO, error) {
	var invData []dto.InvestmentResponseDTO
	allActiveInv, err := r.queries.GetAllActiveInv(ctx)
	if err != nil {
		return nil, err
	}

	for _, value := range allActiveInv {
		user := userDataDTO.UserDataDTO{
			Email:    value.Email,
			Fullname: value.Fullname,
			Username: value.Username,
		}

		invData = append(invData, dto.InvestmentResponseDTO{
			ID:             value.ID,
			User:           user,
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
