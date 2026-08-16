-- --------------------------------------
-- employee.employee_provisioning
-- One row per invite link emailed to a newly created employee.
-- --------------------------------------
CREATE TABLE employee.employee_provisioning (
    id          UUID PRIMARY KEY,
    employee_id UUID NOT NULL,
    email       VARCHAR(150) NOT NULL,

    token_hash  TEXT NOT NULL,

    expires_at  TIMESTAMPTZ NOT NULL,
    consumed_at TIMESTAMPTZ,

    created_at  TIMESTAMPTZ DEFAULT now()
);

-- Token lookup on the "set your password" callback
CREATE UNIQUE INDEX idx_employee_provisioning_token_hash
    ON employee.employee_provisioning(token_hash);

-- Re-invite / "does this employee already have a pending invite" lookups
CREATE INDEX idx_employee_provisioning_employee_id
    ON employee.employee_provisioning(employee_id);
