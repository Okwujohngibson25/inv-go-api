package dto

import (
	"time"

	"github.com/coinserveringo/internal/user/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type WithdrawalDTO struct {
	Amount        decimal.Decimal `json:"amount" binding:"required"`
	CoinType      string          `json:"coin_type" binding:"required"`
	WalletAddress string          `json:"wallet_address" binding:"required"`
	UserID        uuid.UUID       `json:"-"`
}

type UserWithdrawalDTO struct {
	ID            uuid.UUID       `json:"id"`
	UserID        uuid.UUID       `json:"-"`
	User          dto.UserDataDTO `json:"user"`
	Status        string          `json:"status"`
	Amount        decimal.Decimal `json:"amount"`
	TransactionID string          `json:"transaction_id"`
	CoinType      string          `json:"coin_type"`
	WalletAddress string          `json:"wallet_address"`
	CreatedAt     time.Time       `json:"created_at"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
