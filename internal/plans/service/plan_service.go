package service

import "github.com/coinserveringo/internal/plans/repository"

type PlanService struct {
	planRepo repository.PlanRepository
}

func NewPlanService(planRepo repository.PlanRepository) *PlanService {
	return &PlanService{planRepo: planRepo}
}
