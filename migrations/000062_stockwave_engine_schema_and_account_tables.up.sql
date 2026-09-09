-- ============================================================
-- Stockwave engine: rename the `inventory` schema to `stockwave_engine`, drop
-- the five inventory tables 000013 created, and lay down the account layer.
--
-- The engine is being rewritten. Everything 000013 built (accounts,
-- inventories, items, prices, histories) is replaced by the new design, so the
-- tables go rather than being migrated forward — none of them carries data
-- worth preserving into a schema whose shape has not been settled yet.
--
-- The schema is RENAMED rather than dropped-and-recreated so the reversal in
-- 000062.down.sql is exact and so `inventory` stops existing in one step: no
-- window where both names are live and models could write to the wrong one.
--
-- ⚠ DEPLOY ORDER: the application must STOP READING AND WRITING the inventory
-- tables before this migration runs. The models that map to them
-- (stockwave_engine/models/inventory.go, purchase_order.go) and everything that
-- reaches them through stockwave_engine/repo must be gone from the build first.
-- ============================================================

BEGIN;

-- ------------------------------------------
-- 1) inventory -> stockwave_engine
--
-- ALTER SCHEMA ... RENAME has no IF EXISTS form, so the guard is explicit. It
-- is a no-op on a database where the rename already happened (or where
-- `inventory` was never created), which keeps this migration safe to re-run
-- after a forced version fix.
-- ------------------------------------------
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_namespace WHERE nspname = 'inventory')
   AND NOT EXISTS (SELECT 1 FROM pg_namespace WHERE nspname = 'stockwave_engine')
    THEN
        EXECUTE 'ALTER SCHEMA inventory RENAME TO stockwave_engine';
    END IF;
END
$$;

-- Covers the case where `inventory` never existed (a database created after
-- 000001 was amended, or a partially-applied history).
CREATE SCHEMA IF NOT EXISTS stockwave_engine;

-- ------------------------------------------
-- 2) Drop 000013's tables.
--
-- They arrive here already renamed into stockwave_engine by step 1. Dropped
-- child-first even though there are no FK constraints between them, so the
-- order still reads as the dependency order the data had.
-- ------------------------------------------
DROP TABLE IF EXISTS stockwave_engine.inventory_item_histories;
DROP TABLE IF EXISTS stockwave_engine.inventory_item_prices;
DROP TABLE IF EXISTS stockwave_engine.inventory_items;
DROP TABLE IF EXISTS stockwave_engine.inventories;
DROP TABLE IF EXISTS stockwave_engine.inventory_accounts;

-- --------------------------------------
-- 3) stockwave_engine.org_stockwave_engine_accounts
--
-- One organization gets exactly one stockwave-engine account — the switch that
-- says the org has taken the service. The table exists (rather than scoping to
-- organization_id everywhere) so an individual, non-org account stays possible
-- later without reshaping the schema.
--
-- is_active is the kill switch that turns stockwave off for the org without
-- deleting a single stock row — a switched-off account is re-activated by the
-- start flow, never replaced. IDs are app-generated (uuid.New()).
-- --------------------------------------
CREATE TABLE stockwave_engine.org_stockwave_engine_accounts (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,

    is_active       BOOLEAN NOT NULL DEFAULT true,

    created_at      TIMESTAMPTZ DEFAULT now()
);

-- One account per organization. No soft delete on this table, so the
-- constraint is unconditional rather than partial. It is also the race backstop
-- under the service's "already started" check.
CREATE UNIQUE INDEX uq_org_stockwave_engine_accounts_organization_id
    ON stockwave_engine.org_stockwave_engine_accounts(organization_id);

-- --------------------------------------
-- 4) stockwave_engine.stockwave_engine_account_settings
--
-- The account's org-wide tunables, kept off the account row so the hot "is
-- stockwave on for this org?" read stays a single narrow row.
--
-- It carries no settings columns yet: the row is provisioned with the account
-- so every later setting is an ALTER plus an UPDATE, never a backfill of
-- missing rows. organization_id is denormalised alongside
-- org_stockwave_engine_account_id so the tenant predicate never needs a join.
-- --------------------------------------
CREATE TABLE stockwave_engine.stockwave_engine_account_settings (
    id                              UUID PRIMARY KEY,
    org_stockwave_engine_account_id UUID NOT NULL,
    organization_id                 UUID NOT NULL,

    updated_at                      TIMESTAMPTZ DEFAULT now()
);

-- One settings row per account — the reciprocal of the one-account-per-org
-- index above.
CREATE UNIQUE INDEX uq_stockwave_engine_account_settings_account_id
    ON stockwave_engine.stockwave_engine_account_settings(org_stockwave_engine_account_id);

-- The tenant predicate every read carries.
CREATE INDEX idx_stockwave_engine_account_settings_organization_id
    ON stockwave_engine.stockwave_engine_account_settings(organization_id);

COMMIT;
