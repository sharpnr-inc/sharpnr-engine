-- ============================================================
-- DOWN: restore sentinel_policies.scope (000035's column, index and CHECK) and
-- reconstruct its values from the attachment tables.
--
-- The reconstruction mirrors 000035's own backfill, extended with the org
-- attachments 000051 introduced. Precedence: organization > role > group > user,
-- so a policy attached several ways lands in its broadest bucket — matching how
-- the pre-000052 engine read it (`scope = 'organization'` won regardless of
-- other attachments).
-- ============================================================

BEGIN;

-- Re-add nullable first so the reconstruction can run before NOT NULL applies.
ALTER TABLE sentinel_engine.sentinel_policies
    ADD COLUMN IF NOT EXISTS scope TEXT;

-- Unattached policies: deliberately NOT 'organization'.
--
-- Under the pre-000052 query, scope='organization' grants a policy to every
-- employee in the org. A policy authored after 000052 and attached to nothing
-- is granted to nobody, so parking it in the org bucket on rollback would
-- silently widen access. 'user' is arbitrary but inert — it is a pure UI label
-- for a policy that no attachment reaches.
UPDATE sentinel_engine.sentinel_policies SET scope = 'user' WHERE scope IS NULL;

UPDATE sentinel_engine.sentinel_policies p SET scope = 'user'
 WHERE EXISTS (SELECT 1 FROM sentinel_engine.user_policies up
                WHERE up.sentinel_policy_id = p.id AND up.deleted_at IS NULL);

UPDATE sentinel_engine.sentinel_policies p SET scope = 'group'
 WHERE EXISTS (SELECT 1 FROM sentinel_engine.groups_sentinel_policies gp
                WHERE gp.sentinel_policy_id = p.id AND gp.deleted_at IS NULL);

UPDATE sentinel_engine.sentinel_policies p SET scope = 'role'
 WHERE EXISTS (SELECT 1 FROM sentinel_engine.role_policies rp
                WHERE rp.sentinel_policy_id = p.id AND rp.deleted_at IS NULL);

UPDATE sentinel_engine.sentinel_policies p SET scope = 'organization'
 WHERE EXISTS (SELECT 1 FROM sentinel_engine.org_sentinel_policies op
                WHERE op.sentinel_policy_id = p.id AND op.deleted_at IS NULL);

ALTER TABLE sentinel_engine.sentinel_policies
    ALTER COLUMN scope SET DEFAULT 'organization';

ALTER TABLE sentinel_engine.sentinel_policies
    ALTER COLUMN scope SET NOT NULL;

CREATE INDEX IF NOT EXISTS idx_sentinel_policy_scope
    ON sentinel_engine.sentinel_policies(scope);

ALTER TABLE sentinel_engine.sentinel_policies
    ADD CONSTRAINT chk_sentinel_policy_scope
    CHECK (scope IN ('organization', 'role', 'group', 'user'));

COMMIT;
