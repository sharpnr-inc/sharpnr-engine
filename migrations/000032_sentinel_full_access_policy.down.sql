DELETE FROM sentinel_engine.managed_policies
    WHERE id = 'f0000000-0000-4000-8000-000000000001';

ALTER TABLE sentinel_engine.managed_policies
    DROP COLUMN IF EXISTS all_actions;
