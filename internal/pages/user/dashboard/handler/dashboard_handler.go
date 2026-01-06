package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/pages/user/dashboard/service"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DashboardHandler struct {
	dashboardService *service.DashboardService
}

func NewDashboardHandler(dashboardService *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{dashboardService: dashboardService}
}

// User Dashboard godoc
// @Summary User dashboard page
// @Description Returns dashboard data for the logged-in user
// @Tags User Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /user/dashboard [get]
func (h *DashboardHandler) GetDashboard(ctx *gin.Context) {
	context := ctx.Request.Context()
	userID := ctx.MustGet("userid").(uuid.UUID)
	dashboardData, err := h.dashboardService.GetDashboardData(context, userID)
	if err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "user dashboard data fetched successfully",
		Data:    dashboardData,
	})
}
