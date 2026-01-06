package apperrors

import (
	"errors"
	"net/http"

	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
)

// ValidationError
type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

// ConflictError
type ConflictError struct {
	Message string
}

func (e ConflictError) Error() string {
	return e.Message
}

// NotFoundError
type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

// InternalError
type InternalError struct {
	Message string
}

func (e InternalError) Error() string {
	return e.Message
}

func HandleError(ctx *gin.Context, err error) {
	switch e := err.(type) {
	case ValidationError:
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: e.Message,
		})
	case ConflictError:
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: e.Message,
		})
	case NotFoundError:
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: e.Message,
		})
	case InternalError:
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: e.Message,
		})
	default:
		// fallback for unknown errors
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: "internal server error",
		})
	}
}

var (
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("conflict")
	ErrInvalid  = errors.New("invalid")
)
