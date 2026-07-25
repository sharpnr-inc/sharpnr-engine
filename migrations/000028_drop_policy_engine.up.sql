-- ============================================================
-- Drop the legacy policy_engine schema.
--
-- policy_engine held an older CRUD-grid RBAC design (role / permission /
-- employee_role, created in 000001 + 000026). It has been superseded by the
-- Cedar-based sentinel_engine schema (000027). No application code reads or
-- writes policy_engine anymore, so we remove it entirely.
-- ============================================================

DROP TABLE IF EXISTS policy_engine.employee_role;
DROP TABLE IF EXISTS policy_engine.permission;
DROP TABLE IF EXISTS policy_engine.role;

DROP SCHEMA IF EXISTS policy_engine;
