CREATE TABLE setup_wizard.organization_setup (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    status          VARCHAR(20) NOT NULL,
    current_step    VARCHAR(50) NOT NULL,
    version         INT DEFAULT 1,
    started_at      TIMESTAMPTZ,
    completed_at    TIMESTAMPTZ,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_organization_setup_organization_id ON setup_wizard.organization_setup(organization_id);