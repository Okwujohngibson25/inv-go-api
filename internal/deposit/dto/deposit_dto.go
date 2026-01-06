package dto

import (
	"time"

	"github.com/coinserveringo/internal/user/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type DepositDTO struct {
	Amount        decimal.Decimal `form:"amount" binding:"required"`
	DepositProof  string          `json:"-"`
	UserID        uuid.UUID       `json:"-"`
	TransactionID string          `json:"-"`
}

type UserDepositDTO struct {
	ID            uuid.UUID       `json:"id"`
	User          dto.UserDataDTO `json:"user"`
	Status        string          `json:"status"`
	Amount        decimal.Decimal `json:"amount"`
	TransactionID string          `json:"transaction_id"`
	DepositProof  string          `json:"deposit_proof"`
	CreatedAt     time.Time       `json:"created_at"`
}
