package dto

import (
	depositdatadto "github.com/coinserveringo/internal/deposit/dto"
	withdrawalresponsedto "github.com/coinserveringo/internal/withdrawal/dto"
)

type TransactionsDTO struct {
	AllDeposits    []depositdatadto.UserDepositDTO           `json:"all_deposits"`
	AllWithdrawals []withdrawalresponsedto.UserWithdrawalDTO `json:"all_withdrawals"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
