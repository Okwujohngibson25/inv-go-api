package service

import (
	"context"
	"errors"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/user/dto"
	"github.com/google/uuid"
)

func (u *UserService) FetchAllUsers(ctx context.Context) ([]dto.UserDataDTO, error) {
	users, err := u.userRepo.GetAllUsers(ctx)
	if err != nil {
		log.Printf("failed to fetch users: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch users"}
	}

	// usersData := mappers.MapUsersToDTOs(users)

	return users, nil
}

func (u *UserService) UpdateUserAcc(ctx context.Context, req *dto.UpdateUserBalanceDTO) error {
	// Check If User Exist
	_, err := u.userRepo.GetUserByID(ctx, req.Userid)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return apperrors.InternalError{Message: "failed to find user data"}
	}

	userUpdate := dto.UserDataDTO{
		WalletBalance:  req.WalletBalance,
		ProfitBalance:  req.ProfitBalance,
		InvestedAmount: req.InvestedAmount,
		ReferalBonus:   req.ReferalBonus,
	}

	if err := u.userRepo.UpdateUserAcc(ctx, req.Userid, &userUpdate); err != nil {
		log.Printf("failed to update user aaccount: %v", err)
		return apperrors.InternalError{Message: "failed to update user balance"}
	}

	return nil
}

func (u *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := u.userRepo.DeleteUserById(ctx, id); err != nil {
		log.Printf("failed to delete user by id: %v", err)
		return apperrors.InternalError{Message: "failed to delete user"}
	}
	return nil
}

const (
	StatusActive    = "active"
	StatusSuspended = "suspended"
)

func (u *UserService) UpdateAccountStatus(ctx context.Context, userID uuid.UUID) error {
	user, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.ValidationError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return apperrors.InternalError{Message: "could not fetch user data"}
	}

	var newStatus string
	switch user.Status {
	case StatusActive:
		newStatus = StatusSuspended
	case StatusSuspended:
		newStatus = StatusActive
	default:
		return apperrors.ValidationError{Message: "unknown status"}
	}

	if err := u.userRepo.UpdateAccountStatus(ctx, user.ID, newStatus); err != nil {
		log.Printf("failed to update user account status: %v", err)
		return apperrors.InternalError{Message: "failed to update user account status"}
	}

	return nil
}
