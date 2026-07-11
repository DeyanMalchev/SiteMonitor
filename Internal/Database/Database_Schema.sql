-- Table 1: The websites/endpoints we want to monitor
CREATE TABLE targets (
                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         url TEXT NOT NULL,
                         environment TEXT NOT NULL DEFAULT 'production', -- e.g., production, staging
                         interval_seconds INTEGER NOT NULL DEFAULT 60,   -- How often to ping (e.g., every 30s, 60s)
                         timeout_seconds INTEGER NOT NULL DEFAULT 5,     -- Maximum time to wait for a response
                         is_active BOOLEAN NOT NULL DEFAULT TRUE,        -- Switch to pause/resume monitoring
                         created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                         updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Table 2: The historical results of every single ping
CREATE TABLE uptime_logs (
                             id BIGSERIAL PRIMARY KEY,                       -- Uses BIGSERIAL because logs grow rapidly
                             target_id UUID NOT NULL REFERENCES targets(id) ON DELETE CASCADE,
                             status_code INTEGER,                            -- e.g., 200, 404, 500 (NULL if network timeout)
                             latency_ms INTEGER NOT NULL,                     -- How fast the site responded in milliseconds
                             is_healthy BOOLEAN NOT NULL,                     -- TRUE if status is 2xx/3xx, FALSE if 4xx/5xx or timeout
                             error_message TEXT,                             -- Stores "timeout", "connection refused", etc., if failed
                             checked_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Crucial Performance Optimization: Indexes
-- Because uptime_logs will have millions of rows, we need indexes to keep queries fast!
CREATE INDEX idx_logs_target_id ON uptime_logs(target_id);
CREATE INDEX idx_logs_checked_at ON uptime_logs(checked_at DESC);