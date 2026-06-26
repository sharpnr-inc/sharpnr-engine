-- ============================================================
-- Sentinel policy attachments: role_policies + user_policies
--
-- Join tables linking sentinel_engine.sentinel_policies to roles and to
-- employees. They are organizational / UI bookkeeping: the engine compiles every
-- enabled org policy regardless of attachment, so these record which policies an
-- admin grouped under a role (or granted directly to a user).
-- ============================================================

BEGIN;

-- ------------------------------------------
-- role_policies: a policy grouped under a role.
-- ------------------------------------------
CREATE TABLE sentinel_engine.role_policies (
    id                 UUID PRIMARY KEY,
    organization_id    UUID NOT NULL,
    role_id            UUID NOT NULL,
    sentinel_policy_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,

    CONSTRAINT fk_role_policies_policy
        FOREIGN KEY (sentinel_policy_id)
        REFERENCES sentinel_engine.sentinel_policies(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_role_policies_org_id     ON sentinel_engine.role_policies(organization_id);
CREATE INDEX idx_role_policies_role_id    ON sentinel_engine.role_policies(role_id);
CREATE INDEX idx_role_policies_policy_id  ON sentinel_engine.role_policies(sentinel_policy_id);
CREATE INDEX idx_role_policies_deleted_at ON sentinel_engine.role_policies(deleted_at);

-- One live attachment per (org, role, policy); a soft-deleted row must not block
-- re-attaching the same pair, hence the partial predicate.
CREATE UNIQUE INDEX uq_role_policy_assignment
    ON sentinel_engine.role_policies(organization_id, role_id, sentinel_policy_id)
    WHERE deleted_at IS NULL;

-- ------------------------------------------
-- user_policies: a policy granted directly to an employee.
-- ------------------------------------------
CREATE TABLE sentinel_engine.user_policies (
    id                 UUID PRIMARY KEY,
    organization_id    UUID NOT NULL,
    employee_id        UUID NOT NULL,
    sentinel_policy_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,

    CONSTRAINT fk_user_policies_policy
        FOREIGN KEY (sentinel_policy_id)
        REFERENCES sentinel_engine.sentinel_policies(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_user_policies_org_id      ON sentinel_engine.user_policies(organization_id);
CREATE INDEX idx_user_policies_employee_id ON sentinel_engine.user_policies(employee_id);
CREATE INDEX idx_user_policies_policy_id   ON sentinel_engine.user_policies(sentinel_policy_id);
CREATE INDEX idx_user_policies_deleted_at  ON sentinel_engine.user_policies(deleted_at);

CREATE UNIQUE INDEX uq_user_policy_assignment
    ON sentinel_engine.user_policies(organization_id, employee_id, sentinel_policy_id)
    WHERE deleted_at IS NULL;

COMMIT;
