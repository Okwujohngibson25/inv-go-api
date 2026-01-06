package service

import (
	"context"
	"fmt"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/withdrawal/dto"
	"github.com/coinserveringo/internal/withdrawal/repository"
	"github.com/coinserveringo/mail"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	userdatadto "github.com/coinserveringo/internal/user/dto"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*userdatadto.UserDataDTO, error)
	UpdateUserBalancebyId(ctx context.Context, id uuid.UUID, newBalance decimal.Decimal) error
}

type WithdrawalService struct {
	withdrawalRepo repository.WithdrawalRepository
	userRepo       UserService
	mailer         *mail.MailService
}

func NewWithdrawalService(withdrawalRepo repository.WithdrawalRepository, userRepo UserService, mailer *mail.MailService) *WithdrawalService {
	return &WithdrawalService{withdrawalRepo: withdrawalRepo, userRepo: userRepo, mailer: mailer}
}

func (w *WithdrawalService) CreateNewWithdrawal(ctx context.Context, req *dto.WithdrawalDTO) error {
	// check if the user exist
	user, err := w.userRepo.GetUserByID(ctx, req.UserID)
	if err != nil {
		fmt.Printf("failed to find user by id: %v", err)
		return apperrors.InternalError{Message: "could not fetch user data"}
	}

	// Generate transaction ID
	transactionID := uuid.New().String()

	if user.WalletBalance.LessThanOrEqual(req.Amount) {
		return apperrors.ValidationError{
			Message: "cannot process transaction insufficient wallet balance",
		}
	}

	if !user.Verified {
		return apperrors.ValidationError{Message: "account not verified, please complete KYC"}
	}

	// create the new withdrawal
	newWithdrawal := dto.UserWithdrawalDTO{
		UserID:        req.UserID,
		Amount:        req.Amount,
		TransactionID: transactionID,
		CoinType:      req.CoinType,
		WalletAddress: req.WalletAddress,
	}

	if err := w.withdrawalRepo.CreateNewWithdrawal(ctx, &newWithdrawal); err != nil {
		log.Printf("failed to create new withdrawal: %v", err)
		return apperrors.InternalError{Message: "failed to request for withdrawal"}
	}
	return nil
}
