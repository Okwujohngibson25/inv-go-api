package handler

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/kyc/dto"
	"github.com/coinserveringo/internal/kyc/service"
	"github.com/coinserveringo/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type KycHandler struct {
	KycService *service.KycService
}

func NewKycHandler(KycService *service.KycService) *KycHandler {
	return &KycHandler{KycService: KycService}
}

// KycUpload godoc
// @Summary Upload KYC documents
// @Description Upload user KYC ID front and back images
// @Tags User Action
// @Accept multipart/form-data
// @Produce json
// @Param id_front formData file true "ID Front Image"
// @Param id_back formData file true "ID Back Image"
// @Success 201 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /user/kyc/upload [post]
func (h *KycHandler) UploadKyc(ctx *gin.Context) {
	userID := ctx.MustGet("userid").(uuid.UUID)
	context := ctx.Request.Context()

	// --- Handle ID Front ---
	idFrontFile, err := ctx.FormFile("id_front")
	if err != nil {
		log.Printf("missing id front img: %v", err)
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: "validation failed",
			Error: gin.H{
				"id-front": "front ID image required",
			},
		})
		return
	}

	idFrontDst := filepath.Join("uploads/kyc", filepath.Base(idFrontFile.Filename))
	if err := ctx.SaveUploadedFile(idFrontFile, idFrontDst); err != nil {
		log.Printf("failed to save id front: %v", err)
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse{
			Success: false,
			Message: "Could not save kyc front image",
			Error:   "failed to store uploaded file",
		})
		return
	}

	// --- Handle ID back ---
	idBackFile, err := ctx.FormFile("id_back")
	if err != nil {
		log.Printf("missing id back image: %v", err)
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: "validation failed",
			Error: gin.H{
				"id-back": "back id image required",
			},
		})
		return
	}
	idBackDst := filepath.Join("uploads/kyc", filepath.Base(idBackFile.Filename))
	if err := ctx.SaveUploadedFile(idBackFile, idBackDst); err != nil {
		log.Printf("failed to save id back: %v", err)
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse{
			Success: false,
			Message: "validation failed",
			Error:   "failed to store uploaded file",
		})
		return
	}

	req := dto.KYCDTO{
		IDImageFront: idFrontDst,
		IDImageBack:  idBackDst,
		UserID:       userID,
	}

	if err := h.KycService.SaveKycDocument(context, &req); err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, response.ApiResponse{
		Success: true,
		Message: "kyc data uploaded successfully! awaiting approval",
	})
}
