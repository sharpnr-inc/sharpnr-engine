-- --------------------------------------
-- organization_property.organization_property
-- --------------------------------------
CREATE TABLE organization_property.organization_property (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    code            VARCHAR(50) NOT NULL,
    name            VARCHAR(150) NOT NULL,
    is_configured   BOOLEAN DEFAULT false,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_organization_property_organization_id ON organization_property.organization_property(organization_id);
CREATE INDEX idx_organization_property_code ON organization_property.organization_property(code);

-- --------------------------------------
-- organization_property.organization_property_address
-- --------------------------------------
CREATE TABLE organization_property.organization_property_address (
    id                       UUID PRIMARY KEY,
    organization_id          UUID NOT NULL,
    organization_property_id UUID NOT NULL,
    address_line             TEXT NOT NULL,
    city                     VARCHAR(100),
    city_code                VARCHAR(20),
    state                    VARCHAR(100),
    country                  VARCHAR(100),
    pincode                  VARCHAR(20),
    created_at               TIMESTAMPTZ DEFAULT now(),
    updated_at               TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_org_property_address_organization_id ON organization_property.organization_property_address(organization_id);
CREATE INDEX idx_org_property_address_property_id ON organization_property.organization_property_address(organization_property_id);
CREATE INDEX idx_org_property_address_city ON organization_property.organization_property_address(city);
CREATE INDEX idx_org_property_address_state ON organization_property.organization_property_address(state);
CREATE INDEX idx_org_property_address_country ON organization_property.organization_property_address(country);
CREATE INDEX idx_org_property_address_pincode ON organization_property.organization_property_address(pincode);

-- --------------------------------------
-- organization_property.organization_property_timezone
-- --------------------------------------
CREATE TABLE organization_property.organization_property_timezone (
    id                       UUID PRIMARY KEY,
    organization_property_id UUID NOT NULL,
    timezone                 VARCHAR(50) NOT NULL,
    created_at               TIMESTAMPTZ DEFAULT now(),
    updated_at               TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_org_property_timezone_property_id ON organization_property.organization_property_timezone(organization_property_id);

-- --------------------------------------
-- organization_property.organization_property_geo_location
-- --------------------------------------
CREATE TABLE organization_property.organization_property_geo_location (
    id                       UUID PRIMARY KEY,
    organization_property_id UUID,
    latitude                 DECIMAL(9,6) NOT NULL,
    longitude                DECIMAL(9,6) NOT NULL,
    accuracy_m               INT,
    source                   VARCHAR(20),
    captured_at              TIMESTAMPTZ,
    updated_at               TIMESTAMPTZ DEFAULT now()
);

-- --------------------------------------
-- organization_property.organization_property_domain
-- --------------------------------------
CREATE TABLE organization_property.organization_property_domain (
    id                       UUID PRIMARY KEY,
    organization_property_id UUID NOT NULL,
    domain                   VARCHAR(50) NOT NULL,
    status                   VARCHAR(20) NOT NULL,
    created_at               TIMESTAMPTZ DEFAULT now(),
    updated_at               TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_org_property_domain_property_id ON organization_property.organization_property_domain(organization_property_id);
CREATE INDEX idx_org_property_domain_domain ON organization_property.organization_property_domain(domain);
CREATE INDEX idx_org_property_domain_status ON organization_property.organization_property_domain(status);

-- --------------------------------------
-- organization_property.organization_property_domain_modules
-- --------------------------------------
CREATE TABLE organization_property.organization_property_domain_modules (
    id                              UUID PRIMARY KEY,
    organization_property_id        UUID,
    organization_property_domain_id UUID,
    module                          VARCHAR(100),
    module_id                       UUID
);

CREATE INDEX idx_org_property_domain_modules_property_id ON organization_property.organization_property_domain_modules(organization_property_id);
CREATE INDEX idx_org_property_domain_modules_domain_id ON organization_property.organization_property_domain_modules(organization_property_domain_id);

-- --------------------------------------
-- organization_property.organization_property_domain_infra
-- --------------------------------------
CREATE TABLE organization_property.organization_property_domain_infra (
    id                              UUID PRIMARY KEY,
    organization_property_id        UUID,
    organization_property_domain_id UUID,
    infra_name                      VARCHAR(255),
    infra_type                      VARCHAR(50),
    is_shared                       BOOLEAN,
    infra_module_id                 UUID
);

CREATE INDEX idx_org_property_domain_infra_property_id ON organization_property.organization_property_domain_infra(organization_property_id);
CREATE INDEX idx_org_property_domain_infra_domain_id ON organization_property.organization_property_domain_infra(organization_property_domain_id);