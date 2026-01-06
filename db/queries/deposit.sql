-- name: CreateNewDeposit :exec
INSERT INTO user_deposits(user_id, amount, transaction_id, deposit_proof) VALUES($1, $2, $3, $4);

-- name: GetAllPendingDepositByUserID :many
SELECT 
    d.*, 
    u.email,
    u.fullname,
    u.username
FROM user_deposits d
JOIN users u ON d.user_id = u.id
WHERE d.status = 'Pending'
  AND d.user_id = $1
ORDER BY d.created_at DESC;


-- name: GetAllDepositsByUserId :many
SELECT 
    d.*, 
    u.email,
    u.fullname,
    u.username
FROM user_deposits d
JOIN users u ON d.user_id = u.id
WHERE d.user_id = $1
ORDER BY d.created_at DESC;

-- name: SumAllDepositAmountByUserID :one
SELECT COALESCE(SUM(amount), 0)::NUMERIC(18,6) AS total_deposited FROM user_deposits WHERE "status" = 'Approved' AND id = $1;

-- name: SumApprovedDeposits :one
SELECT COALESCE(SUM(amount), 0)::NUMERIC(18,6) AS total_deposited FROM user_deposits WHERE "status" = 'Approved';

-- name: GetAllDeposits :many
SELECT 
    d.*, 
    u.email,
    u.fullname,
    u.username
FROM user_deposits d
JOIN users u ON d.user_id = u.id
ORDER BY d.created_at DESC;

-- name: GetAllPendingDeposit :many
SELECT 
    d.*, 
    u.email,
    u.fullname,
    u.username
FROM user_deposits d
JOIN users u ON d.user_id = u.id
WHERE "status" = 'Pending'
ORDER BY d.created_at DESC;

-- name: GetDepositByID :one
SELECT * FROM user_deposits WHERE id = $1;

-- name: ApproveDeposits :exec
UPDATE user_deposits SET "status" = 'Approved' WHERE id = $1;

-- name: DeclineDeposit :exec
UPDATE user_deposits SET "status" = 'Declined' WHERE id = $1;