
-- name: CreatePlan :exec
INSERT INTO plans("name", min_amount, max_amount, duration_days, roi) VALUES($1, $2, $3, $4, $5);

-- name: GetPlanByID :one
SELECT * FROM plans WHERE id = $1;

-- name: DeletePlanById :exec
DELETE FROM plans WHERE id = $1;

-- name: FetchAllPlans :many
SELECT * FROM plans;