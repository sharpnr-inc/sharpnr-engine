-- --------------------------------------
-- employee.employee_bank_account:
-- allow multiple bank accounts per employee. The unique index on
-- employee_id contradicted the is_primary flag (it forced exactly
-- one account, making is_primary meaningless).
-- --------------------------------------
DROP INDEX IF EXISTS employee.idx_employee_bank_account_employee_id;
CREATE INDEX idx_employee_bank_account_employee_id ON employee.employee_bank_account(employee_id);

-- With multiple accounts allowed, new rows must not default to primary,
-- and at most one account per employee may be primary.
ALTER TABLE employee.employee_bank_account ALTER COLUMN is_primary SET DEFAULT false;
CREATE UNIQUE INDEX idx_employee_bank_account_primary
    ON employee.employee_bank_account(employee_id)
    WHERE is_primary;

-- --------------------------------------
-- Drop employee.employee_organization_profile.
-- It duplicated employee_employment (designation) and
-- employee_reporting (manager) and was unused.
-- --------------------------------------
DROP TABLE IF EXISTS employee.employee_organization_profile;
