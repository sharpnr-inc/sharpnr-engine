-- core.organization
CREATE TABLE core.organization (
    id          UUID PRIMARY KEY NOT NULL,
    name        VARCHAR(100) NOT NULL,
    sharpnr_id  VARCHAR(50) NOT NULL,
    status      VARCHAR(100),
    root_email  VARCHAR(255),
    created_at  TIMESTAMPTZ DEFAULT now(),
    updated_at  TIMESTAMPTZ DEFAULT now(),
    deleted_at  TIMESTAMPTZ
);

CREATE UNIQUE INDEX idx_organization_sharpnr_id ON core.organization(sharpnr_id);
CREATE INDEX idx_organization_root_email ON core.organization(root_email);
CREATE INDEX idx_organization_deleted_at ON core.organization(deleted_at);

-- core.organization_verification
CREATE TABLE core.organization_verification (
    id                  UUID PRIMARY KEY,
    organization_id     UUID UNIQUE,
    email_verified      BOOLEAN DEFAULT false,
    mobile_verified     BOOLEAN DEFAULT false,
    physical_verified   BOOLEAN DEFAULT false,
    verified_at         TIMESTAMPTZ,
    verified_by         UUID,
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now()
);

-- core.organization_subscription
CREATE TABLE core.organization_subscription (
    id              UUID PRIMARY KEY,
    organization_id UUID,
    plan_code       VARCHAR(255),
    status          VARCHAR(255),
    renews_at       TIMESTAMPTZ,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_organization_subscription_org_id ON core.organization_subscription(organization_id);