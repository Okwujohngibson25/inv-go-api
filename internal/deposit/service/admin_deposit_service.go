package service

import (
	"context"
	"errors"
	"log"

	"github.com/coinserveringo/config"
	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/deposit/dto"
	"github.com/google/uuid"
)

func (d *DepositService) GetAllDeposits(ctx context.Context) ([]dto.UserDepositDTO, error) {
	deposits, err := d.depositRepo.GetAllDeposits(ctx)
	if err != nil {
		log.Printf("failed to fetch all deposits: %v", err)
		return nil, apperrors.InternalError{Message: "unable to fetch deposit"}
	}

	return deposits, nil
}

func (d *DepositService) GetAllPendingDeposit(ctx context.Context) ([]dto.UserDepositDTO, error) {
	deposits, err := d.depositRepo.GetAllPendingDeposit(ctx)
	if err != nil {
		log.Printf("failed to fetch all pending deposits: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch all pending deposit"}
	}

	return deposits, nil
}

func (d *DepositService) ApproveDeposit(ctx context.Context, depositid uuid.UUID) error {
	// get the deposit
	deposit, err := d.depositRepo.GetDepositByID(ctx, depositid)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.NotFoundError{Message: "deposit not found"}
		}
		log.Printf("failed to find deposit by ID: %v", err)
		return apperrors.InternalError{Message: "unable to fetch deposit"}
	}

	// fetch user data
	user, err := d.userService.GetUserByID(ctx, deposit.User.ID)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return apperrors.InternalError{Message: "failed to find user"}
	}

	// Update User balance
	newUserBalance := user.WalletBalance.Add(deposit.Amount)

	if err := d.userService.UpdateUserBalancebyId(ctx, user.ID, newUserBalance); err != nil {
		log.Printf("failed to update user balance: %v", err)
		return apperrors.InternalError{Message: "failed to update user balance"}
	}

	// Approve the deposit
	if err := d.depositRepo.ApproveDeposits(ctx, depositid); err != nil {
		log.Printf("failed to approve deposit: %v", err)
		return apperrors.InternalError{Message: "failed to approve deposit"}
	}

	// Send welcome email
	go func() {
		config, _ := config.LoadConfig()
		subject := "Your Deposit Has Been Approved! 🎉"
		data := map[string]interface{}{
			"Fullname":      user.Fullname,
			"Amount":        deposit.Amount,
			"WalletBalance": user.WalletBalance,
			"CompanyName":   config.CompanyName,
		}

		if err := d.mailer.SendMail(user.Email, subject, "deposit_approved.html", data); err != nil {
			log.Printf("failed to send email to client: %v", err)
		}
	}()

	return nil
}

func (d *DepositService) DeclineDeposit(ctx context.Context, depositid uuid.UUID) error {
	// get the deposit
	deposit, err := d.depositRepo.GetDepositByID(ctx, depositid)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.NotFoundError{Message: "deposit not found"}
		}
		log.Printf("failed to find deposit by ID: %v", err)
		return apperrors.InternalError{Message: "unable to fetch deposit"}
	}

	// fetch user data
	user, err := d.userService.GetUserByID(ctx, deposit.User.ID)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return apperrors.InternalError{Message: "could not fetch user data"}
	}

	// Decline the deposit
	if err := d.depositRepo.DeclineDeposit(ctx, depositid); err != nil {
		log.Printf("failed to decline deposit: %v", err)
		return apperrors.InternalError{Message: "failed to decline deposit"}
	}

	// Send welcome email
	go func() {
		config, _ := config.LoadConfig()
		subject := "Your Deposit Was Declined!"
		data := map[string]interface{}{
			"Fullname":    user.Fullname,
			"Amount":      deposit.Amount,
			"CompanyName": config.CompanyName,
		}

		if err := d.mailer.SendMail(user.Email, subject, "deposit_declined.html", data); err != nil {
			log.Printf("failed to send email to client: %v", err)
		}
	}()

	return nil
}
