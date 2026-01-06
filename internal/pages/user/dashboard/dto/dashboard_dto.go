package dto

import (
	depositdatadto "github.com/coinserveringo/internal/deposit/dto"
	withdrawalresponsedto "github.com/coinserveringo/internal/withdrawal/dto"
	"github.com/shopspring/decimal"
)

type DashboardDTO struct {
	Username           string                                    `json:"username"`
	WalletBalance      decimal.Decimal                           `json:"wallet_balance"`
	Profit             decimal.Decimal                           `json:"profit"`
	Bonus              decimal.Decimal                           `json:"bonus"`
	TotalDeposit       decimal.Decimal                           `json:"total_deposit"`
	TotalWithdrawal    decimal.Decimal                           `json:"total_withdrawal"`
	PendingDeposits    []depositdatadto.UserDepositDTO           `json:"pending_deposits"`
	PendingWithdrawals []withdrawalresponsedto.UserWithdrawalDTO `json:"pending_withdrawals"`
	ReferalLink        string                                    `json:"referal_link"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
