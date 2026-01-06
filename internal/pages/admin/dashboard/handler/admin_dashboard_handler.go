package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/pages/admin/dashboard/service"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
)

type AdminDashboardHandler struct {
	AdminDashboardService *service.AdminDashboardService
}

func NewAdminDashboardHandler(AdminDashboardService *service.AdminDashboardService) *AdminDashboardHandler {
	return &AdminDashboardHandler{AdminDashboardService: AdminDashboardService}
}

// Admin dashboard data godoc
// @Summary Admin dashboard data
// @Description Returns dashboard data for admin dashboard
// @Tags Admin Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/dashboard [get]
func (h *AdminDashboardHandler) GetAdminDashboardData(ctx *gin.Context) {
	context := ctx.Request.Context()
	adminDashboardData, err := h.AdminDashboardService.GetAdminDashboardData(context)
	if err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "admin dashboard data fetched successfully",
		Data:    adminDashboardData,
	})
}
