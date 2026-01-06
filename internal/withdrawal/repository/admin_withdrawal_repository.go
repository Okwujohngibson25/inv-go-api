package repository

import (
	"context"

	userDataDTO "github.com/coinserveringo/internal/user/dto"
	"github.com/coinserveringo/internal/withdrawal/dto"
	"github.com/google/uuid"
)

func (r *SqlcRepository) ApproveWithdrawal(ctx context.Context, id uuid.UUID) error {
	if err := r.Queries.ApproveWithdrawal(ctx, id); err != nil {
		return err
	}

	return nil
}

func (r *SqlcRepository) DeclineWithdrawal(ctx context.Context, id uuid.UUID) error {
	if err := r.Queries.DeclineWithdrawal(ctx, id); err != nil {
		return err
	}

	return nil
}

func (r *SqlcRepository) GetAllWithdrawal(ctx context.Context) ([]dto.UserWithdrawalDTO, error) {
	var withdrawalData []dto.UserWithdrawalDTO
	allWithdrawal, err := r.Queries.GetAllWithdrawal(ctx)
	if err != nil {
		return nil, err
	}

	for _, value := range allWithdrawal {

		users := userDataDTO.UserDataDTO{
			Username: value.Username,
			Email:    value.Email,
			Fullname: value.Fullname,
		}

		withdrawalData = append(withdrawalData, dto.UserWithdrawalDTO{
			ID:            value.ID,
			UserID:        value.UserID,
			User:          users,
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

func (r *SqlcRepository) GetAllPendingWithdrawal(ctx context.Context) ([]dto.UserWithdrawalDTO, error) {
	var withdrawalData []dto.UserWithdrawalDTO
	allPendingWithdrawal, err := r.Queries.GetAllPendingWithdrawal(ctx)
	if err != nil {
		return nil, err
	}

	for _, value := range allPendingWithdrawal {

		users := userDataDTO.UserDataDTO{
			Username: value.Username,
			Email:    value.Email,
			Fullname: value.Fullname,
		}

		withdrawalData = append(withdrawalData, dto.UserWithdrawalDTO{
			ID:            value.ID,
			UserID:        value.UserID,
			User:          users,
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

func (r *SqlcRepository) GetWithdrawalById(ctx context.Context, id uuid.UUID) (*dto.UserWithdrawalDTO, error) {
	withdrawal, err := r.Queries.GetWithdrawalById(ctx, id)
	if err != nil {
		return nil, err
	}

	users := userDataDTO.UserDataDTO{
		Username: withdrawal.Username,
		Email:    withdrawal.Email,
		Fullname: withdrawal.Fullname,
	}

	return &dto.UserWithdrawalDTO{
		ID:            withdrawal.ID,
		UserID:        withdrawal.UserID,
		User:          users,
		Status:        withdrawal.Status,
		Amount:        withdrawal.Amount,
		TransactionID: withdrawal.TransactionID,
		CoinType:      withdrawal.CoinType,
		WalletAddress: withdrawal.WalletAddress,
		CreatedAt:     withdrawal.CreatedAt,
	}, nil
}
