-- name: CreateTarget :one
INSERT INTO targets (url, name, environment, interval_seconds, timeout_seconds)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListActiveTargets :many
SELECT * FROM targets
WHERE is_active = TRUE;

-- name: InsertUptimeLog :one
INSERT INTO uptime_logs (target_id, status_code, latency_ms, is_healthy, error_message)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTargetStatsByURLorName :many
SELECT
    t.name as target_name,
    t.url as target_url,
    l.status_code,
    l.is_healthy,
    l.latency_ms,
    l.checked_at
FROM uptime_logs l
         INNER JOIN targets t ON l.target_id = t.id
WHERE t.name ILIKE $1 OR t.url ILIKE $2
ORDER BY l.checked_at DESC
LIMIT $3;