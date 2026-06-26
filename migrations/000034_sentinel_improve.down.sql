-- ============================================================
-- DOWN: reverse of 000034 up — drop the policy attachment tables.
-- ============================================================

BEGIN;

DROP TABLE IF EXISTS sentinel_engine.user_policies CASCADE;
DROP TABLE IF EXISTS sentinel_engine.role_policies CASCADE;

COMMIT;
