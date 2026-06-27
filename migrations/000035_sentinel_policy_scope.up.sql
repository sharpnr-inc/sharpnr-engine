-- ============================================================
-- Classify policies by scope.
--
-- Adds sentinel_policies.scope so the UI can group policies as org-wide, role,
-- group, or user. For now this is a classification label only — the engine still
-- compiles every enabled org policy. It will also drive the per-employee policy
-- load once attachment-driven evaluation lands.
-- ============================================================

BEGIN;

ALTER TABLE sentinel_engine.sentinel_policies
    ADD COLUMN scope TEXT NOT NULL DEFAULT 'organization';

-- Backfill existing policies from their attachments: anything linked to a role
-- is a role policy, anything linked directly to a user is a user policy. The rest
-- keep the 'organization' default.
UPDATE sentinel_engine.sentinel_policies p
   SET scope = 'role'
 WHERE EXISTS (
     SELECT 1 FROM sentinel_engine.role_policies rp
     WHERE rp.sentinel_policy_id = p.id
 );

UPDATE sentinel_engine.sentinel_policies p
   SET scope = 'user'
 WHERE EXISTS (
     SELECT 1 FROM sentinel_engine.user_policies up
     WHERE up.sentinel_policy_id = p.id
 );

CREATE INDEX idx_sentinel_policy_scope
    ON sentinel_engine.sentinel_policies(scope);

ALTER TABLE sentinel_engine.sentinel_policies
    ADD CONSTRAINT chk_sentinel_policy_scope
    CHECK (scope IN ('organization', 'role', 'group', 'user'));

COMMIT;
