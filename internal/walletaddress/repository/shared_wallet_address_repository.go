package repository

import (
	"context"

	"github.com/coinserveringo/internal/walletaddress/dto"
)

func (r *SqlcRepository) GetAdminWalletAddress() ([]dto.AdminWalletAddressResponseDTO, error) {
	var walletData []dto.AdminWalletAddressResponseDTO
	wallets, err := r.Queries.GetAdminWalletAddress(context.Background())
	if err != nil {
		return nil, err
	}

	for _, value := range wallets {
		walletData = append(walletData, dto.AdminWalletAddressResponseDTO{
			ID:        value.ID,
			Coin:      value.Coin,
			CoinLogo:  value.CoinLogo,
			Address:   value.Address,
			QrCode:    value.QrCode,
			Network:   value.NetworkType,
			CreatedAt: value.CreatedAt,
		})
	}

	return walletData, nil
}
