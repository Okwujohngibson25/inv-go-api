package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/helpers"
	"github.com/coinserveringo/internal/response"
	"github.com/coinserveringo/internal/withdrawal/dto"
	"github.com/coinserveringo/internal/withdrawal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WithdrawalHandler struct {
	WithdrawalService *service.WithdrawalService
}

func NewWithdrawalHandler(WithdrawalService *service.WithdrawalService) *WithdrawalHandler {
	return &WithdrawalHandler{WithdrawalService: WithdrawalService}
}

// CreateWithdrawal godoc
// @Summary New withdrawal
// @Description Create a new withdrawal for user
// @Tags User Action
// @Accept json
// @Produce json
// @Param input body dto.WithdrawalDTO true "withdrawal data"
// @Success 201 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /user/withdrawal [post]
func (h *WithdrawalHandler) CreateNewWithdrawal(c *gin.Context) {
	ctx := c.Request.Context()
	// Get user ID from gin.context
	userID := c.MustGet("userid").(uuid.UUID)
	var req dto.WithdrawalDTO

	ok := helpers.BindJSON(c, &req)
	if !ok {
		return
	}

	// Setting UserID
	req.UserID = userID

	if err := h.WithdrawalService.CreateNewWithdrawal(ctx, &req); err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "withdrawal created successfully! Awaiting admin approval",
	})
}
