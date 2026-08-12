-- --------------------------------------
-- employee.employee_cost_centers
-- The employee's posting: the one cost center they work out of.
--
-- A transfer does not overwrite the row. It closes the current one by setting
-- effective_to and opens a new one at the destination, so a payroll run or cost
-- report for a past period still resolves the office the employee was actually
-- posted to at the time. Overwriting would silently rewrite history — which is
-- also why "one cost center per employee" is a partial unique index over live,
-- open rows rather than a plain unique on employee_id.
-- --------------------------------------
CREATE TABLE employee.employee_cost_centers (
    id                 UUID PRIMARY KEY,
    employee_id        UUID NOT NULL,
    organization_id    UUID NOT NULL,
    org_cost_center_id UUID NOT NULL,

    effective_from     DATE NOT NULL,
    effective_to       DATE,

    created_at         TIMESTAMPTZ DEFAULT now(),
    updated_at         TIMESTAMPTZ DEFAULT now(),
    deleted_at         TIMESTAMPTZ
);

-- "where is this employee posted" / "what is their posting history".
CREATE INDEX idx_employee_cost_centers_employee_id ON employee.employee_cost_centers(employee_id);

-- "who works out of this cost center" — the headcount read.
CREATE INDEX idx_employee_cost_centers_org_cost_center_id ON employee.employee_cost_centers(org_cost_center_id);

CREATE INDEX idx_employee_cost_centers_organization_id ON employee.employee_cost_centers(organization_id);
CREATE INDEX idx_employee_cost_centers_deleted_at      ON employee.employee_cost_centers(deleted_at);

-- Period lookups: "who was posted here during Q2". A closed row has a bounded
-- effective_to, an open one is NULL.
CREATE INDEX idx_employee_cost_centers_effective_to ON employee.employee_cost_centers(effective_to);

-- An employee has at most one OPEN posting. Closed rows (effective_to set) are
-- the transfer history and must not collide; soft-deleted rows must not block
-- re-posting either, hence both predicates.
CREATE UNIQUE INDEX uq_employee_cost_centers_current_posting
    ON employee.employee_cost_centers(employee_id)
    WHERE effective_to IS NULL AND deleted_at IS NULL;

-- --------------------------------------
-- employee.employee_cost_center_departments
-- Assigns an employee to a department within the cost center they are posted
-- to. Unlike the posting itself this is not one-per-employee: someone can sit
-- across two teams at the same site, so several live rows are allowed and
-- exactly one carries is_primary.
--
-- The cost center is implied by the department (core.cost_center_departments
-- holds org_cost_center_id) and deliberately not repeated here — storing it
-- twice would let the two disagree after a transfer.
-- --------------------------------------
CREATE TABLE employee.employee_cost_center_departments (
    id                        UUID PRIMARY KEY,
    employee_id               UUID NOT NULL,
    organization_id           UUID NOT NULL,
    cost_center_department_id UUID NOT NULL,

    is_primary                BOOLEAN NOT NULL DEFAULT true,

    effective_from            DATE NOT NULL,
    effective_to              DATE,

    created_at                TIMESTAMPTZ DEFAULT now(),
    updated_at                TIMESTAMPTZ DEFAULT now(),
    deleted_at                TIMESTAMPTZ
);

-- "which departments is this employee in".
CREATE INDEX idx_employee_cc_departments_employee_id ON employee.employee_cost_center_departments(employee_id);

-- "who is in this department" — the roster read, and the guard that blocks
-- deleting a department that still has people in it.
CREATE INDEX idx_employee_cc_departments_department_id ON employee.employee_cost_center_departments(cost_center_department_id);

CREATE INDEX idx_employee_cc_departments_organization_id ON employee.employee_cost_center_departments(organization_id);
CREATE INDEX idx_employee_cc_departments_deleted_at      ON employee.employee_cost_center_departments(deleted_at);
CREATE INDEX idx_employee_cc_departments_effective_to    ON employee.employee_cost_center_departments(effective_to);

-- One live assignment per (employee, department) — re-assigning someone to a
-- department they are already in is a no-op, not a second row.
CREATE UNIQUE INDEX uq_employee_cc_departments_live_assignment
    ON employee.employee_cost_center_departments(employee_id, cost_center_department_id)
    WHERE effective_to IS NULL AND deleted_at IS NULL;

-- Exactly one primary department per employee among live, open rows. This is
-- what payroll and cost allocation read to decide which department owns the
-- employee's cost.
CREATE UNIQUE INDEX uq_employee_cc_departments_primary
    ON employee.employee_cost_center_departments(employee_id)
    WHERE is_primary AND effective_to IS NULL AND deleted_at IS NULL;
