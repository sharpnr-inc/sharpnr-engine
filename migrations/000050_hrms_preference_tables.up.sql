-- The HRMS engine's own schema. Created here rather than in 000001 because the
-- engine's first owned table arrives now; attendance, payroll and recruitment
-- will land alongside it.
CREATE SCHEMA IF NOT EXISTS hrms;

-- --------------------------------------
-- hrms.org_hrms_preference
-- One organization's HRMS configuration — a single row per org, created with
-- defaults the first time the org touches HR.
--
-- hr_model is a statement of INTENT, not an access-control switch. Authorization
-- is always decided by Sentinel policy: a central HR carries a type-level grant
-- (every cost center), a local HR an instance-level grant (one cost center).
-- HR tables are cost-center-scoped under both models, so moving from CENTRALIZED
-- to DISTRIBUTED is a change of policy attachments and never a backfill.
--
-- What it does drive: what the HR app renders (a cost-center picker or not),
-- what policy scope is attached by default when an HR role is created, and the
-- warning shown when a configuration contradicts the declared model.
-- --------------------------------------
CREATE TABLE hrms.org_hrms_preference (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,

    hr_model        VARCHAR(20) NOT NULL DEFAULT 'CENTRALIZED',

    -- Only bites under DISTRIBUTED: may a local HR complete an onboarding on
    -- their own, or does central HR sign off? Ignored under CENTRALIZED, where
    -- there is no local HR to gate.
    require_central_approval_for_onboarding BOOLEAN NOT NULL DEFAULT false,

    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

-- One row per organization, and the only lookup this table ever serves.
CREATE UNIQUE INDEX uq_org_hrms_preference_organization_id
    ON hrms.org_hrms_preference(organization_id);
