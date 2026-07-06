CREATE SCHEMA IF NOT EXISTS audit_logs;

-- --------------------------------------
-- audit_logs.logs
-- One row per tenant event. IDs are app-generated (uuid.New()).
-- --------------------------------------
CREATE TABLE audit_logs.logs (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,

    actor_id        UUID,
    actor_type      VARCHAR(50),

    module          VARCHAR(100),
    action          VARCHAR(150) NOT NULL,

    resource_type   VARCHAR(100),
    resource_id     UUID,

    description     TEXT,
    metadata        JSONB,

    created_at      TIMESTAMPTZ DEFAULT now()
);

-- Org history feed (newest first)
CREATE INDEX idx_audit_logs_org_created_at
    ON audit_logs.logs (organization_id, created_at DESC);

-- Per-entity trail lookups
CREATE INDEX idx_audit_logs_org_resource
    ON audit_logs.logs (organization_id, resource_type, resource_id);

-- "What did this employee do" lookups
CREATE INDEX idx_audit_logs_org_actor
    ON audit_logs.logs (organization_id, actor_id);
