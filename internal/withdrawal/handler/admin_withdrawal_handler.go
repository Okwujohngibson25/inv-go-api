package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/helpers"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
)

// Fetch withdrawals godoc
// @Summary Fetch withdrawals
// @Description Fetch all users withdrawals
// @Tags Admin Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/withdrawal [get]
func (h *WithdrawalHandler) GetAllWithdrawal(c *gin.Context) {
	ctx := c.Request.Context()
	withdrawalData, err := h.WithdrawalService.GetAllWithdrawal(ctx)
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}
	message := "withdrawals fetched successfully"

	if len(withdrawalData) == 0 {
		message = "no withdrawal found"
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: message,
		Data:    withdrawalData,
	})
}

// Fetch pending withdrawals godoc
// @Summary Fetch pending withdrawals
// @Description Fetch all pending user withdrawals
// @Tags Admin Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/withdrawal/pending [get]
func (h *WithdrawalHandler) GetAllPendingWithdrawal(c *gin.Context) {
	ctx := c.Request.Context()
	pendingWithdrawalData, err := h.WithdrawalService.GetAllPendingWithdrawal(ctx)
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	message := "pending withdrawal fetched successfully"
	if len(pendingWithdrawalData) == 0 {
		message = "no pending withdrawal found"
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: message,
		Data:    pendingWithdrawalData,
	})
}

// Approve user withdrawals godoc
// @Summary Approve withdrawal
// @Description Approve user withdrawal
// @Tags Admin Action
// @Accept json
// @Produce json
// @Param id path string true "withdrawal ID"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/withdrawal/{id}/approve [put]
func (h *WithdrawalHandler) ApproveWithdrawal(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	withdrawalID, ok := helpers.ParseUUIDParam(c, idStr)
	if !ok {
		return
	}

	if err := h.WithdrawalService.ApproveWithdrawal(ctx, withdrawalID); err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "withdrawal approved successfully",
	})
}

// Decline user withdrawals godoc
// @Summary Decline withdrawal
// @Description Decline user withdrawal
// @Tags Admin Action
// @Accept json
// @Produce json
// @Param id path string true "withdrawal ID"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/withdrawal/{id}/decline [put]
func (h *WithdrawalHandler) DeclineWithdrawal(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	withdrawalID, ok := helpers.ParseUUIDParam(c, idStr)
	if !ok {
		return
	}

	if err := h.WithdrawalService.DeclineWithdrawal(ctx, withdrawalID); err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "withdrawal request declined successfully",
	})
}
