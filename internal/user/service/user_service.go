package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/coinserveringo/config"
	"github.com/coinserveringo/internal/apperrors"
	"github.com/coinserveringo/internal/db"
	"github.com/coinserveringo/internal/user/dto"
	"github.com/coinserveringo/internal/user/repository"
	"github.com/coinserveringo/mail"
	"github.com/coinserveringo/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CustomClaims struct {
	Email  string
	UserID uuid.UUID
	Role   string
	jwt.RegisteredClaims
}

type UserService struct {
	userRepo repository.UserRepository
	mailer   *mail.MailService
	cfg      *config.Config
}

func NewUserService(userRepo repository.UserRepository, mailer *mail.MailService, cfg *config.Config) *UserService {
	return &UserService{userRepo: userRepo, mailer: mailer, cfg: cfg}
}

func (u *UserService) Create(ctx context.Context, user dto.RegisterUserDTO) error {
	// hash incoming password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return apperrors.InternalError{Message: "failed to hash password"}
	}
	hashedPassword := string(password)
	referalLink := u.cfg.AppURL + "/" + uuid.New().String()

	params := db.CreateUserParams{
		Email:       user.Email,
		Role:        "user",
		Fullname:    user.Fullname,
		Username:    user.Username,
		Gender:      user.Gender,
		Country:     user.Country,
		PhoneNumber: user.PhoneNumber,
		Password:    string(hashedPassword),
		ReferalLink: referalLink,
	}

	// create user
	if err := u.userRepo.Create(ctx, &params); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return apperrors.ConflictError{Message: "email or phone number is already taken"}
		}
		log.Printf("failed to create new user: %v", err)
		return apperrors.InternalError{Message: "failed to create user account"}
	}

	// Send welcome email
	go func() {
		subject := "Test welcome Email! 🎉"
		data := map[string]interface{}{
			"Fullname":    user.Fullname,
			"CompanyName": u.cfg.CompanyName,
		}

		if err := u.mailer.SendMail(user.Email, subject, "welcome_email.html", data); err != nil {
			log.Printf("failed to send email to client: %v", err)
		}
	}()

	return nil
}

func (u *UserService) Login(ctx context.Context, user dto.LoginUserDTO) (signedAccess string, rawRefresh string, err error) {
	// find user in database using email and return user

	DbUser, err := u.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return "", "", apperrors.NotFoundError{Message: "invalid credentials"}
		}
		log.Printf("failed to find user using email: %v", err)
		return "", "", apperrors.InternalError{Message: "failed to find user using email"}
	}

	const StatusSuspended = "suspended"

	// Check Account Status
	if DbUser.Status == StatusSuspended {
		return "", "", apperrors.ValidationError{Message: "this account was suspended, contact customer support for help"}
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(DbUser.Password), []byte(user.Password))
	if err != nil {
		return "", "", apperrors.ValidationError{Message: "invalid email or password"}
	}
	// create jwt token

	// create access token
	accessClaims := CustomClaims{
		Email:  DbUser.Email,
		UserID: DbUser.ID,
		Role:   DbUser.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	// sign access token with secret key
	signedAccess, err = accessToken.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		log.Printf("unable to sign Jwt with secretKey: %v", err)
		return "", "", apperrors.InternalError{Message: "unable to sign Jwt with secretKey"}
	}

	// create refresh token
	refreshClaims := CustomClaims{
		Email:  DbUser.Email,
		UserID: DbUser.ID,
		Role:   DbUser.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 7 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// sign refresh token with secret key
	signedRefresh, err := refreshToken.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		log.Printf("unable to sign Jwt with secretKey: %v", err)
		return "", "", apperrors.InternalError{Message: "unable to sign Jwt with secretKey"}
	}

	// send the jwt token back to handler
	return signedAccess, signedRefresh, nil
}

func (u *UserService) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, u.cfg)
	if err != nil {
		log.Printf("failed to verify JWT token: %v", err)
		return "", apperrors.ValidationError{Message: "failed to verify JWT token"}
	}

	// Check user exists / active in DB
	user, err := u.userRepo.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return "", apperrors.ValidationError{Message: "invalid credentials"}
		}
		log.Printf("failed to find user by email: %v", err)
		return "", apperrors.InternalError{Message: "failed to find user by email"}
	}

	// Generate new access token
	newClaims := CustomClaims{
		Email:  user.Email,
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	newAccessJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	newAccessToken, err := newAccessJwt.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		log.Printf("failed to generate new access token: %v", err)
		return "", apperrors.InternalError{Message: "failed to generate new access token"}
	}

	return newAccessToken, nil
}

