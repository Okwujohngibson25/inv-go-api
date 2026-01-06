package dto

import (
	"time"

	"github.com/coinserveringo/internal/user/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type InvestmentDTO struct {
	UserID         uuid.UUID       `json:"-"`
	PlanID         uuid.UUID       `json:"-"`
	Amount         decimal.Decimal `json:"amount" binding:"required"`
	StartDate      time.Time       `json:"-"`
	EndDate        time.Time       `json:"-"`
	ExpectedProfit decimal.Decimal `json:"-"`
}

type InvestmentResponseDTO struct {
	ID             uuid.UUID       `json:"id"`
	User           dto.UserDataDTO `json:"user"`
	PlanID         uuid.UUID       `json:"plan_id"`
	Amount         decimal.Decimal `json:"amount"`
	StartDate      time.Time       `json:"start_date"`
	EndDate        time.Time       `json:"end_date"`
	ExpectedProfit decimal.Decimal `json:"expected_profit"`
	Status         string          `json:"status"`
	CreatedAt      time.Time       `json:"created_at"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
