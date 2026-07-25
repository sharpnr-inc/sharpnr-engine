-- ============================================================
-- DOWN: reverse of 000035 — drop the scope classification.
-- ============================================================

BEGIN;

ALTER TABLE sentinel_engine.sentinel_policies
    DROP CONSTRAINT IF EXISTS chk_sentinel_policy_scope;

DROP INDEX IF EXISTS sentinel_engine.idx_sentinel_policy_scope;

ALTER TABLE sentinel_engine.sentinel_policies
    DROP COLUMN IF EXISTS scope;

COMMIT;
