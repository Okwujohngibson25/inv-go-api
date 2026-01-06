package middleware

import (
	"net/http"

	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
)

func Authorization(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusForbidden, response.ApiResponse{
				Success: false,
				Message: "unauthorized access",
				Error:   "role not found in token",
			})
			return
		}

		if role != requiredRole {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ApiResponse{
				Success: false,
				Message: "forbidden: insufficient permissions",
				Error:   "role mismatch",
			})
			return
		}
		ctx.Next()
	}
}
