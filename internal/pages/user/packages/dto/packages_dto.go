package dto

import (
	invresponsedto "github.com/coinserveringo/internal/investment/dto"
	"github.com/coinserveringo/internal/plans/dto"
	"github.com/shopspring/decimal"
)

type PackagesDTO struct {
	WalletBalance     decimal.Decimal                        `json:"wallet_balance"`
	InvestmentPlans   []dto.PlanResponseDTO                  `json:"investment_plans"`
	OngoingInvestment []invresponsedto.InvestmentResponseDTO `json:"ongoing_investment"`
	InvestmentHistory []invresponsedto.InvestmentResponseDTO `json:"investment_history"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
