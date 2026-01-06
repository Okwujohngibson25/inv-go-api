package service

import (
	"context"
	"log"
	"strings"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/plans/dto"
	"github.com/google/uuid"
	// <-- needed for pq.Error
)

func (p *PlanService) CreatePlan(ctx context.Context, req *dto.PlanDTO) error {
	if err := p.planRepo.CreatePlan(ctx, req); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return apperrors.ConflictError{Message: "plan name already exists"}
		}
		log.Printf("failed to create plan: %v", err)
		return apperrors.InternalError{Message: "failed to create plan"}
	}

	return nil
}

func (p *PlanService) DeletePlan(ctx context.Context, id uuid.UUID) error {
	if err := p.planRepo.DeletePlanById(ctx, id); err != nil {
		log.Printf("failed to delete plan by id: %v", err)
		return apperrors.InternalError{Message: "failed to delete plan"}
	}
	return nil
}
