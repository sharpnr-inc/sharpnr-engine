-- ============================================================
-- Drop sentinel_engine.roles.slug.
--
-- Roles are now identified by name (unique per org); the slug column was
-- redundant. Replaces the (org, slug) unique index with (org, name).
-- ============================================================

BEGIN;

DROP INDEX IF EXISTS sentinel_engine.idx_role_org_slug;

ALTER TABLE sentinel_engine.roles
    DROP COLUMN IF EXISTS slug;

-- name becomes the unique per-org identifier (the slug's old role). Partial so a
-- soft-deleted row does not block re-creating the same name.
CREATE UNIQUE INDEX idx_role_org_name
    ON sentinel_engine.roles(organization_id, name)
    WHERE deleted_at IS NULL;

COMMIT;
