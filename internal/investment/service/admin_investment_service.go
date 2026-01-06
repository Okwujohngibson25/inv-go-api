package service

import (
	"context"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/investment/dto"
)

func (i *InvService) FetchAllActiveinv(ctx context.Context) ([]dto.InvestmentResponseDTO, error) {
	activeInv, err := i.invRepo.GetAllActiveInv(ctx)
	if err != nil {
		log.Printf("failed to fetch all active investment: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch active investment"}
	}

	return activeInv, nil
}
