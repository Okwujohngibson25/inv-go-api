package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
)

// Fetch Active investment godoc
// @Summary Fetch Active investment
// @Description Fetch Active investment
// @Tags Admin Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/investment/active [get]
func (h *InvHandler) GetAllActiveinv(ctx *gin.Context) {
	context := ctx.Request.Context()
	investmentData, err := h.InvService.FetchAllActiveinv(context)
	if err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	message := "active investment fetched successfylly"
	if len(investmentData) == 0 {
		message = "no active investment found"
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: message,
		Data:    investmentData,
	})
}
