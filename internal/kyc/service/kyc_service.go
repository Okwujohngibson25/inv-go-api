package service

import (
	"context"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/kyc/dto"
	"github.com/coinserveringo/internal/kyc/repository"
	userdatadto "github.com/coinserveringo/internal/user/dto"
	"github.com/coinserveringo/mail"
	"github.com/google/uuid"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*userdatadto.UserDataDTO, error)
	UpdateKycStatus(ctx context.Context, id uuid.UUID) error
}

type KycService struct {
	kycRepo  repository.KycRepository
	userRepo UserService
	mailer   *mail.MailService
}

func NewKycService(kycRepo repository.KycRepository, userRepo UserService, mailer *mail.MailService) *KycService {
	return &KycService{kycRepo: kycRepo, userRepo: userRepo, mailer: mailer}
}

func (k *KycService) SaveKycDocument(ctx context.Context, req *dto.KYCDTO) error {
	newKyc := dto.KYCDTO{
		UserID:       req.UserID,
		IDImageFront: req.IDImageFront,
		IDImageBack:  req.IDImageBack,
	}

	if err := k.kycRepo.Savekyc(ctx, &newKyc); err != nil {
		log.Printf("failed to save kyc data: %v", err)
		return apperrors.InternalError{Message: "failed to save kyc"}
	}
	return nil
}
