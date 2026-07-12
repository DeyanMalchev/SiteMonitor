-- name: CreateTarget :one
INSERT INTO targets (url, environment, interval_seconds, timeout_seconds)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListActiveTargets :many
SELECT * FROM targets
WHERE is_active = TRUE;

-- name: InsertUptimeLog :one
INSERT INTO uptime_logs (target_id, status_code, latency_ms, is_healthy, error_message)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTargetStats :many
SELECT status_code, is_healthy, latency_ms, checked_at
FROM uptime_logs
WHERE target_id = $1
ORDER BY checked_at DESC
LIMIT $2;