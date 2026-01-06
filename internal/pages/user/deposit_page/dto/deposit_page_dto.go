package dto

import (
	"github.com/coinserveringo/internal/walletaddress/dto"
	"github.com/shopspring/decimal"
)

type DepositPageDTO struct {
	TotalDeposit       decimal.Decimal                     `json:"total_deposit"`
	AdminWalletAddress []dto.AdminWalletAddressResponseDTO `json:"admin_wallet_address"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
