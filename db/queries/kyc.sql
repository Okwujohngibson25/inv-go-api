

-- name: Savekyc :exec
INSERT INTO user_kycs(user_id, id_image_front, id_image_back) VALUES ($1, $2, $3);

-- name: GetAllKyc :many
SELECT 
    d.*,
    u.username,
    u.email,
    u.fullname
FROM user_kycs d
JOIN users u ON d.user_id = u.id ORDER BY d.created_at DESC;

-- name: GetAllPendingKyc :many
SELECT 
    d.*,
    u.username,
    u.email,
    u.fullname
FROM user_kycs d
JOIN users u ON d.user_id = u.id 
WHERE "status" = 'Pending'
ORDER BY d.created_at DESC;

-- name: GetKycById :one
SELECT * FROM user_kycs WHERE id = $1;

-- name: Approvekyc :exec
UPDATE user_kycs SET "status" = 'Approved' WHERE id = $1;

-- name: Declinekyc :exec
UPDATE user_kycs SET "status" = 'Declined' WHERE id = $1;
