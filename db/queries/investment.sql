
-- name: CreateInv :exec
INSERT INTO investments(user_id, plan_id, amount, end_date, expected_profit)
VALUES
($1, $2, $3, $4, $5);

-- name: GetActiveInvByUserId :many
SELECT * FROM investments WHERE id = $1 AND "status" = 'Ongoing' ORDER BY created_at DESC;

-- name: GetAllInvByUserId :many
SELECT * FROM investments WHERE id = $1 ORDER BY created_at DESC;

-- name: GetAllActiveInv :many
SELECT 
    d.*,
    u.email,
    u.fullname,
    u.username
FROM investments d 
JOIN users u ON d.user_id = u.id
WHERE "status" = 'Ongoing' ORDER BY d.created_at DESC;