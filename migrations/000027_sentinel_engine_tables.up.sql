-- ============================================================
-- Sentinel authorization engine (Cedar-based)
--
-- These tables back sentinel_engine/models and replace the older
-- policy_engine.* CRUD-grid RBAC design (migration 000026). The engine
-- compiles sentinel_policies into per-org Cedar PolicySets and rebuilds
-- principal membership (roles + groups) from these tables on each check.
-- ============================================================

CREATE SCHEMA IF NOT EXISTS sentinel_engine;

-- --------------------------------------
-- sentinel_engine.roles
-- A named entity that Cedar policies reference as Role::"<slug>". Slug is the
-- Cedar entity id and must be stable, since policies embed it as text.
-- --------------------------------------
CREATE TABLE sentinel_engine.roles (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    slug            TEXT NOT NULL,
    name            TEXT NOT NULL,
    description     TEXT,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ
);

-- (org, slug) is unique among live rows; a soft-deleted row must not block
-- re-creating the same slug, hence the partial predicate.
CREATE UNIQUE INDEX idx_role_org_slug
    ON sentinel_engine.roles(organization_id, slug)
    WHERE deleted_at IS NULL;
CREATE INDEX idx_role_deleted_at ON sentinel_engine.roles(deleted_at);

-- --------------------------------------
-- sentinel_engine.groups
-- A named entity users belong to; groups nest via parent_group_id (NULL for a
-- top-level group). Referenced by policies as Group::"<slug>".
-- --------------------------------------
CREATE TABLE sentinel_engine.groups (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    slug            TEXT NOT NULL,
    name            TEXT NOT NULL,
    parent_group_id UUID,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ
);

CREATE UNIQUE INDEX idx_group_org_slug
    ON sentinel_engine.groups(organization_id, slug)
    WHERE deleted_at IS NULL;
CREATE INDEX idx_group_parent_group_id ON sentinel_engine.groups(parent_group_id);
CREATE INDEX idx_group_deleted_at ON sentinel_engine.groups(deleted_at);

-- --------------------------------------
-- sentinel_engine.sentinel_policies
-- One Cedar policy (human-readable text form) for one organization. id doubles
-- as the cedar.PolicyID when the engine builds the PolicySet, so diagnostics
-- point back at this row.
-- --------------------------------------
CREATE TABLE sentinel_engine.sentinel_policies (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    name            TEXT,
    description     TEXT,
    policy          TEXT NOT NULL,
    enabled         BOOLEAN NOT NULL DEFAULT true,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX idx_sentinel_policy_organization_id ON sentinel_engine.sentinel_policies(organization_id);
CREATE INDEX idx_sentinel_policy_deleted_at ON sentinel_engine.sentinel_policies(deleted_at);

-- --------------------------------------
-- sentinel_engine.user_roles
-- Membership edge: an employee is assigned a role. Hard-deleted on
-- un-assignment (no deleted_at) so the same (user, role) pair can be
-- re-assigned later; write a separate audit log if you need history.
-- --------------------------------------
CREATE TABLE sentinel_engine.user_roles (
    organization_id UUID NOT NULL,
    user_id         UUID NOT NULL,
    role_id         UUID NOT NULL,
    assigned_by     UUID,
    created_at      TIMESTAMPTZ DEFAULT now(),
    PRIMARY KEY (user_id, role_id)
);

CREATE INDEX idx_user_roles_organization_id ON sentinel_engine.user_roles(organization_id);
CREATE INDEX idx_user_roles_role_id ON sentinel_engine.user_roles(role_id);

-- --------------------------------------
-- sentinel_engine.user_groups
-- Membership edge: an employee belongs to a group. Hard-deleted on removal,
-- same rationale as user_roles.
-- --------------------------------------
CREATE TABLE sentinel_engine.user_groups (
    organization_id UUID NOT NULL,
    user_id         UUID NOT NULL,
    group_id        UUID NOT NULL,
    assigned_by     UUID,
    created_at      TIMESTAMPTZ DEFAULT now(),
    PRIMARY KEY (user_id, group_id)
);

CREATE INDEX idx_user_groups_organization_id ON sentinel_engine.user_groups(organization_id);
CREATE INDEX idx_user_groups_group_id ON sentinel_engine.user_groups(group_id);
