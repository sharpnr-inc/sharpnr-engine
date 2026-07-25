-- --------------------------------------
-- Restore the one-account-per-employee constraint on
-- employee.employee_bank_account.
-- --------------------------------------
DROP INDEX IF EXISTS employee.idx_employee_bank_account_primary;
DROP INDEX IF EXISTS employee.idx_employee_bank_account_employee_id;
CREATE UNIQUE INDEX idx_employee_bank_account_employee_id ON employee.employee_bank_account(employee_id);
ALTER TABLE employee.employee_bank_account ALTER COLUMN is_primary SET DEFAULT true;

-- --------------------------------------
-- Restore employee.employee_organization_profile.
-- --------------------------------------
CREATE TABLE employee.employee_organization_profile (
    id              UUID PRIMARY KEY,
    employee_id     UUID NOT NULL,
    organization_id UUID NOT NULL,
    department      VARCHAR(100),
    position        VARCHAR(100),
    manager_id      UUID,
    created_at      BIGINT,
    updated_at      BIGINT
);

CREATE UNIQUE INDEX idx_emp_org_profile_employee_id ON employee.employee_organization_profile(employee_id);
CREATE INDEX idx_emp_org_profile_organization_id ON employee.employee_organization_profile(organization_id);
CREATE INDEX idx_emp_org_profile_manager_id ON employee.employee_organization_profile(manager_id);
