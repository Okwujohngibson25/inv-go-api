package repository

import (
	"context"

	"github.com/coinserveringo/internal/db"
	"github.com/coinserveringo/internal/user/dto"
	"github.com/google/uuid"
)

func (r *SqlcRepository) UpdateKycStatus(ctx context.Context, id uuid.UUID) error {
	if err := r.queries.UpdateKycStatus(ctx, id); err != nil {
		return err
	}
	return nil
}

func (r *SqlcRepository) AdminExists(ctx context.Context) (bool, error) {
	count, err := r.queries.AdminExists(ctx)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *SqlcRepository) GetAllUsers(ctx context.Context) ([]dto.UserDataDTO, error) {
	users, err := r.queries.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var result []dto.UserDataDTO

	for _, value := range users {
		result = append(result, dto.UserDataDTO{
			ID:             value.ID,
			Role:           value.Role,
			Email:          value.Email,
			Fullname:       value.Fullname,
			Username:       value.Username,
			Gender:         value.Gender,
			Country:        value.Country,
			PhoneNumber:    value.Country,
			WalletBalance:  value.WalletBalance,
			ProfitBalance:  value.ProfitBalance,
			InvestedAmount: value.InvestedAmount,
			ReferalBonus:   value.ReferalBonus,
			ReferalLink:    value.ReferalLink,
			Verified:       value.Verified,
			Status:         value.Status,
			CreatedAt:      value.CreatedAt,
		})
	}

	return result, nil
}

func (r *SqlcRepository) UpdateUserAcc(ctx context.Context, id uuid.UUID, data *dto.UserDataDTO) error {
	params := db.UpdateUserAccParams{
		WalletBalance:  data.WalletBalance,
		ProfitBalance:  data.ProfitBalance,
		InvestedAmount: data.InvestedAmount,
		ReferalBonus:   data.ReferalBonus,
		ID:             id,
	}

	if err := r.queries.UpdateUserAcc(ctx, params); err != nil {
		return err
	}

	return nil
}

func (r *SqlcRepository) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	if err := r.queries.DeleteUserById(ctx, id); err != nil {
		return err
	}
	return nil
}

func (r *SqlcRepository) UpdateAccountStatus(ctx context.Context, id uuid.UUID, newStatus string) error {
	params := db.UpdateAccountStatusParams{
		Status: newStatus,
		ID:     id,
	}
	if err := r.queries.UpdateAccountStatus(ctx, params); err != nil {
		return err
	}
	return nil
}

func (r *SqlcRepository) CountUsers(ctx context.Context) (int64, error) {
	total, err := r.queries.CountUsers(ctx)
	if err != nil {
		return 0, err
	}
	return total, nil
}
