BEGIN;

ALTER TABLE sentinel_engine.groups
ADD COLUMN slug TEXT;

ALTER TABLE sentinel_engine.groups
ALTER COLUMN slug SET NOT NULL;

CREATE UNIQUE INDEX idx_group_org_slug
ON sentinel_engine.groups (organization_id, slug);

COMMIT;