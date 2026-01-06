package helpers

import (
	"log"
	"net/http"

	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ParseUUIDParam(ctx *gin.Context, param string) (uuid.UUID, bool) {
	idStr := ctx.Param(param)
	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("failed to parse %s as UUID: %v", param, err)
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: "validation error",
			Error:   "invalid " + param,
		})
		return uuid.Nil, false
	}
	return parsedID, true
}
