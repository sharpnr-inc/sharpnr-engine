CREATE SCHEMA IF NOT EXISTS cortex_engine;

-- --------------------------------------
-- cortex_engine.org_cortex_accounts
-- One organization gets exactly one Cortex account — it is the root every
-- other Cortex row hangs off. The table exists (rather than scoping everything
-- straight to organization_id) so an individual, non-org Cortex account stays
-- possible later without reshaping the whole schema. Same shape as
-- cables_engine.org_cables_accounts and sharpcrm_engine.org_sharpcrm_accounts,
-- on purpose: all three are opt-in services.
--
-- is_active is the kill switch that turns Cortex off for the org without
-- deleting any history — a switched-off account is re-activated by the start
-- flow, never replaced. IDs are app-generated (uuid.New()).
--
-- No updated_at column: the model carries none, and the account lifecycle is
-- already recorded on the org's audit trail.
-- --------------------------------------
CREATE TABLE cortex_engine.org_cortex_accounts (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,

    is_active       BOOLEAN NOT NULL DEFAULT true,

    created_at      TIMESTAMPTZ DEFAULT now()
);

-- One account per organization. No soft delete on this table, so the
-- constraint is unconditional rather than partial. It is also the race
-- backstop under the service's "already started" check.
CREATE UNIQUE INDEX uq_org_cortex_accounts_organization_id
    ON cortex_engine.org_cortex_accounts(organization_id);

-- --------------------------------------
-- cortex_engine.cortex_account_settings
-- The account's tunables, kept off the account row so the hot "is Cortex on
-- for this org?" read stays a single narrow row. It carries no settings
-- columns yet — the row is provisioned with the account so every later setting
-- is an ALTER plus an UPDATE, never a backfill of missing rows.
--
-- organization_id is denormalised alongside org_cortex_account_id so the
-- tenant predicate never needs a join back to the account.
-- --------------------------------------
CREATE TABLE cortex_engine.cortex_account_settings (
    id                    UUID PRIMARY KEY,
    org_cortex_account_id UUID NOT NULL,
    organization_id       UUID NOT NULL,

    updated_at            TIMESTAMPTZ DEFAULT now()
);

-- One settings row per account — the reciprocal of the one-account-per-org
-- index above.
CREATE UNIQUE INDEX uq_cortex_account_settings_account_id
    ON cortex_engine.cortex_account_settings(org_cortex_account_id);

-- The tenant predicate every read carries.
CREATE INDEX idx_cortex_account_settings_organization_id
    ON cortex_engine.cortex_account_settings(organization_id);