func (u *UserService) ForgotPassword(ctx context.Context, req *dto.ResetPasswordLinkDTO) error {
	// Check if user exists
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("failed to find user by email: %v", err)
	}

	// If user exists, create token and send email
	if user != nil {
		token := uuid.New().String()
		expiresAt := time.Now().Add(1 * time.Hour)

		reset := dto.UserPasswordReset{
			UserID:    user.ID,
			Token:     token,
			ExpiresAt: expiresAt,
		}

		if err := u.userRepo.CreateResetToken(ctx, &reset); err != nil {
			log.Printf("failed to create reset token: %v", err)
			return apperrors.InternalError{Message: "failed to create reset token"}
		}

		resetLink := fmt.Sprintf("%s/reset-password?token=%s", u.cfg.AppURL, token)

		go func() {
			subject := "Password Reset Request 🔐"
			data := map[string]interface{}{
				"Fullname":    user.Fullname,
				"ResetLink":   resetLink,
				"CompanyName": u.cfg.CompanyName,
			}

			if err := u.mailer.SendMail(user.Email, subject, "password_reset.html", data); err != nil {
				log.Printf("failed to send email to client: %v", err)
			}
		}()
	}

	return nil
}

func (u *UserService) ResetPassword(ctx context.Context, req *dto.ResetPasswordDTO) error {
	// confirm if token exist in our db
	tokenRecord, err := u.userRepo.GetTokenById(ctx, req.Token)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.NotFoundError{Message: "token not found"}
		}
		log.Printf("failed to find token by id: %v", err)
		return apperrors.InternalError{Message: "failed to find token"}

	}

	// check if token is expired
	if time.Now().After(tokenRecord.ExpiresAt) {
		// Delete expired token
		err = u.userRepo.DeleteToken(ctx, tokenRecord.TokenID)
		if err != nil {
			log.Printf("failed to delete token: %v", err)
			return apperrors.InternalError{Message: "failed to delete token"}
		}
		return apperrors.ValidationError{Message: "reset link has expired, please request a new password reset"}
	}

	// check if password and confirm password field matches
	if req.Password != req.ConfirmPassword {
		return apperrors.ValidationError{Message: "password and confirm password mismatch"}
	}

	// Hashpassword
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Printf("failed to hashpassword: %v", err)
		return apperrors.InternalError{Message: "password hashing failed"}
	}

	// Get User by ID
	user, err := u.userRepo.GetUserByID(ctx, tokenRecord.UserID)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return apperrors.InternalError{Message: "could not fetch user data"}
	}

	// update password
	if err := u.userRepo.UpdatePassword(ctx, user.ID, hashedPassword); err != nil {
		log.Printf("failed to update password: %v", err)
		return apperrors.InternalError{Message: "failed to update password"}
	}

	// send success email
	go func() {
		subject := "Password Reset Successful ✅"
		data := map[string]interface{}{
			"Fullname":    user.Fullname,
			"CompanyName": u.cfg.CompanyName,
		}

		if err := u.mailer.SendMail(user.Email, subject, "password_reset_success.html", data); err != nil {
			log.Printf("failed to send email to client: %v", err)
		}
	}()

	return nil
}

func (u *UserService) UpdatePassword(ctx context.Context, req *dto.UpdatePasswordDTO) error {
	// Get The User By Id
	user, err := u.userRepo.GetUserByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return apperrors.NotFoundError{Message: "user not found"}
		}
		log.Printf("failed to find user by id: %v", err)
		return apperrors.InternalError{Message: "could not fetch user data"}
	}

	// Compare if req Password and confirm password are same
	if req.NewPassword != req.ConfirmPassword {
		return apperrors.ValidationError{Message: "password and confirm password mismatch"}
	}

	// Compare useroldpassword to reqoldpassowrd
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		log.Printf("user %s failed password comparison: %v", user.Email, err)
		return apperrors.ValidationError{Message: "invalid credentials"}

	}

	// Hash New Password
	NewPasswordHashed, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		log.Printf("password hashing failed: %v", err)
		return apperrors.InternalError{Message: "password hashing failed"}
	}

	// Update the users table with new password
	if err := u.userRepo.UpdatePassword(ctx, user.ID, NewPasswordHashed); err != nil {
		log.Printf("update password failed: %v", err)
		return apperrors.InternalError{Message: "update password failed"}
	}

	// send success email
	go func() {
		subject := "Password Reset Successful ✅"
		data := map[string]interface{}{
			"Fullname":    user.Fullname,
			"CompanyName": u.cfg.CompanyName,
		}

		if err := u.mailer.SendMail(user.Email, subject, "password_reset_success.html", data); err != nil {
			log.Printf("failed to send email to client: %v", err)
		}
	}()

	return nil
}
