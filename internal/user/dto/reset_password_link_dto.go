package dto

type ResetPasswordLinkDTO struct {
	Email string `json:"email" form:"email" binding:"required"`
}
