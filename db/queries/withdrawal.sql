
-- name: CreateNewWithdrawal :exec
INSERT INTO user_withdrawals(user_id, amount, transaction_id, coin_type, wallet_address) VALUES($1, $2, $3, $4, $5);

-- name: ApproveWithdrawal :exec
UPDATE user_withdrawals SET "status" = 'Approved' WHERE user_id = $1;

-- name: DeclineWithdrawal :exec
UPDATE user_withdrawals SET "status" = 'Declined' WHERE user_id = $1; 


-- name: GetAllWithdrawal :many
SELECT d.*,
        u.username,
        u.email,
        u.fullname
FROM user_withdrawals d
JOIN users u ON d.user_id = u.id ORDER BY d.created_at DESC;

-- name: GetAllPendingWithdrawal :many
SELECT d.*,
        u.username,
        u.email,
        u.fullname
FROM user_withdrawals d 
JOIN users u ON d.user_id = u.id WHERE d."status" = 'Pending';

-- name: GetWithdrawalById :one
SELECT d.*,
        u.username,
        u.email,
        u.fullname
FROM user_withdrawals d
JOIN users u ON d.user_id = u.id WHERE d.id = $1;


-- name: GetAllPendingWithdrawalByID :many
SELECT * FROM user_withdrawals WHERE "status" = 'Pending' AND user_id = $1;

-- name: GetAllWithdrawalByUserId :many
SELECT * FROM user_withdrawals WHERE user_id = $1;

-- name: SumAllWithdrawalAmountByUserID :one
SELECT COALESCE(SUM(amount), 0)::NUMERIC(18,6) AS total_withdrawal FROM user_withdrawals WHERE "status" = 'Approved' AND user_id = $1;

-- name: SumApprovedWithdrawals :one
SELECT COALESCE(SUM(amount), 0)::NUMERIC(18,6) AS total_withdrawal FROM user_withdrawals WHERE "status" = 'Approved';

