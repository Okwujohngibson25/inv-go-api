package service

import (
	"context"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/withdrawal/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (w *WithdrawalService) GetAllPendingWithdrawalByID(ctx context.Context, id uuid.UUID) ([]dto.UserWithdrawalDTO, error) {
	pendingWithdrawals, err := w.withdrawalRepo.GetAllPendingWithdrawalByID(ctx, id)
	if err != nil {
		log.Printf("failed to fetch pending withdrawal by id: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch pending withdrawal by id"}
	}

	return pendingWithdrawals, nil
}

func (w *WithdrawalService) GetAllWithdrawalByUserId(ctx context.Context, id uuid.UUID) ([]dto.UserWithdrawalDTO, error) {
	withdrawal, err := w.withdrawalRepo.GetAllWithdrawalByUserId(ctx, id)
	if err != nil {
		log.Printf("failed to fetch all withdrawal by user id: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch all withdrawal by user id"}
	}

	return withdrawal, nil
}

func (w *WithdrawalService) SumAllWithdrawalAmountByUserID(ctx context.Context, id uuid.UUID) (decimal.Decimal, error) {
	return w.withdrawalRepo.SumAllWithdrawalAmountByUserID(ctx, id)
}

func (w *WithdrawalService) SumApprovedWithdrawals(ctx context.Context) (decimal.Decimal, error) {
	return w.withdrawalRepo.SumApprovedWithdrawals(ctx)
}
