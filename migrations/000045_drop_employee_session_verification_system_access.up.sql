-- EmployeeSession, EmployeeVerification and EmployeeSystemAccess are being
-- replaced by a different employee onboarding design; drop their tables.
DROP TABLE IF EXISTS employee.employee_system_access;
DROP TABLE IF EXISTS employee.employee_session;
DROP TABLE IF EXISTS employee.employee_verification;
