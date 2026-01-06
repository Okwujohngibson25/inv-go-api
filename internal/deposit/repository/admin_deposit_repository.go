package repository

import (
	"context"

	"github.com/coinserveringo/internal/deposit/dto"
	userDataDTO "github.com/coinserveringo/internal/user/dto"
	"github.com/google/uuid"
)

func (r *SqlcRepository) GetAllDeposits(ctx context.Context) ([]dto.UserDepositDTO, error) {
	var depositData []dto.UserDepositDTO
	allDeposits, err := r.queries.GetAllDeposits(ctx)
	if err != nil {
		return nil, err
	}

	for _, value := range allDeposits {

		user := userDataDTO.UserDataDTO{
			Email:    value.Email,
			Fullname: value.Fullname,
			Username: value.Username,
		}

		depositData = append(depositData, dto.UserDepositDTO{
			ID:            value.ID,
			User:          user,
			Status:        value.Status,
			Amount:        value.Amount,
			TransactionID: value.TransactionID,
			DepositProof:  value.DepositProof,
			CreatedAt:     value.CreatedAt,
		})
	}

	return depositData, nil
}

func (r *SqlcRepository) GetAllPendingDeposit(ctx context.Context) ([]dto.UserDepositDTO, error) {
	var depositData []dto.UserDepositDTO
	allPendingDeposits, err := r.queries.GetAllPendingDeposit(ctx)
	if err != nil {
		return nil, err
	}

	for _, value := range allPendingDeposits {

		user := userDataDTO.UserDataDTO{
			Email:    value.Email,
			Fullname: value.Fullname,
			Username: value.Username,
		}

		depositData = append(depositData, dto.UserDepositDTO{
			ID:            value.ID,
			User:          user,
			Status:        value.Status,
			Amount:        value.Amount,
			TransactionID: value.TransactionID,
			DepositProof:  value.DepositProof,
			CreatedAt:     value.CreatedAt,
		})
	}

	return depositData, nil
}

func (r *SqlcRepository) GetDepositByID(ctx context.Context, id uuid.UUID) (*dto.UserDepositDTO, error) {
	deposit, err := r.queries.GetDepositByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.UserDepositDTO{
		ID:            deposit.ID,
		Status:        deposit.Status,
		Amount:        deposit.Amount,
		TransactionID: deposit.TransactionID,
		DepositProof:  deposit.DepositProof,
		CreatedAt:     deposit.CreatedAt,
	}, nil
}

func (r *SqlcRepository) ApproveDeposits(ctx context.Context, id uuid.UUID) error {
	if err := r.queries.ApproveDeposits(ctx, id); err != nil {
		return err
	}

	return nil
}

func (r *SqlcRepository) DeclineDeposit(ctx context.Context, id uuid.UUID) error {
	if err := r.queries.DeclineDeposit(ctx, id); err != nil {
		return err
	}

	return nil
}
