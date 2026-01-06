-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL UNIQUE,
    "role" TEXT NOT NULL,
    fullname VARCHAR(100) NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    gender TEXT NOT NULL,
    country VARCHAR(50) NOT NULL,
    phone_number VARCHAR(15) NOT NULL UNIQUE,
    wallet_balance NUMERIC(18,6) NOT NULL DEFAULT 0,
    profit_balance NUMERIC(18,6) NOT NULL DEFAULT 0,
    invested_amount NUMERIC(18,6) NOT NULL DEFAULT 0,
    referal_bonus NUMERIC(18,6) NOT NULL DEFAULT 0,
    referal_link TEXT NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT FALSE,
    "status" TEXT NOT NULL DEFAULT 'active',
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd
