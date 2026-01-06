package routes

import (
	"github.com/coinserveringo/internal/app"
	"github.com/coinserveringo/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine, myApp *app.App) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Public Routes
	router.POST("/create", myApp.UserHandler.Register)
	router.POST("/login", myApp.UserHandler.Login)
	router.POST("/forgot-password", myApp.UserHandler.ForgotPassword)
	router.POST("/reset-password", myApp.UserHandler.ResetPassword)
	router.POST("/refresh", myApp.UserHandler.Refresh)

	// Protected User Routes
	auth := router.Group("/user", middleware.Authenticate)
	auth.GET("/dashboard", myApp.DashboardHandler.GetDashboard)
	auth.GET("/deposits", myApp.DepositPageHandler.GetdepositData)
	auth.GET("/profiles", myApp.UserHandler.Profile)
	auth.GET("/packages", myApp.PackagesHandler.GetPackages)
	auth.GET("/transactions", myApp.TransactionsHandler.GetTransactions)

	auth.POST("/deposits", myApp.DepositHandler.CreateNewDeposit)
	auth.POST("/withdrawal", myApp.WithdrawalHandler.CreateNewWithdrawal)
	auth.POST("/kyc/upload", myApp.KycHandler.UploadKyc)
	auth.PATCH("/update-password", myApp.UserHandler.UpdatePassword)
	auth.POST("/investment/:id", myApp.InvestmentHandler.Joinplan)

	// Protected Admin Routes
	admin := router.Group("/admin", middleware.Authenticate, middleware.Authorization("admin"))
	admin.GET("/dashboard", myApp.AdminDashboardHandler.GetAdminDashboardData)
	admin.GET("/users", myApp.UserHandler.FetchAllUsersData)
	admin.GET("/users/:id", myApp.UserHandler.FetchUserDataById)
	admin.GET("/deposit", myApp.DepositHandler.GetAllDeposit)
	admin.GET("/deposit/pending", myApp.DepositHandler.GetAllPendingDeposit)
	admin.GET("/withdrawal", myApp.WithdrawalHandler.GetAllWithdrawal)
	admin.GET("/withdrawal/pending", myApp.WithdrawalHandler.GetAllPendingWithdrawal)
	admin.GET("/kyc", myApp.KycHandler.GetAllKycData)
	admin.GET("/kyc/pending", myApp.KycHandler.GetAllPendingKyc)
	admin.GET("/investment/active", myApp.InvestmentHandler.GetAllActiveinv)

	admin.PUT("/users/:id/update", myApp.UserHandler.UpdateUserAcc)
	admin.DELETE("/users/:id/delete", myApp.UserHandler.DeleteUser)
	admin.PATCH("/users/:id/status", myApp.UserHandler.UpdateAccountStatus)
	admin.PUT("/deposit/:id/approve", myApp.DepositHandler.ApproveDeposit)
	admin.PUT("/deposit/:id/decline", myApp.DepositHandler.DeclineDeposit)
	admin.PUT("/withdrawal/:id/approve", myApp.WithdrawalHandler.ApproveWithdrawal)
	admin.PUT("/withdrawal/:id/decline", myApp.WithdrawalHandler.DeclineWithdrawal)
	admin.PUT("/kyc/:id/approve", myApp.KycHandler.ApproveKyc)
	admin.PUT("/kyc/:id/decline", myApp.KycHandler.Declinekyc)
	admin.POST("/plans", myApp.PlanHandler.Create)
	admin.DELETE("/plans/:id", myApp.PlanHandler.DeletePlan)
	admin.POST("/walletaddress", myApp.WalletAddressHandler.Addwallet)
	admin.GET("/walletaddress", myApp.WalletAddressHandler.Getwallet)
	admin.DELETE("/walletaddress/:id", myApp.WalletAddressHandler.DeleteWallet)
}
