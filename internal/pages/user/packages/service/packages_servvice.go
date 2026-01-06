package service

import (
	"context"
	"errors"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	invresponsedto "github.com/coinserveringo/internal/investment/dto"
	"github.com/coinserveringo/internal/pages/user/packages/dto"
	planresponsedto "github.com/coinserveringo/internal/plans/dto"
	userdatadto "github.com/coinserveringo/internal/user/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*userdatadto.UserDataDTO, error)
}

type PlanService interface {
	FetchAllPlans(ctx context.Context) ([]planresponsedto.PlanResponseDTO, error)
}

type InvService interface {
	FetchAllInvByUserId(ctx context.Context, id uuid.UUID) ([]invresponsedto.InvestmentResponseDTO, error)
	FetchActiveInvByUserId(ctx context.Context, id uuid.UUID) ([]invresponsedto.InvestmentResponseDTO, error)
}

type PackagesService struct {
	userService UserService
	planService PlanService
	invService  InvService
}

func NewPackagesService(userService UserService, planService PlanService, invService InvService) *PackagesService {
	return &PackagesService{userService: userService, planService: planService, invService: invService}
}

func (p *PackagesService) GetPackagesDate(ctx context.Context, id uuid.UUID) (*dto.PackagesDTO, error) {
	// fetch User Using Id
	user, err := p.userService.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to fetch user by id: %v", err)
		return nil, apperrors.InternalError{Message: "user not found"}
	}

	// all Inv Made By User
	allInv, err := p.invService.FetchAllInvByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("failed to fetch investment: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch investment"}
	}

	// fetch All Active Inv. By User
	activeInv, err := p.invService.FetchActiveInvByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("failed to fetch active inv by user id: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch active inv."}

	}

	// fetch All Investment Packages
	plans, err := p.planService.FetchAllPlans(ctx)
	if err != nil {
		log.Printf("failed to fetch plan: %v", err)
		return nil, apperrors.InternalError{Message: "plan not found"}
	}

	packagesData := dto.PackagesDTO{
		WalletBalance:     user.WalletBalance,
		InvestmentPlans:   plans,
		OngoingInvestment: activeInv,
		InvestmentHistory: allInv,
	}

	return &packagesData, nil
}
