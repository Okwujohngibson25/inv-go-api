package app

import (
	"context"
	"database/sql"

	"github.com/coinserveringo/config"
	"github.com/coinserveringo/internal/user/handler"
	"github.com/coinserveringo/internal/user/repository"
	"github.com/coinserveringo/internal/user/service"
	"github.com/coinserveringo/mail"

	depositHandler "github.com/coinserveringo/internal/deposit/handler"
	depositRepo "github.com/coinserveringo/internal/deposit/repository"
	depositService "github.com/coinserveringo/internal/deposit/service"

	withdrawalHandler "github.com/coinserveringo/internal/withdrawal/handler"
	withdrawalRepo "github.com/coinserveringo/internal/withdrawal/repository"
	withdrawalService "github.com/coinserveringo/internal/withdrawal/service"

	kycHandler "github.com/coinserveringo/internal/kyc/handler"
	kycRepo "github.com/coinserveringo/internal/kyc/repository"
	kycService "github.com/coinserveringo/internal/kyc/service"

	WalletAddressHandler "github.com/coinserveringo/internal/walletaddress/handler"
	WalletAddressRepo "github.com/coinserveringo/internal/walletaddress/repository"
	WalletAddressService "github.com/coinserveringo/internal/walletaddress/service"

	PlanHandler "github.com/coinserveringo/internal/plans/handler"
	PlanRepo "github.com/coinserveringo/internal/plans/repository"
	PlanService "github.com/coinserveringo/internal/plans/service"

	InvHandler "github.com/coinserveringo/internal/investment/handler"
	InvRepo "github.com/coinserveringo/internal/investment/repository"
	InvService "github.com/coinserveringo/internal/investment/service"

	DashboardHandler "github.com/coinserveringo/internal/pages/user/dashboard/handler"
	DashboardService "github.com/coinserveringo/internal/pages/user/dashboard/service"

	PackagesHandler "github.com/coinserveringo/internal/pages/user/packages/handler"
	PackagesService "github.com/coinserveringo/internal/pages/user/packages/service"

	DepositPageHandler "github.com/coinserveringo/internal/pages/user/deposit_page/handler"
	DepositPageService "github.com/coinserveringo/internal/pages/user/deposit_page/service"

	TransactionsHandler "github.com/coinserveringo/internal/pages/user/transactions/handler"
	TransactionsService "github.com/coinserveringo/internal/pages/user/transactions/service"

	AdminDashboardHandler "github.com/coinserveringo/internal/pages/admin/dashboard/handler"
	AdminDashboardService "github.com/coinserveringo/internal/pages/admin/dashboard/service"
)

type App struct {
	// Domain Handlers
	UserHandler          *handler.UserHandler
	DepositHandler       *depositHandler.DepositHandler
	WithdrawalHandler    *withdrawalHandler.WithdrawalHandler
	KycHandler           *kycHandler.KycHandler
	WalletAddressHandler *WalletAddressHandler.WalletAddressHandler
	PlanHandler          *PlanHandler.PlanHandler
	InvestmentHandler    *InvHandler.InvHandler

	// Page Handlers
	DashboardHandler      *DashboardHandler.DashboardHandler
	PackagesHandler       *PackagesHandler.PackagesHandler
	DepositPageHandler    *DepositPageHandler.DepositPageHandler
	TransactionsHandler   *TransactionsHandler.TransactionsHandler
	AdminDashboardHandler *AdminDashboardHandler.AdminDashboardHandler

	// Shared dependencies
	Mailer *mail.MailService
}

func NewApp(dbSqlc *sql.DB, mailer *mail.MailService, cfg *config.Config) *App {
	// ===== Repositories =====
	userRepo := repository.NewSqlcRepository(dbSqlc)
	depositRepo := depositRepo.NewSqlcRepository(dbSqlc)
	withdrawalRepo := withdrawalRepo.NewSqlcRepository(dbSqlc)
	kycRepo := kycRepo.NewSqlcRepository(dbSqlc)
	walletRepo := WalletAddressRepo.NewSqlcRepository(dbSqlc)
	planRepo := PlanRepo.NewSqlcRepository(dbSqlc)
	investmentRepo := InvRepo.NewSqlcRepository(dbSqlc)

	// ===== Services =====
	userService := service.NewUserService(userRepo, mailer, cfg)
	ctx := context.Background()
	userService.SeedAdmin(ctx, cfg) // seed admin
	depositService := depositService.NewDepositService(depositRepo, userRepo, mailer)
	withdrawalService := withdrawalService.NewWithdrawalService(withdrawalRepo, userService, mailer)
	kycService := kycService.NewKycService(kycRepo, userService, mailer)
	walletService := WalletAddressService.NewWalletaddressRepo(walletRepo)
	planService := PlanService.NewPlanService(planRepo)
	invService := InvService.NewInvService(investmentRepo, userService, planService)

	// Pages services
	dashboardService := DashboardService.NewDashboardService(userService, depositService, withdrawalService)
	packagesService := PackagesService.NewPackagesService(userService, planService, invService)
	depositPageService := DepositPageService.NewDepositPageService(userService, depositService, walletService)
	transactionsService := TransactionsService.NewTransactionsService(userService, depositService, withdrawalService)
	adminDashboardService := AdminDashboardService.NewAdminDashboardService(userService, depositService, withdrawalService)

	// ===== Handlers =====
	userHandler := handler.NewUserhandler(userService, cfg)
	depositHandler := depositHandler.NewDepositHandler(depositService)
	withdrawalHandler := withdrawalHandler.NewWithdrawalHandler(withdrawalService)
	kycHandler := kycHandler.NewKycHandler(kycService)
	walletHandler := WalletAddressHandler.NewWalletAddressHandler(walletService)
	planHandler := PlanHandler.NewPlanHandler(planService)
	investmentHandler := InvHandler.NewInvHandler(invService)

	dashboardHandler := DashboardHandler.NewDashboardHandler(dashboardService)
	packagesHandler := PackagesHandler.NewPackagesHandler(packagesService)
	depositPageHandler := DepositPageHandler.NewDepositPageHandler(depositPageService)
	transactionsHandler := TransactionsHandler.NewTransactionsHandler(transactionsService)
	adminDashboardHandler := AdminDashboardHandler.NewAdminDashboardHandler(adminDashboardService)

	return &App{
		// Domain Handlers
		UserHandler:          userHandler,
		DepositHandler:       depositHandler,
		WithdrawalHandler:    withdrawalHandler,
		KycHandler:           kycHandler,
		WalletAddressHandler: walletHandler,
		PlanHandler:          planHandler,
		InvestmentHandler:    investmentHandler,

		// Page Handlers
		DashboardHandler:      dashboardHandler,
		PackagesHandler:       packagesHandler,
		DepositPageHandler:    depositPageHandler,
		TransactionsHandler:   transactionsHandler,
		AdminDashboardHandler: adminDashboardHandler,

		// Shared dependencies
		Mailer: mailer,
	}
}
