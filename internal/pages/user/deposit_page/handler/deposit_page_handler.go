package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/pages/user/deposit_page/service"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DepositPageHandler struct {
	depositPageService *service.DepositPageService
}

func NewDepositPageHandler(depositPageService *service.DepositPageService) *DepositPageHandler {
	return &DepositPageHandler{depositPageService: depositPageService}
}

// User Dashboard godoc
// @Summary User deposit page data
// @Description Returns deposit page data for the logged-in user
// @Tags User Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /user/deposits [get]
func (h *DepositPageHandler) GetdepositData(ctx *gin.Context) {
	context := ctx.Request.Context()
	userId := ctx.MustGet("userid").(uuid.UUID)
	depositPageData, err := h.depositPageService.GetDepositPackageData(context, userId)
	if err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "deposit page data fetched successfully",
		Data:    depositPageData,
	})
}
