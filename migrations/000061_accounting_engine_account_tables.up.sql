-- The accounting_engine schema already exists (000001 onwards own the books,
-- accounts and journals). This migration only adds the account layer that makes
-- accounting an opt-in service like cables, sharpcrm, cortex, silo, folio and
-- hrms — it does not touch any existing accounting table.
CREATE SCHEMA IF NOT EXISTS accounting_engine;

-- --------------------------------------
-- accounting_engine.org_accounting_engine_accounts
-- One organization gets exactly one accounting-engine account — the switch that
-- says the org has taken the service, distinct from the accounting_book rows
-- that hang off each legal entity. The table exists (rather than scoping to
-- organization_id everywhere) so an individual, non-org account stays possible
-- later without reshaping the schema.
--
-- is_active is the kill switch that turns accounting off for the org without
-- deleting a single ledger row — a switched-off account is re-activated by the
-- start flow, never replaced. IDs are app-generated (uuid.New()).
-- --------------------------------------
CREATE TABLE accounting_engine.org_accounting_engine_accounts (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,

    is_active       BOOLEAN NOT NULL DEFAULT true,

    created_at      TIMESTAMPTZ DEFAULT now()
);

-- One account per organization. No soft delete on this table, so the
-- constraint is unconditional rather than partial. It is also the race backstop
-- under the service's "already started" check.
CREATE UNIQUE INDEX uq_org_accounting_engine_accounts_organization_id
    ON accounting_engine.org_accounting_engine_accounts(organization_id);

-- --------------------------------------
-- accounting_engine.accounting_engine_account_settings
-- The account's org-wide tunables, kept off the account row so the hot "is
-- accounting on for this org?" read stays a single narrow row. Distinct from
-- accounting_engine.accounting_book_setting, which configures one legal
-- entity's book; this row configures the service for the organization.
--
-- It carries no settings columns yet: the row is provisioned with the account
-- so every later setting is an ALTER plus an UPDATE, never a backfill of
-- missing rows.
-- --------------------------------------
CREATE TABLE accounting_engine.accounting_engine_account_settings (
    id                               UUID PRIMARY KEY,
    org_accounting_engine_account_id UUID NOT NULL,
    organization_id                  UUID NOT NULL,

    updated_at                       TIMESTAMPTZ DEFAULT now()
);

-- One settings row per account — the reciprocal of the one-account-per-org
-- index above.
CREATE UNIQUE INDEX uq_accounting_engine_account_settings_account_id
    ON accounting_engine.accounting_engine_account_settings(org_accounting_engine_account_id);

-- The tenant predicate every read carries.
CREATE INDEX idx_accounting_engine_account_settings_organization_id
    ON accounting_engine.accounting_engine_account_settings(organization_id);
