package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type UpdateUserBalanceDTO struct {
	Userid         uuid.UUID       `json:"-"`
	WalletBalance  decimal.Decimal `json:"wallet_balance" binding:"required"`
	ProfitBalance  decimal.Decimal `json:"profit_balance" binding:"required"`
	InvestedAmount decimal.Decimal `json:"invested_amount" binding:"required"`
	ReferalBonus   decimal.Decimal `json:"referal_bonus" binding:"required"`
}
