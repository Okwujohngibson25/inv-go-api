package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/helpers"
	"github.com/coinserveringo/internal/plans/dto"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
)

// Approve Create plan godoc
// @Summary Create new plan
// @Description Create new inv plan
// @Tags Admin Action
// @Accept json
// @Produce json
// @Param input body dto.PlanDTO true "plan data"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/plans [post]
func (h *PlanHandler) Create(ctx *gin.Context) {
	context := ctx.Request.Context()
	var req dto.PlanDTO

	ok := helpers.BindJSON(ctx, &req)
	if !ok {
		return
	}

	if err := h.PlanService.CreatePlan(context, &req); err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, response.ApiResponse{
		Success: true,
		Message: "investment plan created successfully",
	})
}

// Delete investment plan godoc
// @Summary Delete plan
// @Description Delete investment plan
// @Tags Admin Action
// @Accept json
// @Produce json
// @Param id path string true "Plan ID"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/plan/{id} [delete]
func (h *PlanHandler) DeletePlan(ctx *gin.Context) {
	context := ctx.Request.Context()
	planID, ok := helpers.ParseUUIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := h.PlanService.DeletePlan(context, planID); err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "investment plan deleted successfully",
	})
}
