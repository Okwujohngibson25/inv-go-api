package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/pages/user/packages/service"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PackagesHandler struct {
	PackagesService *service.PackagesService
}

func NewPackagesHandler(PackagesService *service.PackagesService) *PackagesHandler {
	return &PackagesHandler{PackagesService: PackagesService}
}

// User Dashboard godoc
// @Summary User investment page data
// @Description Returns investment page data for the logged-in user
// @Tags User Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /user/packages [get]
func (h *PackagesHandler) GetPackages(ctx *gin.Context) {
	context := ctx.Request.Context()
	userId := ctx.MustGet("userid").(uuid.UUID)
	packagesData, err := h.PackagesService.GetPackagesDate(context, userId)
	if err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "package page data fetched successfully",
		Data:    packagesData,
	})
}
