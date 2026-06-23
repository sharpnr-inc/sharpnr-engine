-- ============================================================
-- Sentinel managed policies (AWS-style)
--
-- Adds the authoring/display layer on top of the Cedar engine:
--   * managed_policies — reusable, named permission bundles (like AWS managed
--     policies). org_id NULL = a system policy shipped to every org.
--   * role_policies    — the attachment: role <-> managed policy (many-to-many).
--
-- These two tables are the SOURCE OF TRUTH the UI reads and edits. The engine
-- still reads sentinel_engine.sentinel_policies, whose rows are GENERATED from
-- the attachments above (cedar-go 1.8.0 has no policy templates, so we compile
-- a concrete `principal in Role::"<slug>"` policy per attachment). The two new
-- columns on sentinel_policies tag each generated row with its source so
-- regeneration is idempotent (delete-by-source, re-insert).
-- ============================================================

-- --------------------------------------
-- sentinel_engine.managed_policies
-- A named bundle of actions. `actions` holds Cedar action ids (the catalog Keys
-- in sentinel_engine/actions). `effect` is permit|forbid. resource_type is an
-- optional scope (NULL = any resource in the org).
-- --------------------------------------
CREATE TABLE sentinel_engine.managed_policies (
    id              UUID PRIMARY KEY,
    organization_id UUID,                                  -- NULL = system policy (all orgs)
    name            TEXT NOT NULL,
    description     TEXT,
    actions         TEXT[] NOT NULL DEFAULT '{}',
    effect          TEXT NOT NULL DEFAULT 'permit',        -- 'permit' | 'forbid'
    resource_type   TEXT,                                  -- NULL = any resource
    system          BOOLEAN NOT NULL DEFAULT false,        -- platform-provided; FE cannot edit/delete
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ
);

-- (org, name) unique among live rows; system rows (org_id NULL) are unique by name.
CREATE UNIQUE INDEX idx_managed_policy_org_name
    ON sentinel_engine.managed_policies(organization_id, name)
    WHERE deleted_at IS NULL;
CREATE INDEX idx_managed_policy_organization_id ON sentinel_engine.managed_policies(organization_id);
CREATE INDEX idx_managed_policy_deleted_at ON sentinel_engine.managed_policies(deleted_at);

-- --------------------------------------
-- sentinel_engine.role_policies
-- Attachment edge: a role has a managed policy attached. Hard-deleted on detach
-- (no deleted_at) so the same (role, policy) pair can be re-attached later,
-- mirroring user_roles/user_groups.
-- --------------------------------------
CREATE TABLE sentinel_engine.role_policies (
    organization_id   UUID NOT NULL,
    role_id           UUID NOT NULL,
    managed_policy_id UUID NOT NULL,
    attached_by       UUID,                                -- employee who attached (audit; optional)
    created_at        TIMESTAMPTZ DEFAULT now(),
    PRIMARY KEY (role_id, managed_policy_id)
);

CREATE INDEX idx_role_policies_organization_id ON sentinel_engine.role_policies(organization_id);
CREATE INDEX idx_role_policies_managed_policy_id ON sentinel_engine.role_policies(managed_policy_id);

-- --------------------------------------
-- Tag generated sentinel_policies with their source attachment so the generator
-- can delete-and-replace exactly the rows it owns. NULL on both columns marks a
-- hand-authored / non-generated policy (e.g. the legacy owner policy), which the
-- generator must leave untouched.
-- --------------------------------------
ALTER TABLE sentinel_engine.sentinel_policies
    ADD COLUMN source_role_id           UUID,
    ADD COLUMN source_managed_policy_id UUID;

CREATE INDEX idx_sentinel_policy_source_role_id
    ON sentinel_engine.sentinel_policies(source_role_id);
