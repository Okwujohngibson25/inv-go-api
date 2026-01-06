package service

import (
	"context"
	"errors"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	depositdatadto "github.com/coinserveringo/internal/deposit/dto"
	"github.com/coinserveringo/internal/pages/user/transactions/dto"
	userdatadto "github.com/coinserveringo/internal/user/dto"
	withdrawaldatadto "github.com/coinserveringo/internal/withdrawal/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*userdatadto.UserDataDTO, error)
}

type DepositService interface {
	GetAllDepositsByUserId(ctx context.Context, id uuid.UUID) ([]depositdatadto.UserDepositDTO, error)
}

type WithdrawalService interface {
	GetAllWithdrawalByUserId(ctx context.Context, id uuid.UUID) ([]withdrawaldatadto.UserWithdrawalDTO, error)
}

type TransactionsService struct {
	userService       UserService
	depositService    DepositService
	withdrawalService WithdrawalService
}

func NewTransactionsService(userService UserService, depositService DepositService, withdrawalService WithdrawalService) *TransactionsService {
	return &TransactionsService{userService: userService, depositService: depositService, withdrawalService: withdrawalService}
}

func (t *TransactionsService) GetTransactionData(ctx context.Context, id uuid.UUID) (*dto.TransactionsDTO, error) {
	// Get User By ID
	user, err := t.userService.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to fetch user by id: %v", err)
		return nil, apperrors.InternalError{Message: "user not found"}
	}

	// Get all deposit transactions
	allDeposits, err := t.depositService.GetAllDepositsByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("failed to fetch deposit using user id: %v", err)
		return nil, apperrors.InternalError{Message: "deposit not found"}
	}

	// Get all withdrawal transactions
	allWithdrawals, err := t.withdrawalService.GetAllWithdrawalByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("failed to fetch withdrawal using user id: %v", err)
		return nil, apperrors.InternalError{Message: "withdrawal not found"}
	}

	// Return DTO
	transactions := dto.TransactionsDTO{
		AllDeposits:    allDeposits,
		AllWithdrawals: allWithdrawals,
	}

	return &transactions, nil
}
