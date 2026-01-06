-- +goose Up
-- +goose StatementBegin
CREATE TABLE admin_wallet_addresses(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    coin TEXT NOT NULL,
    coin_logo TEXT NOT NULL,
    "address" TEXT NOT NULL,
    qr_code TEXT NOT NULL,
    network_type TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()


)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS admin_wallet_addresses;
-- +goose StatementEnd
