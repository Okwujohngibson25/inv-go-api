package repository

import (
	"context"

	"github.com/coinserveringo/internal/withdrawal/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (r *SqlcRepository) GetAllPendingWithdrawalByID(ctx context.Context, id uuid.UUID) ([]dto.UserWithdrawalDTO, error) {
	var withdrawalData []dto.UserWithdrawalDTO
	pendingWithdrawal, err := r.Queries.GetAllPendingWithdrawalByID(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, value := range pendingWithdrawal {
		withdrawalData = append(withdrawalData, dto.UserWithdrawalDTO{
			ID:            value.ID,
			UserID:        value.UserID,
			Status:        value.Status,
			Amount:        value.Amount,
			TransactionID: value.TransactionID,
			CoinType:      value.CoinType,
			WalletAddress: value.WalletAddress,
			CreatedAt:     value.CreatedAt,
		})
	}

	return withdrawalData, nil
}

func (r *SqlcRepository) GetAllWithdrawalByUserId(ctx context.Context, id uuid.UUID) ([]dto.UserWithdrawalDTO, error) {
	var withdrawalData []dto.UserWithdrawalDTO
	allWithdrawals, err := r.Queries.GetAllWithdrawalByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, value := range allWithdrawals {
		withdrawalData = append(withdrawalData, dto.UserWithdrawalDTO{
			ID:            value.ID,
			UserID:        value.UserID,
			Status:        value.Status,
			Amount:        value.Amount,
			TransactionID: value.TransactionID,
			CoinType:      value.CoinType,
			WalletAddress: value.WalletAddress,
			CreatedAt:     value.CreatedAt,
		})
	}

	return withdrawalData, nil
}

func (r *SqlcRepository) SumAllWithdrawalAmountByUserID(ctx context.Context, id uuid.UUID) (decimal.Decimal, error) {
	sum, err := r.Queries.SumAllWithdrawalAmountByUserID(ctx, id)
	if err != nil {
		return decimal.Zero, err
	}

	return sum, nil
}

func (r *SqlcRepository) SumApprovedWithdrawals(ctx context.Context) (decimal.Decimal, error) {
	sum, err := r.Queries.SumApprovedWithdrawals(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	return sum, nil
}
