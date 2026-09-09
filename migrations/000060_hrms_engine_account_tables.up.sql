-- The HRMS engine moves onto the same shape as every other opt-in engine: its
-- own `<engine>_engine` schema, an account row per organization, and a
-- settings row hanging off it.
--
-- org_hrms_preference goes away with it. hr_model and the onboarding-approval
-- flag were a statement of intent that nothing read: authorization is decided
-- by Sentinel policy, HR tables are cost-center-scoped under either model, and
-- the table never earned the code path it implied. Anything the HR app needs
-- to know about how the org runs HR belongs on the settings row below, added
-- as a column when a reader for it actually exists.
DROP TABLE IF EXISTS hrms.org_hrms_preference;

-- The `hrms` schema is renamed to `hrms_engine` to match cables_engine,
-- sharpcrm_engine, cortex_engine, silo_engine and folio_engine. It held only
-- the preference table dropped above, so this is a drop-and-create rather than
-- an ALTER ... RENAME. No CASCADE: if a later HRMS table has been added there
-- the drop fails loudly instead of silently destroying it.
DROP SCHEMA IF EXISTS hrms;

CREATE SCHEMA IF NOT EXISTS hrms_engine;

-- --------------------------------------
-- hrms_engine.org_hrms_accounts
-- One organization gets exactly one HRMS account — the root every other HRMS
-- row hangs off. The table exists (rather than scoping everything straight to
-- organization_id) so an individual, non-org HRMS account stays possible later
-- without reshaping the schema. Same shape as the cables, sharpcrm, cortex,
-- silo and folio account tables, on purpose: all six are opt-in services.
--
-- is_active is the kill switch that turns HRMS off for the org without
-- deleting any employee record — a switched-off account is re-activated by the
-- start flow, never replaced. IDs are app-generated (uuid.New()).
-- --------------------------------------
CREATE TABLE hrms_engine.org_hrms_accounts (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,

    is_active       BOOLEAN NOT NULL DEFAULT true,

    created_at      TIMESTAMPTZ DEFAULT now()
);

-- One account per organization. No soft delete on this table, so the
-- constraint is unconditional rather than partial. It is also the race backstop
-- under the service's "already started" check.
CREATE UNIQUE INDEX uq_org_hrms_accounts_organization_id
    ON hrms_engine.org_hrms_accounts(organization_id);

-- --------------------------------------
-- hrms_engine.hrms_account_settings
-- The account's tunables — attendance, payroll and recruitment flags as those
-- flows land — kept off the account row so the hot "is HRMS on for this org?"
-- read stays a single narrow row. It carries no settings columns yet: the row
-- is provisioned with the account so every later setting is an ALTER plus an
-- UPDATE, never a backfill of missing rows.
--
-- organization_id is denormalised alongside org_hrms_account_id so the tenant
-- predicate never needs a join back to the account.
-- --------------------------------------
CREATE TABLE hrms_engine.hrms_account_settings (
    id                  UUID PRIMARY KEY,
    org_hrms_account_id UUID NOT NULL,
    organization_id     UUID NOT NULL,

    updated_at          TIMESTAMPTZ DEFAULT now()
);

-- One settings row per account — the reciprocal of the one-account-per-org
-- index above.
CREATE UNIQUE INDEX uq_hrms_account_settings_account_id
    ON hrms_engine.hrms_account_settings(org_hrms_account_id);

-- The tenant predicate every read carries.
CREATE INDEX idx_hrms_account_settings_organization_id
    ON hrms_engine.hrms_account_settings(organization_id);
