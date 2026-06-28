BEGIN;

-- ============================================================
-- user_roles
-- ============================================================

ALTER TABLE sentinel_engine.user_roles
DROP CONSTRAINT IF EXISTS user_roles_pkey;

ALTER TABLE sentinel_engine.user_roles
ADD CONSTRAINT user_roles_pkey
PRIMARY KEY (user_id, role_id);

ALTER TABLE sentinel_engine.user_roles
DROP COLUMN IF EXISTS id;


-- ============================================================
-- user_groups
-- ============================================================

ALTER TABLE sentinel_engine.user_groups
DROP CONSTRAINT IF EXISTS user_groups_pkey;

ALTER TABLE sentinel_engine.user_groups
ADD CONSTRAINT user_groups_pkey
PRIMARY KEY (user_id, group_id);

ALTER TABLE sentinel_engine.user_groups
DROP COLUMN IF EXISTS id;

COMMIT;