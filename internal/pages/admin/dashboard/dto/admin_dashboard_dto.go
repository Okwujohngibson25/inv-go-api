package dto

import (
	depositdatadto "github.com/coinserveringo/internal/deposit/dto"
	userwithdrawaldto "github.com/coinserveringo/internal/withdrawal/dto"
	"github.com/shopspring/decimal"
)

type AdminDashboardDTO struct {
	MembersCount                   int64                                 `json:"members_count"`
	SumAmountOfApprovedWithdrawals decimal.Decimal                       `json:"sum_withdrawal_amount"`
	SumAmountOfApprovedDeposits    decimal.Decimal                       `json:"sum_deposit_amount"`
	AllPendingDepoists             []depositdatadto.UserDepositDTO       `json:"pending_deposits"`
	AllPendingWithdrawals          []userwithdrawaldto.UserWithdrawalDTO `json:"pending_withdrawals"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
