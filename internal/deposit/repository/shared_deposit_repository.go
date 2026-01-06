package repository

import (
	"context"

	"github.com/coinserveringo/internal/deposit/dto"
	userDataDTO "github.com/coinserveringo/internal/user/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (r *SqlcRepository) GetAllDepositsByUserId(ctx context.Context, id uuid.UUID) ([]dto.UserDepositDTO, error) {
	var deposits []dto.UserDepositDTO
	allDeposits, err := r.queries.GetAllDepositsByUserId(context.Background(), id)
	if err != nil {
		return nil, err
	}
	for _, value := range allDeposits {

		userDTO := userDataDTO.UserDataDTO{
			Email:     value.Email,
			Fullname:  value.Fullname,
			Username:  value.Username,
			CreatedAt: value.CreatedAt,
		}

		deposits = append(deposits, dto.UserDepositDTO{
			ID:            value.ID,
			User:          userDTO,
			Status:        value.Status,
			Amount:        value.Amount,
			TransactionID: value.TransactionID,
			DepositProof:  value.DepositProof,
			CreatedAt:     value.CreatedAt,
		})
	}

	return deposits, nil
}

func (r *SqlcRepository) GetAllPendingDepositByUserID(ctx context.Context, id uuid.UUID) ([]dto.UserDepositDTO, error) {
	var deposits []dto.UserDepositDTO
	pendingDeposits, err := r.queries.GetAllDepositsByUserId(context.Background(), id)
	if err != nil {
		return nil, err
	}

	for _, value := range pendingDeposits {

		userDTO := userDataDTO.UserDataDTO{
			Email:     value.Email,
			Fullname:  value.Fullname,
			Username:  value.Username,
			CreatedAt: value.CreatedAt,
		}

		deposits = append(deposits, dto.UserDepositDTO{
			ID:            value.ID,
			User:          userDTO,
			Status:        value.Status,
			Amount:        value.Amount,
			TransactionID: value.TransactionID,
			DepositProof:  value.DepositProof,
			CreatedAt:     value.CreatedAt,
		})
	}

	return deposits, nil
}

func (r *SqlcRepository) SumAllDepositAmountByUserID(ctx context.Context, id uuid.UUID) (decimal.Decimal, error) {
	sumDeposits, err := r.queries.SumAllDepositAmountByUserID(context.Background(), id)
	if err != nil {
		return decimal.Zero, err
	}

	return sumDeposits, nil
}

func (r *SqlcRepository) SumApprovedDeposits(ctx context.Context) (decimal.Decimal, error) {
	sumDeposits, err := r.queries.SumApprovedDeposits(context.Background())
	if err != nil {
		return decimal.Zero, err
	}

	return sumDeposits, nil
}
