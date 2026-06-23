-- ============================================================
-- "Full Access" system managed policy + all-actions support
--
-- Until now the owner's grant was a hand-authored row written straight into
-- sentinel_policies at provisioning, so it never appeared in the managed-policy
-- catalog and the owner role showed no privileges in the UI. This migration
-- makes the owner grant a first-class managed policy:
--
--   * all_actions — when true, the generator emits a wildcard `action` (every
--     action, including ones added later) instead of an explicit action list.
--   * a single SYSTEM "Full Access" managed policy (organization_id NULL, shared
--     by every org) with a well-known id, attached to each org's owner role at
--     provisioning. The id is a fixed application-chosen UUID (not DB-generated).
-- ============================================================

ALTER TABLE sentinel_engine.managed_policies
    ADD COLUMN all_actions BOOLEAN NOT NULL DEFAULT false;

INSERT INTO sentinel_engine.managed_policies
    (id, organization_id, name, description, actions, effect, all_actions, system)
VALUES (
    'f0000000-0000-4000-8000-000000000001',
    NULL,
    'Full Access',
    'Grants every action within the organization. Attached to the owner role.',
    '{}',
    'permit',
    true,
    true
)
ON CONFLICT (id) DO NOTHING;
