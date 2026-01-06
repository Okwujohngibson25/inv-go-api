package service

import (
	"context"
	"errors"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/deposit/dto"
	"github.com/coinserveringo/internal/deposit/repository"
	userdatadto "github.com/coinserveringo/internal/user/dto"
	"github.com/coinserveringo/mail"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*userdatadto.UserDataDTO, error)
	UpdateUserBalancebyId(ctx context.Context, id uuid.UUID, newBalance decimal.Decimal) error
}

type DepositService struct {
	depositRepo repository.DepositRepository
	userService UserService
	mailer      *mail.MailService
}

func NewDepositService(depositRepo repository.DepositRepository, userService UserService, mailer *mail.MailService) *DepositService {
	return &DepositService{depositRepo: depositRepo, userService: userService, mailer: mailer}
}

func (d *DepositService) CreateNewDeposit(ctx context.Context, req dto.DepositDTO) error {
	// check if user exist
	_, err := d.userService.GetUserByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return apperrors.InternalError{Message: "could not fetch user data"}
	}

	// Generate Transaction ID
	transactionID := uuid.New().String()

	newDeposit := dto.DepositDTO{
		UserID:        req.UserID,
		Amount:        req.Amount,
		TransactionID: transactionID,
		DepositProof:  req.DepositProof,
	}

	if err := d.depositRepo.CreateNewDeposit(ctx, &newDeposit); err != nil {
		log.Printf("failed to create new deposit: %v", err)
		return apperrors.InternalError{Message: "failed to make deposit"}
	}
	return nil
}
