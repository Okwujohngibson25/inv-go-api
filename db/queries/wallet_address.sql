
-- name: AddAdminWalletAddress :exec
INSERT INTO admin_wallet_addresses(user_id, coin, coin_logo, "address", qr_code, network_type) VALUES($1, $2, $3, $4, $5, $6);

-- name: DeleteWalletAddressByID :exec
DELETE FROM admin_wallet_addresses WHERE id = $1;

-- name: GetAdminWalletAddress :many
SELECT * FROM admin_wallet_addresses;
