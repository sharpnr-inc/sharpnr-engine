DROP TABLE IF EXISTS hrms_engine.hrms_account_settings;
DROP TABLE IF EXISTS hrms_engine.org_hrms_accounts;

-- No CASCADE: if a later HRMS table is still present the drop fails loudly
-- rather than silently destroying it.
DROP SCHEMA IF EXISTS hrms_engine;

-- Put the `hrms` schema and its preference table back exactly as 000050 left
-- them, so rolling this migration back lands on that migration's state. Rows
-- are not recoverable — the up migration drops them.
CREATE SCHEMA IF NOT EXISTS hrms;

CREATE TABLE hrms.org_hrms_preference (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,

    hr_model        VARCHAR(20) NOT NULL DEFAULT 'CENTRALIZED',

    require_central_approval_for_onboarding BOOLEAN NOT NULL DEFAULT false,

    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX uq_org_hrms_preference_organization_id
    ON hrms.org_hrms_preference(organization_id);
