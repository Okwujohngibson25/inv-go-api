package service

import (
	"context"
	"errors"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/user/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (u *UserService) UpdateUserBalancebyId(ctx context.Context, id uuid.UUID, newBalance decimal.Decimal) error {
	return u.userRepo.UpdateUserBalancebyId(ctx, id, newBalance)
}

func (u *UserService) UpdateKycStatus(ctx context.Context, id uuid.UUID) error {
	return u.userRepo.UpdateKycStatus(ctx, id)
}

func (u *UserService) CountUsers(ctx context.Context) (int64, error) {
	return u.userRepo.CountUsers(ctx)
}

func (u *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*dto.UserDataDTO, error) {
	user, err := u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return &dto.UserDataDTO{}, apperrors.NotFoundError{Message: "error finding user"}
		}
		log.Printf("error finding user: %v", err)
		return &dto.UserDataDTO{}, apperrors.InternalError{Message: "error encountered when finding user"}
	}

	return user, nil
}
