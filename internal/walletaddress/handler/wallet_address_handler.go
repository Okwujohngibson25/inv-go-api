package handler

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/helpers"
	"github.com/coinserveringo/internal/response"
	"github.com/coinserveringo/internal/walletaddress/dto"
	"github.com/coinserveringo/internal/walletaddress/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WalletAddressHandler struct {
	WalletAddressService *service.WalletaddressService
}

func NewWalletAddressHandler(WalletAddressService *service.WalletaddressService) *WalletAddressHandler {
	return &WalletAddressHandler{WalletAddressService: WalletAddressService}
}

// Add wallet godoc
// @Summary Add wallet
// @Description Add wallet as admin
// @Tags Admin Action
// @Accept multipart/form-data
// @Produce json
// @Param coin_logo formData file true "Coinlogo Image"
// @Param qr_code formData file true "QR code  Image"
// @Param coin formData string true "Coin name"
// @Param address formData string true "Wallet address"
// @Param network formData string true "Coin network"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/walletaddress [post]
func (h *WalletAddressHandler) Addwallet(ctx *gin.Context) {
	userID := ctx.MustGet("userid").(uuid.UUID)

	var req dto.AdminWalletAddressDTO

	ok := helpers.BindMultipart(ctx, &req)
	if !ok {
		return
	}

	// --- Handle Coin Logo---
	coinLogoFile, err := ctx.FormFile("coin_logo")
	if err != nil {
		log.Printf("missing coin logo img: %v", err)
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: "validation error",
			Error:   "coin logo img is required",
		})
		return
	}

	coinLogoDst := filepath.Join("uploads/wallet-images", filepath.Base(coinLogoFile.Filename))
	if err := ctx.SaveUploadedFile(coinLogoFile, coinLogoDst); err != nil {
		log.Printf("failed to save coin logo: %v", err)
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse{
			Success: false,
			Message: "Could not save coin logo img",
			Error:   "failed to store uploaded file",
		})
		return
	}

	// --- Handle QR Code ---
	qrCodeFile, err := ctx.FormFile("qr_code")
	if err != nil {
		log.Printf("missing id front img: %v", err)
		ctx.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: "validation error",
			Error:   "qr code is required",
		})
		return
	}

	qrCodeDst := filepath.Join("uploads/wallet-images", filepath.Base(qrCodeFile.Filename))
	if err := ctx.SaveUploadedFile(qrCodeFile, qrCodeDst); err != nil {
		log.Printf("failed to save ID front: %v", err)
		ctx.JSON(http.StatusInternalServerError, response.ApiResponse{
			Success: false,
			Message: "Could not save coin logo img",
			Error:   "failed to store uploaded file",
		})
		return
	}

	req.UserID = userID
	req.CoinLogo = coinLogoDst
	req.QrCode = qrCodeDst

	if err := h.WalletAddressService.AddAdminWalletAddress(&req); err != nil {
		apperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "wallet address added successfully",
	})
}

// Delete wallet godoc
// @Summary delete wallet
// @Description delete selected wallet
// @Tags Admin Action
// @Accept json
// @Produce json
// @Param id path string true "wallet ID"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/walletaddress/{id} [delete]
func (h *WalletAddressHandler) DeleteWallet(ctx *gin.Context) {
	walletID, ok := helpers.ParseUUIDParam(ctx, "id")
	if !ok {
		return
	}

	if err := h.WalletAddressService.DeleteWallet(walletID); err != nil {
		apperrors.HandleError(ctx, err)
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "wallet deleted successfully",
	})
}

// Get wallets godoc
// @Summary Get all admin wallets
// @Description Get all admin wallets
// @Tags Admin Dashboard
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Security BearerAuth
// @Router /admin/walletaddress [get]
func (h *WalletAddressHandler) Getwallet(ctx *gin.Context) {
	wallets, err := h.WalletAddressService.FetchAdminWalletAddress()
	if err != nil {
		apperrors.HandleError(ctx, err)
	}

	message := "wallets fetched successfully"
	if len(wallets) == 0 {
		message = "no wallet found"
	}

	ctx.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: message,
		Data:    wallets,
	})
}
