CREATE SCHEMA IF NOT EXISTS folio_engine;

-- --------------------------------------
-- folio_engine.org_folio_accounts
-- One organization gets exactly one Folio account — it is the root every other
-- Folio row hangs off. The table exists (rather than scoping everything
-- straight to organization_id) so an individual, non-org Folio account stays
-- possible later without reshaping the whole schema. Same shape as the cables,
-- sharpcrm, cortex and silo account tables, on purpose: all five are opt-in
-- services.
--
-- is_active is the kill switch that turns Folio off for the org without
-- deleting any document — a switched-off account is re-activated by the start
-- flow, never replaced. IDs are app-generated (uuid.New()).
-- --------------------------------------
CREATE TABLE folio_engine.org_folio_accounts (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,

    is_active       BOOLEAN NOT NULL DEFAULT true,

    created_at      TIMESTAMPTZ DEFAULT now()
);

-- One account per organization. No soft delete on this table, so the
-- constraint is unconditional rather than partial. It is also the race backstop
-- under the service's "already started" check.
CREATE UNIQUE INDEX uq_org_folio_accounts_organization_id
    ON folio_engine.org_folio_accounts(organization_id);

-- --------------------------------------
-- folio_engine.folio_account_settings
-- The account's tunables — default sharing, templates, retention — kept off the
-- account row so the hot "is Folio on for this org?" read stays a single narrow
-- row. It carries no settings columns yet: the row is provisioned with the
-- account so every later setting is an ALTER plus an UPDATE, never a backfill
-- of missing rows.
--
-- organization_id is denormalised alongside org_folio_account_id so the tenant
-- predicate never needs a join back to the account.
-- --------------------------------------
CREATE TABLE folio_engine.folio_account_settings (
    id                   UUID PRIMARY KEY,
    org_folio_account_id UUID NOT NULL,
    organization_id      UUID NOT NULL,

    updated_at           TIMESTAMPTZ DEFAULT now()
);

-- One settings row per account — the reciprocal of the one-account-per-org
-- index above.
CREATE UNIQUE INDEX uq_folio_account_settings_account_id
    ON folio_engine.folio_account_settings(org_folio_account_id);

-- The tenant predicate every read carries.
CREATE INDEX idx_folio_account_settings_organization_id
    ON folio_engine.folio_account_settings(organization_id);
