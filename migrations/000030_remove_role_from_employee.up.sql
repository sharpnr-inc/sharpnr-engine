-- --------------------------------------
-- Drop the legacy `role` column from employee.employee.
-- Roles are now managed via the Sentinel RBAC engine
-- (policy_engine.employee_role / sentinel tables).
-- --------------------------------------
ALTER TABLE employee.employee DROP COLUMN IF EXISTS role;
