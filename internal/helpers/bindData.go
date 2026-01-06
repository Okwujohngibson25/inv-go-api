package helpers

import (
	"log"
	"net/http"

	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
)

func BindJSON(ctx *gin.Context, dest any) bool {
	if err := ctx.ShouldBindJSON(dest); err != nil {
		log.Printf("failed to bind request body: %v", err)

		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: "invalid request body",
			// Error:   "failed to parse request body",
			Error: err.Error(),
		})
		return false
	}
	return true
}

func BindMultipart(ctx *gin.Context, dest any) bool {
	if err := ctx.ShouldBind(dest); err != nil {
		log.Printf("failed to bind request body: %v", err)

		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: "invalid request body",
			// Error:   "failed to parse request body",
			Error: err.Error(),
		})
		return false
	}
	return true
}
