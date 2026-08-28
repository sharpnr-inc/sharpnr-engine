-- ============================================================
-- Drop sentinel_policies.scope.
--
-- `scope` (added in 000035) was doing two jobs: a UI grouping label, and the
-- enforcement flag for org-wide applicability — `scope = 'organization'` was
-- the one value the engine's union query read. 000051 replaced that second job
-- with org_sentinel_policies, so applicability is now uniformly
-- attachment-driven and the column has no behaviour left to carry.
--
-- ⚠ DEPLOY ORDER: the application must STOP READING AND WRITING `scope` before
-- this migration runs. At minimum:
--   * sentinel/repo/engine_store.go   — GetAllPoliciesOfUser must read
--                                       org_sentinel_policies, not p.scope
--   * sentinel/repo/sentinel_policy.go — ListPoliciesByOrganizationID's scope
--                                       filter, and GetPolciesScopedToOrganization
--   * sentinel/models/policy.go       — the SentinelPolicy.Scope GORM field
--                                       (Create/Save emit the column)
-- Running this against a build that still references the column breaks
-- authorization, policy writes, and org provisioning.
-- ============================================================

BEGIN;

-- ------------------------------------------
-- 1) Final safety backfill.
--
-- 000051 backfilled the attachments that existed at that moment, but any policy
-- created through CreatePolicy since then wrote `scope` and no attachment row.
-- Re-run the same idempotent insert so nothing org-wide is lost when the column
-- disappears. Identical to 000051's backfill (deterministic md5 id, NOT EXISTS
-- guard), so it is a no-op when there has been no drift.
-- ------------------------------------------
INSERT INTO sentinel_engine.org_sentinel_policies (
    id, organization_id, sentinel_policy_id, attached_by, created_at
)
SELECT md5('org-policy:' || p.organization_id::text || ':' || p.id::text)::uuid,
       p.organization_id,
       p.id,
       NULL,
       COALESCE(p.created_at, NOW())
FROM sentinel_engine.sentinel_policies p
WHERE p.deleted_at IS NULL
  AND p.scope = 'organization'
  AND NOT EXISTS (
      SELECT 1 FROM sentinel_engine.org_sentinel_policies op
      WHERE op.organization_id = p.organization_id
        AND op.sentinel_policy_id = p.id
        AND op.deleted_at IS NULL
  );

-- ------------------------------------------
-- 2) Drop the column and everything 000035 hung off it.
--    DROP COLUMN would cascade to both, but naming them keeps the reversal in
--    000052.down.sql an exact mirror.
-- ------------------------------------------
ALTER TABLE sentinel_engine.sentinel_policies
    DROP CONSTRAINT IF EXISTS chk_sentinel_policy_scope;

DROP INDEX IF EXISTS sentinel_engine.idx_sentinel_policy_scope;

ALTER TABLE sentinel_engine.sentinel_policies
    DROP COLUMN IF EXISTS scope;

COMMIT;
