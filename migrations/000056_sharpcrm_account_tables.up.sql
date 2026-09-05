CREATE SCHEMA IF NOT EXISTS sharpcrm_engine;

-- --------------------------------------
-- sharpcrm_engine.org_sharpcrm_accounts
-- One organization gets exactly one SharpCRM account — it is the root every
-- other CRM row hangs off. The table exists (rather than scoping everything
-- straight to organization_id) so an individual, non-org CRM account stays
-- possible later without reshaping the whole schema. Same shape as
-- cables_engine.org_cables_accounts, on purpose: both are opt-in services.
--
-- is_active is the kill switch that turns SharpCRM off for the org without
-- deleting any history — a switched-off account is re-activated by the start
-- flow, never replaced. IDs are app-generated (uuid.New()).
-- --------------------------------------
CREATE TABLE sharpcrm_engine.org_sharpcrm_accounts (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,

    is_active       BOOLEAN NOT NULL DEFAULT true,

    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

-- One account per organization. No soft delete on this table, so the
-- constraint is unconditional rather than partial. It is also the race
-- backstop under the service's "already started" check.
CREATE UNIQUE INDEX uq_org_sharpcrm_accounts_organization_id
    ON sharpcrm_engine.org_sharpcrm_accounts(organization_id);

-- --------------------------------------
-- sharpcrm_engine.sharpcrm_account_settings
-- The account's tunables, kept off the account row so the hot "is CRM on for
-- this org?" read stays a single narrow row. It carries no settings columns
-- yet — the row is provisioned with the account so every later setting is an
-- ALTER plus an UPDATE, never a backfill of missing rows.
--
-- organization_id is denormalised alongside org_sharpcrm_account_id so the
-- tenant predicate never needs a join back to the account.
-- --------------------------------------
CREATE TABLE sharpcrm_engine.sharpcrm_account_settings (
    id                      UUID PRIMARY KEY,
    org_sharpcrm_account_id UUID NOT NULL,
    organization_id         UUID NOT NULL,

    last_updated_at         TIMESTAMPTZ DEFAULT now()
);

-- One settings row per account — the reciprocal of the one-account-per-org
-- index above.
CREATE UNIQUE INDEX uq_sharpcrm_account_settings_account_id
    ON sharpcrm_engine.sharpcrm_account_settings(org_sharpcrm_account_id);

-- The tenant predicate every read carries.
CREATE INDEX idx_sharpcrm_account_settings_organization_id
    ON sharpcrm_engine.sharpcrm_account_settings(organization_id);
