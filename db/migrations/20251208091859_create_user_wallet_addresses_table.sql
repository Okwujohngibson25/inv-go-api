-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_wallet_addresses(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    btc_btc_address TEXT,
    btc_bep20_address TEXT,
    eth_erc20_address TEXT,
    eth_bep20_address TEXT,
    usdt_trc20_address TEXT,
    usdt_bep20_address TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_wallet_addresses;
-- +goose StatementEnd
