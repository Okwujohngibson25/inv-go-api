-- name: CreateResetToken :exec
INSERT INTO user_password_resets(user_id, token, expires_at) VALUES($1, $2, $3);

-- name: GetTokenById :one
SELECT * FROM user_password_resets WHERE token = $1;

-- name: DeleteTokenByID :exec
DELETE FROM user_password_resets WHERE id = $1;