-- ============================================================
-- DOWN: reverse of 000036 — drop the Cedar JSON column.
-- ============================================================

BEGIN;

ALTER TABLE sentinel_engine.sentinel_policies
    DROP COLUMN IF EXISTS policy_json;

COMMIT;
