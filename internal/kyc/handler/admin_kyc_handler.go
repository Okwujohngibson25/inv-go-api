package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/helpers"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
)

// Fetch kyc godoc
// @Summary Fetch kyc
// @Description Fetch all users kyc data
// @Tags Admin Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/kyc [get]
func (h *KycHandler) GetAllKycData(ctx *gin.Context) {
	context := ctx.Request.Context()
	kycData, err := h.KycService.GetAllKyc(context)
	if err != nil {
		apperrors.HandleError(ctx, err)
		return
	}
	message := "kyc data fetched successfully"
	if len(kycData) == 0 {
		message = "no kyc data found"
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: message,
		Data:    kycData,
	})
}

// Fetch pending kyc godoc
// @Summary Fetch pending kyc
// @Description Fetch all pending users kyc data
// @Tags Admin Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/kyc/pending [get]
func (h *KycHandler) GetAllPendingKyc(ctx *gin.Context) {
	context := ctx.Request.Context()
	pendingKycData, err := h.KycService.GetAllPendingKyc(context)
	if err != nil {
		apperrors.HandleError(ctx, err)
		return
	}
	message := "pending kyc data fetched successfully"
	if len(pendingKycData) == 0 {
		message = "no pending kyc data found"
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: message,
		Data:    pendingKycData,
	})
}

// Approve user Kyc godoc
// @Summary Approve kyc
// @Description Approve user kyc
// @Tags Admin Action
// @Accept json
// @Produce json
// @Param id path string true "kyc ID"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/kyc/{id}/approve [put]
func (h *KycHandler) ApproveKyc(ctx *gin.Context) {
	context := ctx.Request.Context()
	kycID, ok := helpers.ParseUUIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := h.KycService.ApproveKyc(context, kycID); err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "kyc approved successfully",
	})
}

// Decline user Kyc godoc
// @Summary Decline kyc
// @Description Decline user kyc
// @Tags Admin Action
// @Accept json
// @Produce json
// @Param id path string true "kyc ID"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/kyc/{id}/decline [put]
func (h *KycHandler) Declinekyc(ctx *gin.Context) {
	context := ctx.Request.Context()
	kycID, ok := helpers.ParseUUIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := h.KycService.Declinekyc(context, kycID); err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "kyc declined successfully",
	})
}
