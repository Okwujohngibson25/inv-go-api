package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/investment/dto"
	"github.com/coinserveringo/internal/investment/repository"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	plandatadto "github.com/coinserveringo/internal/plans/dto"
	userdatadto "github.com/coinserveringo/internal/user/dto"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*userdatadto.UserDataDTO, error)
	UpdateUserBalancebyId(ctx context.Context, id uuid.UUID, newBalance decimal.Decimal) error
}

type PlanService interface {
	FindPlanByID(ctx context.Context, id uuid.UUID) (*plandatadto.PlanResponseDTO, error)
}

type InvService struct {
	invRepo  repository.InvRepository
	userRepo UserService
	planRepo PlanService
}

func NewInvService(invRepo repository.InvRepository, userRepo UserService, planRepo PlanService) *InvService {
	return &InvService{invRepo: invRepo, userRepo: userRepo, planRepo: planRepo}
}

func (i *InvService) Joininv(ctx context.Context, req *dto.InvestmentDTO) error {
	// Get User Wallet Balance
	user, err := i.userRepo.GetUserByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return apperrors.InternalError{Message: "unable to fetch user"}
	}

	// get inv plan by id
	plan, err := i.planRepo.FindPlanByID(ctx, req.PlanID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NotFoundError{Message: "plan not found"}
		}
		log.Printf("failed to find plan by id: %v", err)
		return apperrors.InternalError{Message: "failed to find plan"}
	}

	// check if they are eligable to join this plan
	if user.WalletBalance.LessThan(plan.MinAmount) {
		return apperrors.ValidationError{
			Message: fmt.Sprintf("insufficient wallet balance — the minimum required amount for this plan is $%.2f", plan.MinAmount),
		}
	}

	if req.Amount.GreaterThan(plan.MaxAmount) {
		return apperrors.ValidationError{
			Message: fmt.Sprintf("amount exceeds the maximum limit for this plan (maximum allowed: $%.2f)", plan.MaxAmount),
		}
	}

	if req.Amount.LessThan(plan.MinAmount) {
		return apperrors.ValidationError{
			Message: fmt.Sprintf("amount $%.2f is below the minimum required amount of $%.2f for this plan", req.Amount, plan.MinAmount),
		}
	}

	// create the join reqs
	newInv := dto.InvestmentDTO{
		UserID:         req.UserID,
		PlanID:         req.PlanID,
		Amount:         req.Amount,
		EndDate:        time.Now().Add(time.Hour * 24 * time.Duration(plan.Duration)),
		ExpectedProfit: req.Amount.Mul(plan.Roi),
	}

	// pass to repo to subscribe to plan
	if err := i.invRepo.CreateInv(ctx, &newInv); err != nil {
		log.Printf("failed to create investment: %v", err)
		return apperrors.InternalError{Message: "failed to create new investment"}
	}

	// Update The User WalletBalance
	newUserBalance := user.WalletBalance.Sub(req.Amount)
	if err := i.userRepo.UpdateUserBalancebyId(ctx, user.ID, newUserBalance); err != nil {
		log.Printf("failed to update user balance")
		return apperrors.InternalError{Message: "failed to update user balance"}
	}

	return nil
}
