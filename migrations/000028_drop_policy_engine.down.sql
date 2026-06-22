-- Recreate the legacy policy_engine schema and tables (mirrors 000001 + 000026).

CREATE SCHEMA IF NOT EXISTS policy_engine;

-- --------------------------------------
-- policy_engine.role
-- A reusable, named bundle of access scoped to one organization.
-- --------------------------------------
CREATE TABLE policy_engine.role (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    name            VARCHAR(100) NOT NULL,
    description     VARCHAR(255),
    is_system       BOOLEAN NOT NULL DEFAULT false,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX idx_role_organization_id ON policy_engine.role(organization_id);
CREATE INDEX idx_role_deleted_at ON policy_engine.role(deleted_at);

-- --------------------------------------
-- policy_engine.permission
-- One row of the access grid: for a role, on a resource type, optionally
-- scoped to a property, the allowed CRUD actions. property_id NULL = all
-- properties; resource_type '*' = all resource types.
-- --------------------------------------
CREATE TABLE policy_engine.permission (
    id            UUID PRIMARY KEY,
    role_id       UUID NOT NULL,
    property_id   UUID,
    resource_type VARCHAR(50) NOT NULL,
    can_view      BOOLEAN NOT NULL DEFAULT false,
    can_create    BOOLEAN NOT NULL DEFAULT false,
    can_edit      BOOLEAN NOT NULL DEFAULT false,
    can_delete    BOOLEAN NOT NULL DEFAULT false,
    created_at    TIMESTAMPTZ DEFAULT now(),
    updated_at    TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_permission_role_id ON policy_engine.permission(role_id);
CREATE INDEX idx_permission_property_id ON policy_engine.permission(property_id);

-- --------------------------------------
-- policy_engine.employee_role
-- Assigns a role to an employee (many-to-many).
-- --------------------------------------
CREATE TABLE policy_engine.employee_role (
    id              UUID PRIMARY KEY,
    employee_id     UUID NOT NULL,
    role_id         UUID NOT NULL,
    organization_id UUID NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX uq_employee_role ON policy_engine.employee_role(employee_id, role_id);
CREATE INDEX idx_employee_role_employee_id ON policy_engine.employee_role(employee_id);
CREATE INDEX idx_employee_role_role_id ON policy_engine.employee_role(role_id);
CREATE INDEX idx_employee_role_organization_id ON policy_engine.employee_role(organization_id);
