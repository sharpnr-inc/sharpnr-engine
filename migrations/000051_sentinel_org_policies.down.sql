-- Reverting drops the org-wide attachments. This is lossless only while
-- `sentinel_policies.scope` is still the source of truth for org-wide
-- applicability — once the engine reads this table and scope is dropped, the
-- attachment rows are the only record and this down migration loses them.
BEGIN;

DROP TABLE IF EXISTS sentinel_engine.org_sentinel_policies;

COMMIT;
