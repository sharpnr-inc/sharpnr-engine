-- ==========================================
-- UP MIGRATION
-- ==========================================

BEGIN;

-- Drop old tables
DROP TABLE IF EXISTS sentinel_engine.managed_policies CASCADE;
DROP TABLE IF EXISTS sentinel_engine.role_policies CASCADE;

-- Remove obsolete columns
ALTER TABLE sentinel_engine.sentinel_policies
    DROP COLUMN IF EXISTS source_role_id,
    DROP COLUMN IF EXISTS source_managed_policy_id;

-- ==========================================
-- Create role_policies
-- ==========================================

CREATE TABLE sentinel_engine.role_policies (
    id UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    role_id UUID NOT NULL,
    sentinel_policy_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,

    CONSTRAINT fk_role_policies_policy
        FOREIGN KEY (sentinel_policy_id)
        REFERENCES sentinel_engine.sentinel_policies(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_role_policies_org_id
    ON sentinel_engine.role_policies(organization_id);

CREATE INDEX idx_role_policies_role_id
    ON sentinel_engine.role_policies(role_id);

CREATE INDEX idx_role_policies_policy_id
    ON sentinel_engine.role_policies(sentinel_policy_id);

CREATE INDEX idx_role_policies_deleted_at
    ON sentinel_engine.role_policies(deleted_at);

-- Prevent duplicate assignments
CREATE UNIQUE INDEX uq_role_policy_assignment
    ON sentinel_engine.role_policies(
        organization_id,
        role_id,
        sentinel_policy_id
    )
    WHERE deleted_at IS NULL;

-- ==========================================
-- Create user_policies
-- ==========================================

CREATE TABLE sentinel_engine.user_policies (
    id UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    employee_id UUID NOT NULL,
    sentinel_policy_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,

    CONSTRAINT fk_user_policies_policy
        FOREIGN KEY (sentinel_policy_id)
        REFERENCES sentinel_engine.sentinel_policies(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_user_policies_org_id
    ON sentinel_engine.user_policies(organization_id);

CREATE INDEX idx_user_policies_employee_id
    ON sentinel_engine.user_policies(employee_id);

CREATE INDEX idx_user_policies_policy_id
    ON sentinel_engine.user_policies(sentinel_policy_id);

CREATE INDEX idx_user_policies_deleted_at
    ON sentinel_engine.user_policies(deleted_at);

-- Prevent duplicate assignments
CREATE UNIQUE INDEX uq_user_policy_assignment
    ON sentinel_engine.user_policies(
        organization_id,
        employee_id,
        sentinel_policy_id
    )
    WHERE deleted_at IS NULL;

COMMIT;