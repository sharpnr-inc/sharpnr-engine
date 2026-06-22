-- --------------------------------------
-- Restore the `role` column on employee.employee.
-- --------------------------------------
ALTER TABLE employee.employee ADD COLUMN IF NOT EXISTS role VARCHAR(50) NOT NULL DEFAULT '';
