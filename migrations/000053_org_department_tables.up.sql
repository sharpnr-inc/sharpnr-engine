-- --------------------------------------
-- core.org_departments
-- An ORG-LEVEL department — Finance, Housekeeping, Engineering — hanging off
-- the organization directly. This is the track used by organizations that want
-- departments without managing cost centers; an org that does manage cost
-- centers defines its departments in core.cost_center_departments instead. The
-- two tracks are independent by design and never mix.
--
-- code is the short human key ("FIN"), derived from the name when the caller
-- does not supply one. It is unique per organization among live rows.
--
-- head_employee_id is the employee who runs the department. Nullable: a
-- department is often created before anyone is put in charge of it.
-- --------------------------------------
CREATE TABLE core.org_departments (
    id               UUID PRIMARY KEY,
    organization_id  UUID NOT NULL,

    name             VARCHAR(150) NOT NULL,
    description      TEXT,
    code             VARCHAR(50) NOT NULL,

    head_employee_id UUID,

    created_at       TIMESTAMPTZ DEFAULT now(),
    updated_at       TIMESTAMPTZ DEFAULT now(),
    deleted_at       TIMESTAMPTZ
);

-- The tenant predicate every read carries.
CREATE INDEX idx_org_departments_organization_id  ON core.org_departments(organization_id);

CREATE INDEX idx_org_departments_code             ON core.org_departments(code);
CREATE INDEX idx_org_departments_head_employee_id ON core.org_departments(head_employee_id);
CREATE INDEX idx_org_departments_deleted_at       ON core.org_departments(deleted_at);

-- (org, code) and (org, name) are unique among live rows; a soft-deleted row
-- must not block re-using either, hence the partial predicate.
CREATE UNIQUE INDEX uq_org_departments_org_code
    ON core.org_departments(organization_id, code)
    WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX uq_org_departments_org_name
    ON core.org_departments(organization_id, lower(name))
    WHERE deleted_at IS NULL;

-- --------------------------------------
-- core.org_departments_employees
-- Membership: which employees belong to an org-level department. A department
-- has many employees and an employee may sit in more than one.
--
-- organization_id is denormalised from the department so every tenant-scoped
-- read of this table is a single indexed predicate rather than a join.
-- --------------------------------------
CREATE TABLE core.org_departments_employees (
    id                UUID PRIMARY KEY,
    organization_id   UUID NOT NULL,
    org_department_id UUID NOT NULL,
    employee_id       UUID NOT NULL,

    created_at        TIMESTAMPTZ DEFAULT now(),
    updated_at        TIMESTAMPTZ DEFAULT now(),
    deleted_at        TIMESTAMPTZ
);

CREATE INDEX idx_org_departments_employees_organization_id   ON core.org_departments_employees(organization_id);

-- "who is in this department" — the roster read, and the guard that blocks
-- deleting a department that still has people in it.
CREATE INDEX idx_org_departments_employees_department_id     ON core.org_departments_employees(org_department_id);

-- "which departments is this employee in".
CREATE INDEX idx_org_departments_employees_employee_id       ON core.org_departments_employees(employee_id);

CREATE INDEX idx_org_departments_employees_deleted_at        ON core.org_departments_employees(deleted_at);

-- One live membership row per (department, employee).
CREATE UNIQUE INDEX uq_org_departments_employees_department_employee
    ON core.org_departments_employees(org_department_id, employee_id)
    WHERE deleted_at IS NULL;
