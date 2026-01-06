package service

import (
	"context"
	"errors"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/pages/user/dashboard/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	depositdatadto "github.com/coinserveringo/internal/deposit/dto"
	userdatadto "github.com/coinserveringo/internal/user/dto"
	withdrawaldatadto "github.com/coinserveringo/internal/withdrawal/dto"
)

type (
	UserService interface {
		GetUserByID(ctx context.Context, id uuid.UUID) (*userdatadto.UserDataDTO, error)
	}

	DepositService interface {
		GetAllPendingDepositByUserID(ctx context.Context, id uuid.UUID) ([]depositdatadto.UserDepositDTO, error)
		SumAllDepositAmountByUserID(ctx context.Context, id uuid.UUID) (decimal.Decimal, error)
	}

	WithdrawalService interface {
		GetAllPendingWithdrawalByID(ctx context.Context, id uuid.UUID) ([]withdrawaldatadto.UserWithdrawalDTO, error)
		SumAllWithdrawalAmountByUserID(ctx context.Context, id uuid.UUID) (decimal.Decimal, error)
	}
)

type DashboardService struct {
	userService       UserService
	depositService    DepositService
	withdrawalService WithdrawalService
}

func NewDashboardService(userService UserService, depositService DepositService, withdrawalService WithdrawalService) *DashboardService {
	return &DashboardService{userService: userService, depositService: depositService, withdrawalService: withdrawalService}
}

func (d *DashboardService) GetDashboardData(ctx context.Context, id uuid.UUID) (*dto.DashboardDTO, error) {
	// Get User By ID
	user, err := d.userService.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return nil, apperrors.InternalError{Message: "user not found"}
	}

	// Get All Pending Deposits
	pendingDeposits, err := d.depositService.GetAllPendingDepositByUserID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NotFoundError{Message: "no pending deposit found"}
		}
		log.Printf("failed to fetch pending deposit: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch pending deposits"}
	}

	// Get All pending Withdrawals
	pendingWithdrawls, err := d.withdrawalService.GetAllPendingWithdrawalByID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NotFoundError{Message: "no pending withdrawals found"}
		}
		log.Printf("failed to fetch pending withdrawal: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch pending withdrawals"}
	}

	// Get Total Amount Of User Deposits
	totalAmountOfDeposit, err := d.depositService.SumAllDepositAmountByUserID(ctx, user.ID)
	if err != nil {
		log.Printf("failed to sum deposit amount: %v", err)
		return nil, apperrors.InternalError{Message: "failed to sum deposit amount"}
	}

	// Get Total Amount of User Withdrawal
	totalAmountOfWithdrawal, err := d.withdrawalService.SumAllWithdrawalAmountByUserID(ctx, user.ID)
	if err != nil {
		log.Printf("failed to sum withdrawal amount: %v", err)
		return nil, apperrors.InternalError{Message: "failed to sum withdrawal amount"}
	}

	return &dto.DashboardDTO{
		Username:           user.Username,
		WalletBalance:      user.WalletBalance,
		Profit:             user.ProfitBalance,
		Bonus:              user.ReferalBonus,
		TotalDeposit:       totalAmountOfDeposit,
		TotalWithdrawal:    totalAmountOfWithdrawal,
		PendingDeposits:    pendingDeposits,
		PendingWithdrawals: pendingWithdrawls,
		ReferalLink:        user.ReferalLink,
	}, nil
}
