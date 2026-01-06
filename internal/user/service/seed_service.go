package service

import (
	"context"
	"fmt"
	"log"

	"github.com/coinserveringo/config"
	"github.com/coinserveringo/internal/db"
	"github.com/coinserveringo/utils"
)

func (u *UserService) SeedAdmin(ctx context.Context, cfg *config.Config) {
	exists, err := u.userRepo.AdminExists(ctx)
	if err != nil {
		log.Println("Error checking admin existence:", err)
		return
	}

	if exists {
		log.Println("Admin already exists, skipping seeding.")
		return
	}

	hashedpassword, err := utils.HashPassword(cfg.AdminPassword)
	if err != nil {
		log.Println("failed to hash .env admin password:", err)
		return
	}

	// Create the super admin
	admin := db.CreateUserParams{
		Email:       cfg.AdminEmail,
		Role:        "admin",
		Fullname:    "System Administrator",
		Username:    "N/A",
		Gender:      "N/A",
		Country:     "N/A",
		PhoneNumber: "N/A",
		ReferalLink: "N/A",
		Password:    hashedpassword,
	}

	if err := u.userRepo.Create(ctx, &admin); err != nil {
		fmt.Println(err)
	}

	log.Println("Super admin created successfully")
}
