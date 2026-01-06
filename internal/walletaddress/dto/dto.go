package dto

import (
	"time"

	"github.com/google/uuid"
)

type AdminWalletAddressDTO struct {
	UserID   uuid.UUID `json:"-"`
	Coin     string    `form:"coin" binding:"required"`
	CoinLogo string    `json:"-"`
	Address  string    `form:"address" binding:"required"`
	QrCode   string    `json:"-"`
	Network  string    `form:"network" binding:"required"`
}

type AdminWalletAddressResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Coin      string    `json:"coin"`
	CoinLogo  string    `json:"coin_logo"`
	Address   string    `json:"address"`
	QrCode    string    `json:"qr_code"`
	Network   string    `json:"network"`
	CreatedAt time.Time `json:"created_at"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
