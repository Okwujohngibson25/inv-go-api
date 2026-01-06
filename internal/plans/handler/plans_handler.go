package handler

import (
	"github.com/coinserveringo/internal/plans/service"
)

type PlanHandler struct {
	PlanService *service.PlanService
}

func NewPlanHandler(PlanService *service.PlanService) *PlanHandler {
	return &PlanHandler{PlanService: PlanService}
}
