package repository

import (
	"context"

	"github.com/coinserveringo/internal/kyc/dto"
	userDataDTO "github.com/coinserveringo/internal/user/dto"
	"github.com/google/uuid"
)

func (r *SqlcRepository) GetAllKyc(ctx context.Context) ([]dto.UserKycDTO, error) {
	var kycData []dto.UserKycDTO
	allKYC, err := r.queries.GetAllKyc(ctx)
	if err != nil {
		return nil, err
	}

	for _, value := range allKYC {
		user := userDataDTO.UserDataDTO{
			Username: value.Username,
			Email:    value.Email,
			Fullname: value.Fullname,
		}
		kycData = append(kycData, dto.UserKycDTO{
			ID:           value.ID,
			User:         user,
			Status:       value.Status,
			IDImageFront: value.IDImageFront,
			IDImageBack:  value.IDImageBack,
			CreatedAt:    value.CreatedAt,
		})

	}

	return kycData, nil
}

func (r *SqlcRepository) GetAllPendingKyc(ctx context.Context) ([]dto.UserKycDTO, error) {
	var kycData []dto.UserKycDTO
	pendingKyc, err := r.queries.GetAllPendingKyc(ctx)
	if err != nil {
		return nil, err
	}

	for _, value := range pendingKyc {
		user := userDataDTO.UserDataDTO{
			Username: value.Username,
			Email:    value.Email,
			Fullname: value.Fullname,
		}
		kycData = append(kycData, dto.UserKycDTO{
			ID:           value.ID,
			User:         user,
			Status:       value.Status,
			IDImageFront: value.IDImageFront,
			IDImageBack:  value.IDImageBack,
			CreatedAt:    value.CreatedAt,
		})

	}

	return kycData, nil
}

func (r *SqlcRepository) GetKycById(ctx context.Context, id uuid.UUID) (*dto.UserKycDTO, error) {
	kycData, err := r.queries.GetKycById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.UserKycDTO{
		ID:           kycData.ID,
		Status:       kycData.Status,
		IDImageFront: kycData.IDImageFront,
		IDImageBack:  kycData.IDImageBack,
		CreatedAt:    kycData.CreatedAt,
	}, nil
}

func (r *SqlcRepository) Approvekyc(ctx context.Context, id uuid.UUID) error {
	if err := r.queries.Approvekyc(context.Background(), id); err != nil {
		return err
	}

	return nil
}

func (r *SqlcRepository) Declinekyc(ctx context.Context, id uuid.UUID) error {
	if err := r.queries.Declinekyc(ctx, id); err != nil {
		return err
	}
	return nil
}
