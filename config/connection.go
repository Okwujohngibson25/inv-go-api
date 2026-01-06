package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg *Config) *sql.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.SSLMode,
	)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	log.Println("connected to database")

	_, err = conn.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	if err != nil {
		log.Fatalf("failed to create uuid-ossp extension: %v", err)
	}

	return conn
}
