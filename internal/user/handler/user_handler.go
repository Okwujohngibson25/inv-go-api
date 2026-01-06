package handler

import (
	"net/http"
	"time"

	"github.com/coinserveringo/config"
	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/helpers"
	"github.com/coinserveringo/internal/response"
	"github.com/coinserveringo/internal/user/dto"
	"github.com/coinserveringo/internal/user/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserService *service.UserService
	cfg         *config.Config
}

func NewUserhandler(UserService *service.UserService, cfg *config.Config) *UserHandler {
	return &UserHandler{UserService: UserService, cfg: cfg}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account with email, fullname, phone, and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body dto.RegisterUserDTO true "User registration data"
// @Success 201 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /create [post]
func (h *UserHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	var req dto.RegisterUserDTO

	ok := helpers.BindJSON(c, &req)
	if !ok {
		return
	}

	if err := h.UserService.Create(ctx, req); err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.ApiResponse{
		Success: true,
		Message: "user account created successfully",
	})
}

// Login godoc
// @Summary Login User
// @Description Login user account using email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body dto.LoginUserDTO true "login user data"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var req dto.LoginUserDTO

	ok := helpers.BindJSON(c, &req)
	if !ok {
		return
	}

	accessToken, refreshToken, err := h.UserService.Login(ctx, req)
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	cookieMaxAge := int((7 * 24 * time.Hour).Seconds()) // 7 days
	c.SetCookie(
		"refresh_token",
		refreshToken,
		cookieMaxAge,
		"/",              // path
		"yourdomain.com", // domain (change in production)
		true,             // secure → HTTPS only
		true,             // httpOnly
	)

	// return access token
	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "logged in successfully",
		Data: gin.H{
			"access_token": accessToken,
		},
	})
}

// Refresh token godoc
// @Summary Refresh token
// @Description verify refresh token and send back a new access token
// @Tags Authentication
// @Accept json
// @Produce json
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /refresh [post]
func (h *UserHandler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusUnauthorized, response.ApiResponse{
			Success: false,
			Message: "authentication failed",
			Error:   "missing refresh token",
		})
		return
	}

	accessToken, err := h.UserService.RefreshAccessToken(ctx, refreshToken)
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "access token refreshed",
		Data: gin.H{
			"access_token": accessToken,
		},
	})
}

// ForgotPassword godoc
// @Summary Request password reset link
// @Description Sends a password reset link to the user's email
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body dto.ResetPasswordLinkDTO true "Password reset request data"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /forgot-password [post]
func (h *UserHandler) ForgotPassword(c *gin.Context) {
	ctx := c.Request.Context()
	var req dto.ResetPasswordLinkDTO

	ok := helpers.BindJSON(c, &req)
	if !ok {
		return
	}

	if err := h.UserService.ForgotPassword(ctx, &req); err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "if this email exists, a reset link has been sent",
	})
}

// ResetPassword godoc
// @Summary Reset user password
// @Description Reset user password from token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body dto.ResetPasswordDTO true "Password reset request data"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /reset-password [post]
func (h *UserHandler) ResetPassword(c *gin.Context) {
	ctx := c.Request.Context()
	var req dto.ResetPasswordDTO

	ok := helpers.BindJSON(c, &req)
	if !ok {
		return
	}

	if err := h.UserService.ResetPassword(ctx, &req); err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "password reset successful",
	})
}

// UpdatePassword godoc
// @Summary Update user password
// @Description update user password from profile settings
// @Tags User Action
// @Accept json
// @Produce json
// @Param input body dto.UpdatePasswordDTO true "Password Update data"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /user/update-password [patch]
func (h *UserHandler) UpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.MustGet("userid").(uuid.UUID)
	var req dto.UpdatePasswordDTO

	ok := helpers.BindJSON(c, &req)
	if !ok {
		return
	}

	req.UserID = userID

	if err := h.UserService.UpdatePassword(ctx, &req); err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "password update successful",
	})
}

// User Dashboard godoc
// @Summary User profile page data
// @Description Returns user profile page data for the logged-in user
// @Tags User Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /user/profiles [get]
func (h *UserHandler) Profile(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.MustGet("userid").(uuid.UUID)

	userData, err := h.UserService.GetUserByID(ctx, userID)
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "profile fetched successfully",
		Data:    userData,
	})
}
