BEGIN;

-- ============================================================
-- user_roles
-- ============================================================

ALTER TABLE sentinel_engine.user_roles
DROP CONSTRAINT IF EXISTS user_roles_pkey;

ALTER TABLE sentinel_engine.user_roles
ADD COLUMN id UUID;

UPDATE sentinel_engine.user_roles
SET id = gen_random_uuid()
WHERE id IS NULL;

ALTER TABLE sentinel_engine.user_roles
ALTER COLUMN id SET NOT NULL;

ALTER TABLE sentinel_engine.user_roles
ADD CONSTRAINT user_roles_pkey PRIMARY KEY (id);


-- ============================================================
-- user_groups
-- ============================================================

ALTER TABLE sentinel_engine.user_groups
DROP CONSTRAINT IF EXISTS user_groups_pkey;

ALTER TABLE sentinel_engine.user_groups
ADD COLUMN id UUID;

UPDATE sentinel_engine.user_groups
SET id = gen_random_uuid()
WHERE id IS NULL;

ALTER TABLE sentinel_engine.user_groups
ALTER COLUMN id SET NOT NULL;

ALTER TABLE sentinel_engine.user_groups
ADD CONSTRAINT user_groups_pkey PRIMARY KEY (id);

COMMIT;