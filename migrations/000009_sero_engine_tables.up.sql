CREATE TABLE sero_engine.sero_outbox (
    id         UUID PRIMARY KEY,
    event_name VARCHAR(255),
    payload    JSONB,
    status     VARCHAR(50) DEFAULT 'PENDING',
    retries    INT DEFAULT 0,
    last_error TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_sero_outbox_event_name ON sero_engine.sero_outbox(event_name);
CREATE INDEX idx_sero_outbox_status ON sero_engine.sero_outbox(status);
CREATE INDEX idx_sero_outbox_deleted_at ON sero_engine.sero_outbox(deleted_at);