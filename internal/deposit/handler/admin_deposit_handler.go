package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/helpers"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
)

// Fetch deposits godoc
// @Summary Fetch deposits
// @Description Fetch all users deposits
// @Tags Admin Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/deposit [get]
func (h *DepositHandler) GetAllDeposit(ctx *gin.Context) {
	context := ctx.Request.Context()
	deposits, err := h.DepositService.GetAllDeposits(context)
	if err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	message := "Deposits fetched successfully"
	if len(deposits) == 0 {
		message = "No deposits found"
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: message,
		Data:    deposits,
	})
}

// Fetch pending deposits godoc
// @Summary Fetch pending deposits
// @Description Fetch all pending users deposits
// @Tags Admin Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/deposit/pending [get]
func (h *DepositHandler) GetAllPendingDeposit(ctx *gin.Context) {
	context := ctx.Request.Context()
	depositData, err := h.DepositService.GetAllPendingDeposit(context)
	if err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	message := "Pending deposits fetched successfully"
	if len(depositData) == 0 {
		message = "No pending deposits found"
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: message,
		Data:    depositData,
	})
}

// Approve user deposits godoc
// @Summary Approve deposit
// @Description Approve user deposits
// @Tags Admin Action
// @Accept json
// @Produce json
// @Param id path string true "Deposit ID"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/deposit/{id}/approve [put]
func (h *DepositHandler) ApproveDeposit(ctx *gin.Context) {
	context := ctx.Request.Context()
	depositID, ok := helpers.ParseUUIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := h.DepositService.ApproveDeposit(context, depositID); err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "deposit approved successfully",
	})
}

// Decline user deposits godoc
// @Summary Decline deposit
// @Description Decline user deposits
// @Tags Admin Action
// @Accept json
// @Produce json
// @Param id path string true "id" Example(550e8400-e29b-41d4-a716-446655440000)
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/deposit/{id}/decline [put]
func (h *DepositHandler) DeclineDeposit(ctx *gin.Context) {
	context := ctx.Request.Context()
	depositID, ok := helpers.ParseUUIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := h.DepositService.DeclineDeposit(context, depositID); err != nil {
		apperrors.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "Deposit declined successfully",
	})
}
