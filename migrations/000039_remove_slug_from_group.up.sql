BEGIN;

DROP INDEX IF EXISTS sentinel_engine.idx_group_org_slug;

ALTER TABLE sentinel_engine.groups
DROP COLUMN IF EXISTS slug;

COMMIT;