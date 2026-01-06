package service

import (
	"log"
	"strings"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/walletaddress/dto"
	"github.com/coinserveringo/internal/walletaddress/repository"
	"github.com/google/uuid"
)

type WalletaddressService struct {
	walletRepo repository.WalletAddressRepo
}

func NewWalletaddressRepo(walletRepo repository.WalletAddressRepo) *WalletaddressService {
	return &WalletaddressService{walletRepo: walletRepo}
}

func (wa *WalletaddressService) AddAdminWalletAddress(req *dto.AdminWalletAddressDTO) error {
	wallet := dto.AdminWalletAddressResponseDTO{
		UserID:   req.UserID,
		Coin:     req.Coin,
		CoinLogo: req.CoinLogo,
		Address:  req.Address,
		QrCode:   req.QrCode,
		Network:  req.Network,
	}

	if err := wa.walletRepo.AddAdminWalletAddress(&wallet); err != nil {

		if strings.Contains(err.Error(), "duplicate key") {
			return apperrors.ConflictError{Message: "coin name or wallet address already exists"}
		}
		log.Printf("failed to add admin wallet address: %v", err)
		return apperrors.InternalError{Message: "failed to add admin wallet"}
	}

	return nil
}

func (wa *WalletaddressService) DeleteWallet(id uuid.UUID) error {
	if err := wa.walletRepo.DeleteWalletAddressByID(id); err != nil {
		log.Printf("failed to delete wallet: %v", err)
		return apperrors.InternalError{Message: "failed to delete wallet"}
	}

	return nil
}
