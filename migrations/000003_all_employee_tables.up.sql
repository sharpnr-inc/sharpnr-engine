-- --------------------------------------
-- employee.employee (base table first)
-- --------------------------------------
CREATE TABLE employee.employee (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    first_name      VARCHAR(50) NOT NULL,
    middle_name     VARCHAR(50),
    last_name       VARCHAR(50),
    nick_name       VARCHAR(50),
    email           VARCHAR(150) NOT NULL,
    phone_number    VARCHAR(20),
    role            VARCHAR(50) NOT NULL,
    status          VARCHAR(20) DEFAULT 'ACTIVE',
    employee_code   VARCHAR(50),
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX idx_employee_organization_id ON employee.employee(organization_id);
CREATE INDEX idx_employee_email ON employee.employee(email);
CREATE INDEX idx_employee_deleted_at ON employee.employee(deleted_at);

-- --------------------------------------
-- employee.employee_auth
-- --------------------------------------
CREATE TABLE employee.employee_auth (
    id              UUID PRIMARY KEY,
    employee_id     UUID NOT NULL,
    username        VARCHAR(100) NOT NULL,
    email           VARCHAR(150) NOT NULL,
    password_hash   TEXT NOT NULL,
    is_active       BOOLEAN DEFAULT true,
    is_provisioned  BOOLEAN DEFAULT false,
    last_login_at   TIMESTAMPTZ,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_employee_auth_employee_id ON employee.employee_auth(employee_id);
CREATE UNIQUE INDEX idx_employee_auth_username ON employee.employee_auth(username);
CREATE UNIQUE INDEX idx_employee_auth_email ON employee.employee_auth(email);

-- --------------------------------------
-- employee.employee_verification
-- --------------------------------------
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

-- --------------------------------------
-- employee.employee_session
-- --------------------------------------
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

-- --------------------------------------
-- employee.employee_system_access
-- --------------------------------------
CREATE TABLE employee.employee_system_access (
    id          UUID PRIMARY KEY,
    employee_id UUID,
    app         VARCHAR(50),
    online      BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ DEFAULT now(),
    updated_at  TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_employee_system_access_employee_id ON employee.employee_system_access(employee_id);

-- --------------------------------------
-- employee.employee_onboarding
-- --------------------------------------
CREATE TABLE employee.employee_onboarding (
    id              UUID PRIMARY KEY,
    employee_id     UUID NOT NULL,
    status          VARCHAR(50) NOT NULL,
    last_updated_by UUID,
    notes           TEXT,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_employee_onboarding_employee_id ON employee.employee_onboarding(employee_id);
CREATE INDEX idx_employee_onboarding_status ON employee.employee_onboarding(status);

-- --------------------------------------
-- employee.employee_profile
-- --------------------------------------
CREATE TABLE employee.employee_profile (
    id                UUID PRIMARY KEY,
    employee_id       UUID NOT NULL,
    gender            VARCHAR(10),
    date_of_birth     DATE,
    address           TEXT,
    emergency_contact VARCHAR(100),
    blood_group       VARCHAR(5),
    created_at        TIMESTAMPTZ DEFAULT now(),
    updated_at        TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_employee_profile_employee_id ON employee.employee_profile(employee_id);

-- --------------------------------------
-- employee.employee_document
-- --------------------------------------
CREATE TABLE employee.employee_document (
    id          UUID PRIMARY KEY,
    employee_id UUID NOT NULL,
    type        VARCHAR(50),
    url         TEXT,
    verified    BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_employee_document_employee_id ON employee.employee_document(employee_id);

-- --------------------------------------
-- employee.employee_bank_account
-- --------------------------------------
CREATE TABLE employee.employee_bank_account (
    id             UUID PRIMARY KEY,
    employee_id    UUID NOT NULL,
    bank_name      VARCHAR(100),
    account_number VARCHAR(50),
    ifsc           VARCHAR(20),
    upi            VARCHAR(100),
    is_primary     BOOLEAN DEFAULT true,
    created_at     TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_employee_bank_account_employee_id ON employee.employee_bank_account(employee_id);

-- --------------------------------------
-- employee.employee_reporting
-- --------------------------------------
CREATE TABLE employee.employee_reporting (
    id          UUID PRIMARY KEY,
    employee_id UUID NOT NULL,
    manager_id  UUID,
    level       SMALLINT,
    created_at  TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_employee_reporting_employee_id ON employee.employee_reporting(employee_id);
CREATE INDEX idx_employee_reporting_manager_id ON employee.employee_reporting(manager_id);

-- --------------------------------------
-- employee.employee_employment
-- --------------------------------------
CREATE TABLE employee.employee_employment (
    id                UUID PRIMARY KEY,
    employee_id       UUID NOT NULL,
    employment_type   VARCHAR(30) NOT NULL,
    designation       VARCHAR(100) NOT NULL,
    confirmation_date DATE,
    joined_at         DATE NOT NULL,
    created_at        TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_employee_employment_employee_id ON employee.employee_employment(employee_id);

-- --------------------------------------
-- employee.employee_compensation
-- --------------------------------------
CREATE TABLE employee.employee_compensation (
    id             UUID PRIMARY KEY,
    employee_id    UUID NOT NULL,
    base_salary    DECIMAL(15,2) NOT NULL,
    currency       VARCHAR(3) DEFAULT 'USD',
    effective_from DATE NOT NULL,
    effective_to   DATE,
    created_at     TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_employee_compensation_employee_id ON employee.employee_compensation(employee_id);

-- --------------------------------------
-- employee.employee_salary_component
-- --------------------------------------
CREATE TABLE employee.employee_salary_component (
    id              UUID PRIMARY KEY,
    compensation_id UUID NOT NULL,
    component       VARCHAR(50) NOT NULL,
    amount          DECIMAL(15,2) NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_employee_salary_component_compensation_id ON employee.employee_salary_component(compensation_id);

-- --------------------------------------
-- employee.employee_leave_type
-- --------------------------------------
CREATE TABLE employee.employee_leave_type (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    name            VARCHAR(50) NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_employee_leave_type_organization_id ON employee.employee_leave_type(organization_id);

-- --------------------------------------
-- employee.employee_leave_balance
-- --------------------------------------
CREATE TABLE employee.employee_leave_balance (
    id            UUID PRIMARY KEY,
    employee_id   UUID NOT NULL,
    leave_type_id UUID NOT NULL,
    balance       DECIMAL(5,2) DEFAULT 0,
    updated_at    TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_employee_leave_balance_employee_id ON employee.employee_leave_balance(employee_id);
CREATE INDEX idx_employee_leave_balance_leave_type_id ON employee.employee_leave_balance(leave_type_id);

-- --------------------------------------
-- employee.employee_leave_request
-- --------------------------------------
CREATE TABLE employee.employee_leave_request (
    id            UUID PRIMARY KEY,
    employee_id   UUID NOT NULL,
    leave_type_id UUID NOT NULL,
    from_date     DATE NOT NULL,
    to_date       DATE NOT NULL,
    status        VARCHAR(20) DEFAULT 'PENDING',
    approved_by   UUID,
    approved_at   TIMESTAMPTZ,
    reason        TEXT,
    created_at    TIMESTAMPTZ DEFAULT now(),
    updated_at    TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_employee_leave_request_employee_id ON employee.employee_leave_request(employee_id);
CREATE INDEX idx_employee_leave_request_leave_type_id ON employee.employee_leave_request(leave_type_id);
CREATE INDEX idx_employee_leave_request_status ON employee.employee_leave_request(status);

-- --------------------------------------
-- employee.employee_organization_profile
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