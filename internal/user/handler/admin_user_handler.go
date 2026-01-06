package handler

import (
	"net/http"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/helpers"
	"github.com/coinserveringo/internal/response"
	"github.com/coinserveringo/internal/user/dto"
	"github.com/gin-gonic/gin"
)

// Fetch all users godoc
// @Summary Fetch all users
// @Description Fetch all user data
// @Tags Admin Dashboard
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/users [get]
func (h *UserHandler) FetchAllUsersData(c *gin.Context) {
	ctx := c.Request.Context()
	usersData, err := h.UserService.FetchAllUsers(ctx)
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	message := "users fetched successfully"

	if len(usersData) == 0 {
		message = "no user record found"
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: message,
		Data:    usersData,
	})
}

// Fetch user godoc
// @Summary Fetch user
// @Description Fetch user data
// @Tags Admin Dashboard
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/users/{id} [get]
func (h *UserHandler) FetchUserDataById(c *gin.Context) {
	ctx := c.Request.Context()
	UserID, ok := helpers.ParseUUIDParam(c, "id")
	if !ok {
		return
	}

	userData, err := h.UserService.GetUserByID(ctx, UserID)
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "user fetched successfully",
		Data:    userData,
	})
}

// Update user account godoc
// @Summary Update user account
// @Description update user account as an admin
// @Tags Admin Action
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param input body dto.UpdateUserBalanceDTO true "Update user data"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/users/{id}/update [put]
func (h *UserHandler) UpdateUserAcc(c *gin.Context) {
	ctx := c.Request.Context()
	userID, ok := helpers.ParseUUIDParam(c, "id")
	if !ok {
		return
	}

	var req dto.UpdateUserBalanceDTO

	ok = helpers.BindJSON(c, &req)
	if !ok {
		return
	}

	req.Userid = userID

	if err := h.UserService.UpdateUserAcc(ctx, &req); err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "user account updated successfully",
	})
}

// Delete user account godoc
// @Summary Delete user account
// @Description delete user account as an admin
// @Tags Admin Action
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/users/{id}/delete [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	userID, ok := helpers.ParseUUIDParam(c, "id")
	if !ok {
		return
	}

	if err := h.UserService.DeleteUser(ctx, userID); err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "user deleted successfully",
	})
}

// Update user account status godoc
// @Summary Update user account status
// @Description update user account status as an admin
// @Tags Admin Action
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/users/{id}/status [patch]
func (h *UserHandler) UpdateAccountStatus(c *gin.Context) {
	ctx := c.Request.Context()
	userID, ok := helpers.ParseUUIDParam(c, "id")
	if !ok {
		return
	}

	if err := h.UserService.UpdateAccountStatus(ctx, userID); err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "account updated successfully",
	})
}
