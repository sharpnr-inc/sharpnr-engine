-- --------------------------------------
-- legal_entity.legal_entity
-- --------------------------------------
CREATE TABLE legal_entity.legal_entity (
    id                  UUID PRIMARY KEY,
    organization_id     UUID NOT NULL,
    legal_name          VARCHAR(255) NOT NULL,
    display_name        VARCHAR(255),
    entity_type         VARCHAR(20) NOT NULL,
    incorporation_date  DATE,
    fy_start_month      INT NOT NULL DEFAULT 4,
    is_active           BOOLEAN NOT NULL DEFAULT true,
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now(),
    deleted_at          TIMESTAMPTZ
);

CREATE INDEX idx_legal_entity_organization_id ON legal_entity.legal_entity(organization_id);
CREATE INDEX idx_legal_entity_deleted_at ON legal_entity.legal_entity(deleted_at);

-- --------------------------------------
-- legal_entity.legal_entity_contact
-- --------------------------------------
CREATE TABLE legal_entity.legal_entity_contact (
    id              UUID PRIMARY KEY,
    legal_entity_id UUID NOT NULL,
    type            VARCHAR(30) NOT NULL,
    first_name      VARCHAR(100) NOT NULL,
    last_name       VARCHAR(100),
    designation     VARCHAR(100),
    email           VARCHAR(255),
    mobile          VARCHAR(15),
    din             VARCHAR(10),
    is_primary      BOOLEAN NOT NULL DEFAULT false,
    is_active       BOOLEAN NOT NULL DEFAULT true,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX idx_legal_entity_contact_legal_entity_id ON legal_entity.legal_entity_contact(legal_entity_id);
CREATE INDEX idx_legal_entity_contact_deleted_at ON legal_entity.legal_entity_contact(deleted_at);

-- --------------------------------------
-- legal_entity.legal_entity_identifier
-- --------------------------------------
CREATE TABLE legal_entity.legal_entity_identifier (
    id              UUID PRIMARY KEY,
    legal_entity_id UUID NOT NULL,
    type            VARCHAR(30) NOT NULL,
    value           VARCHAR(100) NOT NULL,
    is_verified     BOOLEAN NOT NULL DEFAULT false,
    verified_at     TIMESTAMPTZ,
    expires_at      TIMESTAMPTZ,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_legal_entity_identifier_legal_entity_id ON legal_entity.legal_entity_identifier(legal_entity_id);

-- --------------------------------------
-- legal_entity.legal_entity_bank_account
-- --------------------------------------
CREATE TABLE legal_entity.legal_entity_bank_account (
    id                   UUID PRIMARY KEY,
    legal_entity_id      UUID NOT NULL,
    account_holder_name  VARCHAR(355) NOT NULL,
    account_number       VARCHAR(512) NOT NULL,
    ifsc                 VARCHAR(11) NOT NULL,
    bank_name            VARCHAR(100),
    branch_name          VARCHAR(100),
    micr_code            VARCHAR(9),
    account_type         VARCHAR(15) NOT NULL DEFAULT 'CURRENT',
    is_primary           BOOLEAN NOT NULL DEFAULT false,
    is_verified          BOOLEAN NOT NULL DEFAULT false,
    verified_at          TIMESTAMPTZ,
    is_active            BOOLEAN NOT NULL DEFAULT true,
    created_at           TIMESTAMPTZ DEFAULT now(),
    updated_at           TIMESTAMPTZ DEFAULT now(),
    deleted_at           TIMESTAMPTZ
);

CREATE INDEX idx_legal_entity_bank_account_legal_entity_id ON legal_entity.legal_entity_bank_account(legal_entity_id);
CREATE INDEX idx_legal_entity_bank_account_deleted_at ON legal_entity.legal_entity_bank_account(deleted_at);

-- --------------------------------------
-- legal_entity.legal_entity_address
-- --------------------------------------
CREATE TABLE legal_entity.legal_entity_address (
    id              UUID PRIMARY KEY,
    legal_entity_id UUID NOT NULL,
    type            VARCHAR(20) NOT NULL,
    address_line1   VARCHAR(255) NOT NULL,
    address_line2   VARCHAR(255),
    landmark        VARCHAR(255),
    city            VARCHAR(100) NOT NULL,
    state           VARCHAR(100) NOT NULL,
    gst_state_code  VARCHAR(2),
    pincode         VARCHAR(10) NOT NULL,
    country         VARCHAR(2) NOT NULL DEFAULT 'IN',
    is_primary      BOOLEAN NOT NULL DEFAULT false,
    property_id     UUID,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX idx_legal_entity_address_legal_entity_id ON legal_entity.legal_entity_address(legal_entity_id);
CREATE INDEX idx_legal_entity_address_deleted_at ON legal_entity.legal_entity_address(deleted_at);

-- --------------------------------------
-- legal_entity.legal_entity_compliance
-- --------------------------------------
CREATE TABLE legal_entity.legal_entity_compliance (
    id                    UUID PRIMARY KEY,
    legal_entity_id       UUID NOT NULL,
    type                  VARCHAR(30) NOT NULL,
    license_number        VARCHAR(100) NOT NULL,
    issuing_authority     VARCHAR(255),
    valid_from            TIMESTAMPTZ,
    valid_to              TIMESTAMPTZ,
    renewal_reminder_days INT NOT NULL DEFAULT 30,
    property_id           UUID,
    module                VARCHAR(30),
    document_url          VARCHAR(512),
    is_active             BOOLEAN NOT NULL DEFAULT true,
    created_at            TIMESTAMPTZ DEFAULT now(),
    updated_at            TIMESTAMPTZ DEFAULT now(),
    deleted_at            TIMESTAMPTZ
);

CREATE INDEX idx_legal_entity_compliance_legal_entity_id ON legal_entity.legal_entity_compliance(legal_entity_id);
CREATE INDEX idx_legal_entity_compliance_property_id ON legal_entity.legal_entity_compliance(property_id);
CREATE INDEX idx_legal_entity_compliance_deleted_at ON legal_entity.legal_entity_compliance(deleted_at);

-- --------------------------------------
-- legal_entity.legal_entity_tax_registration
-- --------------------------------------
CREATE TABLE legal_entity.legal_entity_tax_registration (
    id                    UUID PRIMARY KEY,
    legal_entity_id       UUID NOT NULL,
    tax_type              VARCHAR(10) NOT NULL,
    registration_number   VARCHAR(20) NOT NULL,
    gst_registration_type VARCHAR(20) NOT NULL DEFAULT 'REGULAR',
    state                 VARCHAR(100),
    gst_state_code        VARCHAR(2),
    country               VARCHAR(2) NOT NULL DEFAULT 'IN',
    effective_from        TIMESTAMPTZ,
    effective_to          TIMESTAMPTZ,
    is_active             BOOLEAN NOT NULL DEFAULT true,
    created_at            TIMESTAMPTZ DEFAULT now(),
    updated_at            TIMESTAMPTZ DEFAULT now(),
    deleted_at            TIMESTAMPTZ
);

CREATE INDEX idx_legal_entity_tax_registration_legal_entity_id ON legal_entity.legal_entity_tax_registration(legal_entity_id);
CREATE INDEX idx_legal_entity_tax_registration_deleted_at ON legal_entity.legal_entity_tax_registration(deleted_at);

-- --------------------------------------
-- legal_entity.legal_entity_tax_category
-- --------------------------------------
CREATE TABLE legal_entity.legal_entity_tax_category (
    id              UUID PRIMARY KEY,
    legal_entity_id UUID NOT NULL,
    name            VARCHAR(100) NOT NULL,
    code            VARCHAR(50) NOT NULL,
    description     VARCHAR(255),
    sac_code        VARCHAR(10),
    hsn_code        VARCHAR(10),
    is_active       BOOLEAN NOT NULL DEFAULT true,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_legal_entity_tax_category_legal_entity_id ON legal_entity.legal_entity_tax_category(legal_entity_id);

-- --------------------------------------
-- legal_entity.legal_entity_tax_rule
-- --------------------------------------
CREATE TABLE legal_entity.legal_entity_tax_rule (
    id                               UUID PRIMARY KEY,
    legal_entity_id                  UUID NOT NULL,
    legal_entity_tax_category_id     UUID NOT NULL,
    legal_entity_tax_registration_id UUID NOT NULL,
    tax_type                         VARCHAR(10) NOT NULL,
    rate                             DECIMAL(5,2) NOT NULL,
    condition_json                   JSONB,
    min_value                        DECIMAL(15,2),
    max_value                        DECIMAL(15,2),
    priority                         INT NOT NULL DEFAULT 99,
    effective_from                   TIMESTAMPTZ,
    effective_to                     TIMESTAMPTZ,
    is_active                        BOOLEAN NOT NULL DEFAULT true,
    created_at                       TIMESTAMPTZ DEFAULT now(),
    updated_at                       TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_legal_entity_tax_rule_legal_entity_id ON legal_entity.legal_entity_tax_rule(legal_entity_id);
CREATE INDEX idx_legal_entity_tax_rule_tax_category_id ON legal_entity.legal_entity_tax_rule(legal_entity_tax_category_id);
CREATE INDEX idx_legal_entity_tax_rule_tax_registration_id ON legal_entity.legal_entity_tax_rule(legal_entity_tax_registration_id);

-- --------------------------------------
-- legal_entity.legal_entity_tax_component
-- --------------------------------------
CREATE TABLE legal_entity.legal_entity_tax_component (
    id                       UUID PRIMARY KEY,
    legal_entity_tax_rule_id UUID NOT NULL,
    component_code           VARCHAR(10) NOT NULL,
    rate                     DECIMAL(5,2) NOT NULL,
    liability_account_id     UUID NOT NULL,
    created_at               TIMESTAMPTZ DEFAULT now(),
    updated_at               TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_legal_entity_tax_component_tax_rule_id ON legal_entity.legal_entity_tax_component(legal_entity_tax_rule_id);

-- --------------------------------------
-- legal_entity.tax_category_account_mapping
-- --------------------------------------
CREATE TABLE legal_entity.tax_category_account_mapping (
    id                          UUID PRIMARY KEY,
    legal_entity_id             UUID NOT NULL,
    legal_entity_tax_category_id UUID NOT NULL,
    accounting_book_id          UUID NOT NULL,
    receivable_account_id       UUID NOT NULL,
    income_account_id           UUID NOT NULL,
    is_active                   BOOLEAN NOT NULL DEFAULT true,
    created_at                  TIMESTAMPTZ DEFAULT now(),
    updated_at                  TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_tax_category_account_mapping_legal_entity_id ON legal_entity.tax_category_account_mapping(legal_entity_id);
CREATE INDEX idx_tax_category_account_mapping_tax_category_id ON legal_entity.tax_category_account_mapping(legal_entity_tax_category_id);
CREATE INDEX idx_tax_category_account_mapping_accounting_book_id ON legal_entity.tax_category_account_mapping(accounting_book_id);