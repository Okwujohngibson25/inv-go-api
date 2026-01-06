package dto

import (
	"time"

	"github.com/coinserveringo/internal/user/dto"
	"github.com/google/uuid"
)

type KYCDTO struct {
	IDImageFront string
	IDImageBack  string
	UserID       uuid.UUID
}

type UserKycDTO struct {
	ID           uuid.UUID       `json:"id"`
	User         dto.UserDataDTO `json:"user"`
	Status       string          `json:"status"`
	IDImageFront string          `json:"id_front"`
	IDImageBack  string          `json:"id_back"`
	CreatedAt    time.Time       `json:"created_at"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
