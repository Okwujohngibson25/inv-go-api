package service

import (
	"context"
	"errors"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/pages/user/deposit_page/dto"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/google/uuid"

	userdatadto "github.com/coinserveringo/internal/user/dto"
	adminwalletresponsedto "github.com/coinserveringo/internal/walletaddress/dto"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*userdatadto.UserDataDTO, error)
}

type DepositService interface {
	SumAllDepositAmountByUserID(ctx context.Context, id uuid.UUID) (decimal.Decimal, error)
}

type WalletAddressService interface {
	FetchAdminWalletAddress() ([]adminwalletresponsedto.AdminWalletAddressResponseDTO, error)
}

type DepositPageService struct {
	userService          UserService
	depositService       DepositService
	walletAddressService WalletAddressService
}

func NewDepositPageService(userService UserService, depositService DepositService, walletAddressService WalletAddressService) *DepositPageService {
	return &DepositPageService{userService: userService, depositService: depositService, walletAddressService: walletAddressService}
}

func (d *DepositPageService) GetDepositPackageData(ctx context.Context, id uuid.UUID) (*dto.DepositPageDTO, error) {
	// Get User By ID
	user, err := d.userService.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return nil, apperrors.InternalError{Message: "user not found"}
	}

	// Get sum of all user deposits
	sumOfDepositAmount, err := d.depositService.SumAllDepositAmountByUserID(ctx, user.ID)
	if err != nil {
		log.Printf("failed to get sum of deposit amount: %v", err)
		return nil, apperrors.InternalError{Message: "failed to get sum of deposit amount"}
	}

	// Fetch admin wallet
	adminWallets, err := d.walletAddressService.FetchAdminWalletAddress()
	if err != nil {
		log.Printf("failed to fetch admin wallet: %v", err)
		return nil, apperrors.InternalError{Message: "admin wallet not found"}
	}

	depositPackageData := dto.DepositPageDTO{
		TotalDeposit:       sumOfDepositAmount,
		AdminWalletAddress: adminWallets,
	}

	return &depositPackageData, nil
}
