package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/pages/user/transactions/service"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionsHandler struct {
	transactionsService *service.TransactionsService
}

func NewTransactionsHandler(transactionsService *service.TransactionsService) *TransactionsHandler {
	return &TransactionsHandler{transactionsService: transactionsService}
}

// User Dashboard godoc
// @Summary User transaction page data
// @Description Returns transaction page data for the logged-in user
// @Tags User Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /user/transactions [get]
func (h *TransactionsHandler) GetTransactions(ctx *gin.Context) {
	context := ctx.Request.Context()
	userID := ctx.MustGet("userid").(uuid.UUID)
	transactionsPageData, err := h.transactionsService.GetTransactionData(context, userID)
	if err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "transaction page data fetched successfully",
		Data:    transactionsPageData,
	})
}
