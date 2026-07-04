-- ============================================================
-- Fix legacy seeded owner policies for attachment-driven evaluation.
--
-- The engine no longer builds Role/Group Cedar entities: a policy applies to a
-- user because it is org-scoped or ATTACHED to them (directly / via roles /
-- via groups), and policy text targets the generic principal. Owner policies
-- seeded before this rework read
--
--     permit (principal in Role::"owner", action, resource)
--     when { resource.orgId == principal.orgId };
--
-- which can never match now: no Role parents exist and neither entity carries
-- an orgId attribute, so owners were denied everything (403s despite the role).
--
-- This migration rewrites those rows to the current seed form and attaches
-- them to their org's Owner role, mirroring what ProvisionOrganization creates
-- today: plain full-access text, scope = 'role', one role_policies edge.
-- ============================================================

BEGIN;

-- 1) Rewrite the Cedar text of legacy owner full-access policies and move them
--    to the role bucket. Matched by the legacy text pattern so hand-authored
--    policies are left alone.
UPDATE sentinel_engine.sentinel_policies
SET policy     = E'permit (\n  principal,\n  action,\n  resource\n);\n',
    scope      = 'role',
    updated_at = NOW()
WHERE deleted_at IS NULL
  AND policy ILIKE '%principal in role::"owner"%';

-- 2) Attach each org's owner full-access policy to that org's Owner role where
--    the edge is missing (new provisioning creates it; legacy seeds did not).
--    The edge id is derived deterministically (md5 of the key columns) instead
--    of gen_random_uuid(): ids in this codebase are app-generated, and a
--    deterministic id also keeps the insert idempotent on re-run.
INSERT INTO sentinel_engine.role_policies (id, organization_id, role_id, sentinel_policy_id, created_at)
SELECT md5('role-policy:' || r.id::text || ':' || p.id::text)::uuid,
       r.organization_id,
       r.id,
       p.id,
       NOW()
FROM sentinel_engine.sentinel_policies p
JOIN sentinel_engine.roles r
  ON  r.organization_id = p.organization_id
  AND r.name = 'Owner'
  AND r.deleted_at IS NULL
WHERE p.deleted_at IS NULL
  AND p.scope = 'role'
  AND p.name IN ('owner-full-access', 'Owner Full Access')
  AND NOT EXISTS (
      SELECT 1 FROM sentinel_engine.role_policies rp
      WHERE rp.role_id = r.id
        AND rp.sentinel_policy_id = p.id
        AND rp.deleted_at IS NULL
  );

COMMIT;
