-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_deposits(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    "status" TEXT NOT NULL DEFAULT 'Pending',
    amount NUMERIC(18,6) NOT NULL,
    transaction_id TEXT NOT NULL,
    deposit_proof TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_deposits;
-- +goose StatementEnd
