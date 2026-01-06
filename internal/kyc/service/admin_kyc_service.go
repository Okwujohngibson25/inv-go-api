package service

import (
	"context"
	"errors"
	"log"

	"github.com/coinserveringo/config"
	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/kyc/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (k *KycService) GetAllKyc(ctx context.Context) ([]dto.UserKycDTO, error) {
	kyc, err := k.kycRepo.GetAllKyc(ctx)
	if err != nil {
		log.Printf("failed to fetch all kyc data: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch kyc data"}
	}

	return kyc, nil
}

func (k *KycService) GetAllPendingKyc(ctx context.Context) ([]dto.UserKycDTO, error) {
	pendingKyc, err := k.kycRepo.GetAllPendingKyc(ctx)
	if err != nil {
		log.Printf("failed to fetch all pending kyc data: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch all pending kyc data"}
	}

	return pendingKyc, nil
}

func (k *KycService) ApproveKyc(ctx context.Context, id uuid.UUID) error {
	// Get The Kyc
	kyc, err := k.kycRepo.GetKycById(ctx, id)
	if err != nil {
		log.Printf("failed to get kyc data by id: %v", err)
		return apperrors.InternalError{Message: "failed to fetch kyc data by id"}
	}

	// Get The User
	user, err := k.userRepo.GetUserByID(ctx, kyc.User.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return apperrors.InternalError{Message: "could not fetch user data"}
	}

	// Approve Kyc
	if err := k.kycRepo.Approvekyc(ctx, id); err != nil {
		log.Printf("failed to apporve kyc: %v", err)
		return apperrors.InternalError{Message: "failed to approve kyc"}
	}

	// Update User Kyc Status
	if err := k.userRepo.UpdateKycStatus(ctx, user.ID); err != nil {
		log.Printf("failed to update kyc status: %v", err)
		return apperrors.InternalError{Message: "failed to update kyc status"}
	}

	// Send Mail
	go func() {
		config, _ := config.LoadConfig()
		subject := "Your Kyc Was Approved! 🎉"
		data := map[string]interface{}{
			"Fullname":    user.Fullname,
			"CompanyName": config.CompanyName,
		}

		if err := k.mailer.SendMail(user.Email, subject, "kyc_verified.html", data); err != nil {
			log.Printf("failed to send email to client: %v", err)
		}
	}()

	return nil
}

func (k *KycService) Declinekyc(ctx context.Context, id uuid.UUID) error {
	if err := k.kycRepo.Declinekyc(ctx, id); err != nil {
		log.Printf("failed to decline kyc: %v", err)
		return apperrors.InternalError{Message: "failed to decline kyc"}
	}
	return nil
}
