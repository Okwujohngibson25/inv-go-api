package handler

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/deposit/dto"
	"github.com/coinserveringo/internal/deposit/service"
	"github.com/coinserveringo/internal/helpers"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DepositHandler struct {
	DepositService *service.DepositService
}

func NewDepositHandler(DepositService *service.DepositService) *DepositHandler {
	return &DepositHandler{DepositService: DepositService}
}

// CreateDeposit godoc
// @Summary New deposit
// @Description Create a new deposit for a user
// @Tags User Action
// @Accept multipart/form-data
// @Produce json
// @Param amount formData number true "Deposit Amount"
// @Param deposit_proof formData file true "Deposit Proof Image"
// @Success 201 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /user/deposits [post]
func (h *DepositHandler) CreateNewDeposit(ctx *gin.Context) {
	context := ctx.Request.Context()
	userID := ctx.MustGet("userid").(uuid.UUID)
	var req dto.DepositDTO

	if !helpers.BindMultipart(ctx, &req) {
		return
	}

	// save depoit img
	idDepositFile, err := ctx.FormFile("deposit_proof")
	if err != nil {
		log.Printf("missing deposit proof: %v", err)
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: "validation failed",
			Error: gin.H{
				"deposit_proof": "deposit proof is required",
			},
		})
		return
	}

	idFrontDst := filepath.Join("uploads/depositProof", filepath.Base(idDepositFile.Filename))
	if err := ctx.SaveUploadedFile(idDepositFile, idFrontDst); err != nil {
		log.Printf("failed to save deposit proof: %v", err)
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse{
			Success: false,
			Message: "Could not save deposit proof",
			Error:   "failed to store uploaded file",
		})
		return
	}

	NewDeposit := dto.DepositDTO{
		UserID:       userID,
		Amount:       req.Amount,
		DepositProof: idFrontDst,
	}

	if err := h.DepositService.CreateNewDeposit(context, NewDeposit); err != nil {
		apperrors.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, response.ApiResponse{
		Success: true,
		Message: "Deposit created successfully! Awaiting admin approval",
	})
}
