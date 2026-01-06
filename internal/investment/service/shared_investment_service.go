package service

import (
	"context"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/investment/dto"
	"github.com/google/uuid"
)

func (i *InvService) FetchActiveInvByUserId(ctx context.Context, id uuid.UUID) ([]dto.InvestmentResponseDTO, error) {
	inv, err := i.invRepo.GetActiveInvByUserId(ctx, id)
	if err != nil {
		log.Printf("failed to fetch active investment by user id: %v", err)
		return nil, err
	}

	return inv, nil
}

func (i *InvService) FetchAllInvByUserId(ctx context.Context, id uuid.UUID) ([]dto.InvestmentResponseDTO, error) {
	inv, err := i.invRepo.GetAllInvByUserId(ctx, id)
	if err != nil {
		log.Printf("failed to fetch all investment by user id: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch all investment by user id"}
	}

	return inv, nil
}
