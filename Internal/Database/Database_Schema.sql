CREATE TABLE Targets (
                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         url TEXT NOT NULL,
                         name TEXT,                                      -- NEW: Optional friendly name (e.g. 'YouTube')
                         environment TEXT NOT NULL DEFAULT 'production',
                         interval_seconds INTEGER NOT NULL DEFAULT 60,
                         timeout_seconds INTEGER NOT NULL DEFAULT 5,
                         is_active BOOLEAN NOT NULL DEFAULT TRUE,
                         created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                         updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE Uptime_logs (
                             id BIGSERIAL PRIMARY KEY,
                             target_id UUID NOT NULL REFERENCES targets(id) ON DELETE CASCADE,
                             status_code INTEGER,
                             latency_ms INTEGER NOT NULL,
                             is_healthy BOOLEAN NOT NULL,
                             error_message TEXT,
                             checked_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_logs_target_id ON Uptime_logs(target_id);
CREATE INDEX idx_logs_checked_at ON Uptime_logs(checked_at DESC);