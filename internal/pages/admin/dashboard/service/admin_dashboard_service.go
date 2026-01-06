package service

import (
	"context"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	depositdatadto "github.com/coinserveringo/internal/deposit/dto"
	"github.com/coinserveringo/internal/pages/admin/dashboard/dto"
	withdrawaldto "github.com/coinserveringo/internal/withdrawal/dto"
	"github.com/shopspring/decimal"
)

type (
	UserService interface {
		CountUsers(ctx context.Context) (int64, error)
	}
	DepositService interface {
		SumApprovedDeposits(ctx context.Context) (decimal.Decimal, error)
		GetAllPendingDeposit(ctx context.Context) ([]depositdatadto.UserDepositDTO, error)
	}
	WithdrawalService interface {
		SumApprovedWithdrawals(ctx context.Context) (decimal.Decimal, error)
		GetAllPendingWithdrawal(ctx context.Context) ([]withdrawaldto.UserWithdrawalDTO, error)
	}
)

type AdminDashboardService struct {
	userService       UserService
	depositService    DepositService
	withdrawalService WithdrawalService
}

func NewAdminDashboardService(userService UserService, depositService DepositService, withdrawalService WithdrawalService) *AdminDashboardService {
	return &AdminDashboardService{userService: userService, depositService: depositService, withdrawalService: withdrawalService}
}

func (d *AdminDashboardService) GetAdminDashboardData(ctx context.Context) (*dto.AdminDashboardDTO, error) {
	// get total number of users
	userCount, err := d.userService.CountUsers(ctx)
	if err != nil {
		log.Printf("failed to count user: %v", err)
		return nil, apperrors.InternalError{Message: "user count operation failed"}
	}

	// get sum of approved deposits
	sumApprovedDepost, err := d.depositService.SumApprovedDeposits(ctx)
	if err != nil {
		log.Printf("failed to sum approved deposit: %v", err)
		return nil, apperrors.InternalError{Message: "failed to sum approved deposit"}
	}

	// get sum of approved withdrawals
	sumApprovedWithdrawal, err := d.withdrawalService.SumApprovedWithdrawals(ctx)
	if err != nil {
		log.Printf("failed to sum approved withdrawal: %v", err)
		return nil, apperrors.InternalError{Message: "failed to sum approved withdrawal"}
	}

	// all pending depoists
	pendingDeposits, err := d.depositService.GetAllPendingDeposit(ctx)
	if err != nil {
		log.Printf("failed to get all pending deposit: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch all pending deposit"}
	}

	// all pending withdrawals
	pendingWithdrawals, err := d.withdrawalService.GetAllPendingWithdrawal(ctx)
	if err != nil {
		log.Printf("failed to get all pending withdrawal: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch all pending withdrawal"}
	}

	// send all in a dto back as json
	data := dto.AdminDashboardDTO{
		MembersCount:                   userCount,
		SumAmountOfApprovedWithdrawals: sumApprovedWithdrawal,
		SumAmountOfApprovedDeposits:    sumApprovedDepost,
		AllPendingDepoists:             pendingDeposits,
		AllPendingWithdrawals:          pendingWithdrawals,
	}

	return &data, nil
}
