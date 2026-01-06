

-- name: CreateUser :exec
INSERT INTO users (email, "role", fullname, username, gender, country, phone_number, referal_link, password) 
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUserBalanceByID :exec
UPDATE users SET wallet_balance = $1 WHERE id = $2;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateUserPassword :exec
UPDATE users SET password = $1 WHERE id = $2;

-- name: UpdateKycStatus :exec
UPDATE users SET verified = 'TRUE' WHERE id = $1;

-- name: AdminExists :one
SELECT COUNT(*) FROM users WHERE "role" = 'admin';

-- name: GetAllUsers :many
SELECT * FROM users WHERE "role" = 'user';

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateAccountStatus :exec
UPDATE users SET status = $1 WHERE id = $2;

-- name: CountUsers :one
SELECT COUNT(*) FROM users WHERE "role" = 'user';

-- name: UpdateUserAcc :exec
UPDATE users SET wallet_balance = $1, profit_balance = $2, invested_amount = $3, referal_bonus = $4 WHERE id = $5;