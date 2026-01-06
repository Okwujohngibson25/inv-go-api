-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_withdrawals(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    "status" TEXT NOT NULL DEFAULT 'Pending',
    amount NUMERIC(18,6) NOT NULL,
    transaction_id TEXT NOT NULL,
    coin_type TEXT NOT NULL,
    wallet_address TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_withdrawals;
-- +goose StatementEnd
