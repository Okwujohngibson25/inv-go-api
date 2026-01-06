package service

import (
	"context"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/deposit/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (d *DepositService) GetAllPendingDepositByUserID(ctx context.Context, id uuid.UUID) ([]dto.UserDepositDTO, error) {
	pendingDeposits, err := d.depositRepo.GetAllPendingDepositByUserID(ctx, id)
	if err != nil {
		log.Printf("failed to fetch all pending deposit by user id: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch all pending deposit by user id"}
	}

	return pendingDeposits, nil
}

func (d *DepositService) SumAllDepositAmountByUserID(ctx context.Context, id uuid.UUID) (decimal.Decimal, error) {
	return d.depositRepo.SumAllDepositAmountByUserID(ctx, id)
}

func (d *DepositService) GetAllDepositsByUserId(ctx context.Context, id uuid.UUID) ([]dto.UserDepositDTO, error) {
	deposits, err := d.depositRepo.GetAllDepositsByUserId(ctx, id)
	if err != nil {
		log.Printf("failed to fetch all deposits by user id: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch all deposits by user id"}
	}

	return deposits, nil
}

func (d *DepositService) SumApprovedDeposits(ctx context.Context) (decimal.Decimal, error) {
	return d.depositRepo.SumApprovedDeposits(ctx)
}
