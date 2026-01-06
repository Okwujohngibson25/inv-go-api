package service

import (
	"log"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/walletaddress/dto"
)

func (wa *WalletaddressService) FetchAdminWalletAddress() ([]dto.AdminWalletAddressResponseDTO, error) {
	wallets, err := wa.walletRepo.GetAdminWalletAddress()
	if err != nil {
		log.Printf("failed to fetch wallet: %v", err)
		return nil, apperrors.InternalError{Message: "failed to fetch wallet"}
	}

	return wallets, nil
}
