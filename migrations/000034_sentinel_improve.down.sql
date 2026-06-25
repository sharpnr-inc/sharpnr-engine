-- ==========================================
-- DOWN MIGRATION
-- ==========================================

BEGIN;

-- Drop new tables
DROP TABLE IF EXISTS sentinel_engine.user_policies CASCADE;
DROP TABLE IF EXISTS sentinel_engine.role_policies CASCADE;

-- Restore removed columns
ALTER TABLE sentinel_engine.sentinel_policies
    ADD COLUMN source_role_id UUID NULL,
    ADD COLUMN source_managed_policy_id UUID NULL;

-- Recreate old role_policies table
CREATE TABLE sentinel_engine.role_policies (
    id UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    role_id UUID NOT NULL,
    managed_policy_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);

-- Recreate old managed_policies table
CREATE TABLE sentinel_engine.managed_policies (
    id UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);

COMMIT;