package service

import (
	"context"
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/plans/dto"
	"github.com/google/uuid"
)

func (p *PlanService) FindPlanByID(ctx context.Context, id uuid.UUID) (*dto.PlanResponseDTO, error) {
	plan, err := p.planRepo.GetPlanByID(ctx, id)
	if err != nil {
		log.Printf("failed to fetch plan by id: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch plan by id"}
	}

	return plan, nil
}

func (p *PlanService) FetchAllPlans(ctx context.Context) ([]dto.PlanResponseDTO, error) {
	plans, err := p.planRepo.FetchAllPlans(ctx)
	if err != nil {
		log.Printf("failed to fetch all plans: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch all plans"}
	}

	return plans, nil
}
