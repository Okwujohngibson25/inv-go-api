package dto

import "github.com/google/uuid"

type UpdatePasswordDTO struct {
	UserID          uuid.UUID `json:"-"`
	OldPassword     string    `json:"old_password"  binding:"required"`
	NewPassword     string    `json:"new_password"  binding:"required"`
	ConfirmPassword string    `json:"confirm_password"  binding:"required"`
}
