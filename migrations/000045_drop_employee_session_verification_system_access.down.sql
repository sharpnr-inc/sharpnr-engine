CREATE TABLE employee.employee_verification (
    id          UUID PRIMARY KEY,
    employee_id UUID,
    email       VARCHAR(150),
    type        VARCHAR(20),
    token_hash  TEXT,
    expires_at  TIMESTAMPTZ,
    verified_at TIMESTAMPTZ,
    created_at  TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_employee_verification_employee_id ON employee.employee_verification(employee_id);
CREATE INDEX idx_employee_verification_email ON employee.employee_verification(email);

CREATE TABLE employee.employee_session (
    id          UUID PRIMARY KEY,
    employee_id UUID,
    app         VARCHAR(50),
    device_info TEXT,
    expires_at  TIMESTAMPTZ,
    revoked_at  TIMESTAMPTZ,
    created_at  TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_employee_session_employee_id ON employee.employee_session(employee_id);

CREATE TABLE employee.employee_system_access (
    id          UUID PRIMARY KEY,
    employee_id UUID,
    app         VARCHAR(50),
    online      BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ DEFAULT now(),
    updated_at  TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_employee_system_access_employee_id ON employee.employee_system_access(employee_id);
