package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/helpers"
	"github.com/coinserveringo/internal/investment/dto"
	"github.com/coinserveringo/internal/investment/service"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InvHandler struct {
	InvService *service.InvService
}

func NewInvHandler(InvService *service.InvService) *InvHandler {
	return &InvHandler{InvService: InvService}
}

// SubscribePlan godoc
// @Summary Subscribe to an investment plan
// @Description User subscribing to an investment plan
// @Tags User Action
// @Accept json
// @Produce json
// @Param id path string true "Investment Plan ID" Format(uuid)
// @Param input body dto.InvestmentDTO true "join plan data"
// @Success 201 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /user/investment/{id} [post]
func (h *InvHandler) Joinplan(ctx *gin.Context) {
	context := ctx.Request.Context()
	userID := ctx.MustGet("userid").(uuid.UUID)

	planID, ok := helpers.ParseUUIDParam(ctx, "id")
	if !ok {
		return
	}

	var req dto.InvestmentDTO

	ok = helpers.BindJSON(ctx, &req)
	if !ok {
		return
	}

	req.PlanID = planID
	req.UserID = userID

	if err := h.InvService.Joininv(context, &req); err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, response.ApiResponse{
		Success: true,
		Message: "successfully subscribed to plan",
	})
}
