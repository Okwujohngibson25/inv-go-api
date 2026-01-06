package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppENV     string
	AppURL     string
	DBHost     string
	DBUser     string
	DBName     string
	DBPassword string
	DBPort     string
	SSLMode    string

	AdminEmail    string
	AdminPassword string

	JWTSecret string

	SMTPHost       string
	SMTPPort       string
	SMTPEncryption string
	SMTPUser       string
	SMTPPass       string
	SMTPFrom       string

	CompanyName string
	ResetDB     string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		AppENV:     os.Getenv("APP_ENV"),
		AppURL:     os.Getenv("APP_URL"),
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBPort:     os.Getenv("PORT"),
		SSLMode:    os.Getenv("SSL_MODE"),

		AdminEmail:    os.Getenv("ADMIN_EMAIL"),
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),

		JWTSecret: os.Getenv("JWT_SECRET"),

		SMTPHost:       os.Getenv("SMTP_HOST"),
		SMTPPort:       os.Getenv("SMTP_PORT"),
		SMTPEncryption: os.Getenv("SMTP_ENCRYPTION"),
		SMTPUser:       os.Getenv("SMTP_USER"),
		SMTPPass:       os.Getenv("SMTP_PASS"),
		SMTPFrom:       os.Getenv("SMTP_FROM"),

		CompanyName: os.Getenv("COMPANY_NAME"),
		ResetDB:     os.Getenv("RESET_DB"),
	}

	return cfg, nil
}
