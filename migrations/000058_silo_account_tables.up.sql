CREATE SCHEMA IF NOT EXISTS silo_engine;

-- --------------------------------------
-- silo_engine.org_silo_accounts
-- One organization gets exactly one Silo account — it is the root every other
-- Silo row hangs off. The table exists (rather than scoping everything straight
-- to organization_id) so an individual, non-org Silo account stays possible
-- later without reshaping the whole schema. Same shape as the cables, sharpcrm
-- and cortex account tables, on purpose: all four are opt-in services.
--
-- is_active is the kill switch that turns Silo off for the org without deleting
-- any stored data — a switched-off account is re-activated by the start flow,
-- never replaced. IDs are app-generated (uuid.New()).
--
-- No updated_at column: the model carries none, and the account lifecycle is
-- already recorded on the org's audit trail.
-- --------------------------------------
CREATE TABLE silo_engine.org_silo_accounts (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,

    is_active       BOOLEAN NOT NULL DEFAULT true,

    created_at      TIMESTAMPTZ DEFAULT now()
);

-- One account per organization. No soft delete on this table, so the
-- constraint is unconditional rather than partial. It is also the race backstop
-- under the service's "already started" check.
CREATE UNIQUE INDEX uq_org_silo_accounts_organization_id
    ON silo_engine.org_silo_accounts(organization_id);

-- --------------------------------------
-- silo_engine.silo_account_settings
-- The account's tunables — regions, defaults, retention — kept off the account
-- row so the hot "is Silo on for this org?" read stays a single narrow row. It
-- carries no settings columns yet: the row is provisioned with the account so
-- every later setting is an ALTER plus an UPDATE, never a backfill of missing
-- rows.
--
-- organization_id is denormalised alongside org_silo_account_id so the tenant
-- predicate never needs a join back to the account.
-- --------------------------------------
CREATE TABLE silo_engine.silo_account_settings (
    id                  UUID PRIMARY KEY,
    org_silo_account_id UUID NOT NULL,
    organization_id     UUID NOT NULL,

    updated_at          TIMESTAMPTZ DEFAULT now()
);

-- One settings row per account — the reciprocal of the one-account-per-org
-- index above.
CREATE UNIQUE INDEX uq_silo_account_settings_account_id
    ON silo_engine.silo_account_settings(org_silo_account_id);

-- The tenant predicate every read carries.
CREATE INDEX idx_silo_account_settings_organization_id
    ON silo_engine.silo_account_settings(organization_id);
