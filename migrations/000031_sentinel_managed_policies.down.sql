ALTER TABLE sentinel_engine.sentinel_policies
    DROP COLUMN IF EXISTS source_role_id,
    DROP COLUMN IF EXISTS source_managed_policy_id;

DROP TABLE IF EXISTS sentinel_engine.role_policies;
DROP TABLE IF EXISTS sentinel_engine.managed_policies;
