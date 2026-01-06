package repository

import (
	"context"
	"database/sql"

	"github.com/coinserveringo/internal/db"
	"github.com/coinserveringo/internal/walletaddress/dto"
	"github.com/google/uuid"
)

type WalletAddressRepo interface {
	AddAdminWalletAddress(req *dto.AdminWalletAddressResponseDTO) error
	GetAdminWalletAddress() ([]dto.AdminWalletAddressResponseDTO, error)
	DeleteWalletAddressByID(id uuid.UUID) error
}

type SqlcRepository struct {
	Queries *db.Queries
}

func NewSqlcRepository(conn *sql.DB) *SqlcRepository {
	return &SqlcRepository{Queries: db.New(conn)}
}

func (r *SqlcRepository) AddAdminWalletAddress(req *dto.AdminWalletAddressResponseDTO) error {
	params := db.AddAdminWalletAddressParams{
		UserID:      req.UserID,
		Coin:        req.Coin,
		CoinLogo:    req.CoinLogo,
		Address:     req.Address,
		QrCode:      req.QrCode,
		NetworkType: req.Network,
	}
	if err := r.Queries.AddAdminWalletAddress(context.Background(), params); err != nil {
		return err
	}
	return nil
}

func (r *SqlcRepository) DeleteWalletAddressByID(id uuid.UUID) error {
	if err := r.Queries.DeleteWalletAddressByID(context.Background(), id); err != nil {
		return err
	}

	return nil
}
