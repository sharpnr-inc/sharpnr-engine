-- ============================================================
-- Sentinel group policy attachments: groups_sentinel_policies
--
-- Join table linking sentinel_engine.sentinel_policies to groups, mirroring
-- role_policies (000034). Unlike the earlier design where attachments were UI
-- bookkeeping only, attachments now DRIVE evaluation: the engine's
-- GetAllPoliciesOfUser unions org-scoped policies with the policies attached to
-- the user's roles, groups, and the user directly.
-- ============================================================

BEGIN;

CREATE TABLE sentinel_engine.groups_sentinel_policies (
    id                 UUID PRIMARY KEY,
    organization_id    UUID NOT NULL,
    group_id           UUID NOT NULL,
    sentinel_policy_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,

    CONSTRAINT fk_group_policies_policy
        FOREIGN KEY (sentinel_policy_id)
        REFERENCES sentinel_engine.sentinel_policies(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_group_policies_group
        FOREIGN KEY (group_id)
        REFERENCES sentinel_engine.groups(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_group_policies_org_id     ON sentinel_engine.groups_sentinel_policies(organization_id);
CREATE INDEX idx_group_policies_group_id   ON sentinel_engine.groups_sentinel_policies(group_id);
CREATE INDEX idx_group_policies_policy_id  ON sentinel_engine.groups_sentinel_policies(sentinel_policy_id);
CREATE INDEX idx_group_policies_deleted_at ON sentinel_engine.groups_sentinel_policies(deleted_at);

-- One live attachment per (org, group, policy); a soft-deleted row must not
-- block re-attaching the same pair, hence the partial predicate.
CREATE UNIQUE INDEX uq_group_policy_assignment
    ON sentinel_engine.groups_sentinel_policies(organization_id, group_id, sentinel_policy_id)
    WHERE deleted_at IS NULL;

COMMIT;
