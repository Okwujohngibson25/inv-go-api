package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PlanDTO struct {
	Name      string          `json:"name" binding:"required"`
	MinAmount decimal.Decimal `json:"min_amount" binding:"required"`
	MaxAmount decimal.Decimal `json:"max_amount" binding:"required"`
	Duration  int             `json:"duration" binding:"required"`
	Roi       decimal.Decimal `json:"roi" binding:"required"`
}

type PlanResponseDTO struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"plan_name"`
	MinAmount decimal.Decimal `json:"min_amount"`
	MaxAmount decimal.Decimal `json:"max_amount"`
	Duration  int             `json:"plan_duration"`
	Roi       decimal.Decimal `json:"roi"`
	CreatedAt time.Time       `json:"created_at"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
