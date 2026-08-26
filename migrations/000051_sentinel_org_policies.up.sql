-- ============================================================
-- Sentinel org policy attachments: org_sentinel_policies
--
-- The fourth attachment path, mirroring role_policies (000034),
-- user_policies (000034) and groups_sentinel_policies (000040): a policy
-- attached here applies to EVERY user in the organization.
--
-- This replaces `sentinel_policies.scope = 'organization'` as the org-wide
-- mechanism, so that applicability is uniformly attachment-driven — a policy
-- reaches a user iff an attachment edge reaches them, with no special case for
-- the org-wide bucket. Two consequences worth stating:
--
--   * A newly created policy is now attached to nobody, so it affects nobody
--     until it is deliberately attached. The old default (scope defaulting to
--     'organization') silently granted org-wide.
--   * Going org-wide is an INSERT and lifting it is a soft DELETE, both with an
--     attached_by/created_at/deleted_at trail. Previously it was an UPDATE to
--     the policy row itself, which left no history.
--
-- `scope` is left in place by this migration; it reverts to a pure UI
-- classification label once the engine's union query reads this table instead.
-- Dropping the column is a separate migration, after the query switches over.
-- ============================================================

BEGIN;

CREATE TABLE sentinel_engine.org_sentinel_policies (
    id                 UUID PRIMARY KEY,
    organization_id    UUID NOT NULL,
    sentinel_policy_id UUID NOT NULL,
    attached_by        UUID NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,

    CONSTRAINT fk_org_policies_policy
        FOREIGN KEY (sentinel_policy_id)
        REFERENCES sentinel_engine.sentinel_policies(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_org_policies_org_id     ON sentinel_engine.org_sentinel_policies(organization_id);
CREATE INDEX idx_org_policies_policy_id  ON sentinel_engine.org_sentinel_policies(sentinel_policy_id);
CREATE INDEX idx_org_policies_deleted_at ON sentinel_engine.org_sentinel_policies(deleted_at);

-- One live attachment per (org, policy); a soft-deleted row must not block
-- re-attaching the same pair, hence the partial predicate.
CREATE UNIQUE INDEX uq_org_policy_assignment
    ON sentinel_engine.org_sentinel_policies(organization_id, sentinel_policy_id)
    WHERE deleted_at IS NULL;

-- ------------------------------------------
-- Backfill: every policy currently org-wide by scope becomes an attachment, so
-- the new table is authoritative the moment the engine reads it. Soft-deleted
-- policies are skipped; disabled ones are kept (enabled is a separate switch
-- from attachment, and re-enabling must not silently lose the attachment).
--
-- The edge id is derived deterministically (md5 of the key columns) instead of
-- gen_random_uuid(), matching 000041: ids in this codebase are app-generated,
-- and a deterministic id also keeps the insert idempotent on re-run.
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

COMMIT;
