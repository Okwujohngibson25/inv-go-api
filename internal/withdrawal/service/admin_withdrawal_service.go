package service

import (
	"context"
	"errors"
	"log"

	"github.com/coinserveringo/config"
	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/withdrawal/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (w *WithdrawalService) GetAllWithdrawal(ctx context.Context) ([]dto.UserWithdrawalDTO, error) {
	withdrawals, err := w.withdrawalRepo.GetAllWithdrawal(ctx)
	if err != nil {
		log.Printf("failed to get all withdrawal: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch all withdrawal"}
	}

	return withdrawals, nil
}

func (w *WithdrawalService) GetAllPendingWithdrawal(ctx context.Context) ([]dto.UserWithdrawalDTO, error) {
	pendingWithdrawals, err := w.withdrawalRepo.GetAllPendingWithdrawal(ctx)
	if err != nil {
		log.Printf("failed to fetch all pending withdrawl: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch all pending withdrawal"}
	}

	return pendingWithdrawals, nil
}

func (w *WithdrawalService) ApproveWithdrawal(ctx context.Context, id uuid.UUID) error {
	// Get The Withdrawal
	withdrawal, err := w.withdrawalRepo.GetWithdrawalById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NotFoundError{Message: "withdrawal not found"}
		}
		log.Printf("failed to fetch withdrawal by id: %v", err)
		return apperrors.InternalError{Message: "failed to fetch withdrawal"}
	}

	// Get The User
	user, err := w.userRepo.GetUserByID(ctx, withdrawal.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return apperrors.InternalError{Message: "could not fetch user data"}
	}

	// Approve The Withdrawal
	if err := w.withdrawalRepo.ApproveWithdrawal(ctx, id); err != nil {
		log.Printf("failed to approve withdrawal: %v", err)
		return apperrors.InternalError{Message: "failed to approve withdrawal"}
	}

	// Update User Wallet Balance
	NewWalletBalance := user.WalletBalance.Sub(withdrawal.Amount)

	if err := w.userRepo.UpdateUserBalancebyId(ctx, user.ID, NewWalletBalance); err != nil {
		log.Printf("failed to update user balance: %v", err)
		return apperrors.InternalError{Message: "failed to update user balance"}
	}

	// Send Mail
	go func() {
		config, _ := config.LoadConfig()
		subject := "Your Withdrawal Has Been Approved! 🎉"
		data := map[string]interface{}{
			"Fullname":      user.Fullname,
			"Amount":        withdrawal.Amount,
			"WalletBalance": NewWalletBalance,
			"CompanyName":   config.CompanyName,
		}

		if err := w.mailer.SendMail(user.Email, subject, "withdrawal_approve.html", data); err != nil {
			log.Printf("failed to send email to client: %v", err)
		}
	}()

	return nil
}

func (w *WithdrawalService) DeclineWithdrawal(ctx context.Context, id uuid.UUID) error {
	// Get The Withdrawal
	withdrawal, err := w.withdrawalRepo.GetWithdrawalById(ctx, id)
	if err != nil {
		log.Printf("failed to decline withdrawal: %v", err)
		return apperrors.InternalError{Message: "failed to decline deposit"}
	}

	// Get The User
	user, err := w.userRepo.GetUserByID(ctx, withdrawal.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return apperrors.InternalError{Message: "could not fetch user data"}
	}

	// Decline The Withdrawal
	if err := w.withdrawalRepo.DeclineWithdrawal(ctx, id); err != nil {
		log.Printf("failed to decline withdrawal: %v", err)
		return apperrors.InternalError{Message: "failed to decline withdrawal"}
	}

	// Send Mail
	go func() {
		config, _ := config.LoadConfig()
		subject := "Your Withdrawal Has Been Declined!"
		data := map[string]interface{}{
			"Fullname":    user.Fullname,
			"Amount":      withdrawal.Amount,
			"CompanyName": config.CompanyName,
		}

		if err := w.mailer.SendMail(user.Email, subject, "withdrawal_declined.html", data); err != nil {
			log.Printf("failed to send email to client: %v", err)
		}
	}()

	return nil
}
