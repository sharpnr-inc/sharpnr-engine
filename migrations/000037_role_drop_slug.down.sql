-- ============================================================
-- DOWN: restore sentinel_engine.roles.slug.
--
-- Backfills slug from name (unique per org, so the (org, slug) unique index
-- holds) before re-applying the constraint.
-- ============================================================

BEGIN;

DROP INDEX IF EXISTS sentinel_engine.idx_role_org_name;

ALTER TABLE sentinel_engine.roles
    ADD COLUMN slug TEXT;

UPDATE sentinel_engine.roles
   SET slug = name
 WHERE slug IS NULL;

ALTER TABLE sentinel_engine.roles
    ALTER COLUMN slug SET NOT NULL;

CREATE UNIQUE INDEX idx_role_org_slug
    ON sentinel_engine.roles(organization_id, slug)
    WHERE deleted_at IS NULL;

COMMIT;
