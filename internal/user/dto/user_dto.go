package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type RegisterUserDTO struct {
	Email       string `json:"email" binding:"required,email"`
	Fullname    string `json:"fullname" binding:"required,min=2,max=50"`
	Username    string `json:"username" binding:"required,min=3,max=20"`
	Gender      string `json:"gender" binding:"required,oneof=male female other"`
	Country     string `json:"country" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,min=7,max=15"`
	Password    string `json:"password" binding:"required,min=6"`
}

type LoginUserDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDataDTO struct {
	ID             uuid.UUID       `json:"id"`
	Role           string          `json:"role"`
	Email          string          `json:"email"`
	Fullname       string          `json:"fullname"`
	Username       string          `json:"username"`
	Gender         string          `json:"gender"`
	Country        string          `json:"country"`
	PhoneNumber    string          `json:"phone-number"`
	WalletBalance  decimal.Decimal `json:"wallet-balance"`
	ProfitBalance  decimal.Decimal `json:"profit-balance"`
	InvestedAmount decimal.Decimal `json:"invested-amount"`
	ReferalBonus   decimal.Decimal `json:"referal-bonus"`
	ReferalLink    string          `json:"referal-link"`
	Verified       bool            `json:"verified"`
	Status         string          `json:"status"`
	CreatedAt      time.Time       `json:"created-at"`
	Password       string          `json:"-"`
}
