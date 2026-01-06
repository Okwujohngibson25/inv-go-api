package middleware

import (
	"log"
	"net/http"

	"github.com/coinserveringo/config"
	"github.com/coinserveringo/internal/response"
	"github.com/coinserveringo/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	config, _ := config.LoadConfig()
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ApiResponse{
			Success: false,
			Message: "unauthorized access",
			Error:   "No token in request",
		})
		return
	}

	claims, err := utils.VerifyToken(token, config)
	if err != nil {
		log.Printf("failed to verify token: %v", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ApiResponse{
			Success: false,
			Message: "unauthorized access",
			Error:   "failed to verify token",
		})
		return
	}
	ctx.Set("userid", claims.User_ID)
	ctx.Set("role", claims.Role)
	ctx.Next()
}
