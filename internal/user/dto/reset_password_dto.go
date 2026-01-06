package dto

import (
	"time"

	"github.com/google/uuid"
)

type ResetPasswordDTO struct {
	Token           string `json:"token" form:"token" binding:"required"`
	Password        string `json:"password" form:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
}

type UserPasswordReset struct {
	UserID    uuid.UUID
	Token     string
	TokenID   uuid.UUID
	ExpiresAt time.Time
}
