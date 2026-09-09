-- ============================================================
-- DOWN: drop the account layer, rename stockwave_engine back to inventory, and
-- rebuild the five tables 000013 created.
--
-- The table definitions below are a verbatim copy of
-- 000013_inventory_engine_tables.up.sql, so rolling back to 000061 leaves the
-- database in exactly the shape 000013 left it in. Row data is NOT restored:
-- the up migration dropped it, and there is nowhere to read it back from.
-- ============================================================

BEGIN;

-- ------------------------------------------
-- 1) The account layer goes first, so it is not carried into `inventory` by
--    the rename below.
-- ------------------------------------------
DROP TABLE IF EXISTS stockwave_engine.stockwave_engine_account_settings;
DROP TABLE IF EXISTS stockwave_engine.org_stockwave_engine_accounts;

-- ------------------------------------------
-- 2) stockwave_engine -> inventory. Guarded the same way as the up: ALTER
--    SCHEMA ... RENAME has no IF EXISTS form.
-- ------------------------------------------
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_namespace WHERE nspname = 'stockwave_engine')
   AND NOT EXISTS (SELECT 1 FROM pg_namespace WHERE nspname = 'inventory')
    THEN
        EXECUTE 'ALTER SCHEMA stockwave_engine RENAME TO inventory';
    END IF;
END
$$;

CREATE SCHEMA IF NOT EXISTS inventory;

-- --------------------------------------
-- inventory.inventory_accounts
-- --------------------------------------
CREATE TABLE inventory.inventory_accounts (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    name            VARCHAR(255) NOT NULL,
    type            VARCHAR(20),
    is_shared       BOOLEAN DEFAULT false,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX idx_inventory_accounts_organization_id ON inventory.inventory_accounts(organization_id);
CREATE INDEX idx_inventory_accounts_deleted_at ON inventory.inventory_accounts(deleted_at);

-- --------------------------------------
-- inventory.inventories
-- --------------------------------------
CREATE TABLE inventory.inventories (
    id                   UUID PRIMARY KEY,
    inventory_account_id UUID NOT NULL,
    business_unit_id     UUID NOT NULL,
    legal_entity_id      UUID NOT NULL,
    name                 VARCHAR(255) NOT NULL,
    description          TEXT,
    location             VARCHAR(255),
    status               VARCHAR(20) DEFAULT 'ACTIVE',
    created_at           TIMESTAMPTZ DEFAULT now(),
    updated_at           TIMESTAMPTZ DEFAULT now(),
    deleted_at           TIMESTAMPTZ
);

CREATE INDEX idx_inventories_inventory_account_id ON inventory.inventories(inventory_account_id);
CREATE INDEX idx_inventories_business_unit_id ON inventory.inventories(business_unit_id);
CREATE INDEX idx_inventories_legal_entity_id ON inventory.inventories(legal_entity_id);
CREATE INDEX idx_inventories_deleted_at ON inventory.inventories(deleted_at);

-- --------------------------------------
-- inventory.inventory_items
-- --------------------------------------
CREATE TABLE inventory.inventory_items (
    id               UUID PRIMARY KEY,
    inventory_id     UUID NOT NULL,
    supplier_id      UUID,
    cogs_account_id  UUID,
    asset_account_id UUID,
    name             VARCHAR(255) NOT NULL,
    sku              VARCHAR(100),
    description      TEXT,
    category         VARCHAR(30),
    unit             VARCHAR(20),
    quantity         DECIMAL(15,4) DEFAULT 0,
    min_stock_level  DECIMAL(15,4) DEFAULT 0,
    max_stock_level  DECIMAL(15,4) DEFAULT 0,
    status           VARCHAR(20) DEFAULT 'ACTIVE',
    created_at       TIMESTAMPTZ DEFAULT now(),
    updated_at       TIMESTAMPTZ DEFAULT now(),
    deleted_at       TIMESTAMPTZ
);

CREATE INDEX idx_inventory_items_inventory_id ON inventory.inventory_items(inventory_id);
CREATE INDEX idx_inventory_items_supplier_id ON inventory.inventory_items(supplier_id);
CREATE INDEX idx_inventory_items_deleted_at ON inventory.inventory_items(deleted_at);
CREATE UNIQUE INDEX idx_inventory_items_sku ON inventory.inventory_items(sku);

-- --------------------------------------
-- inventory.inventory_item_prices
-- --------------------------------------
CREATE TABLE inventory.inventory_item_prices (
    id                UUID PRIMARY KEY,
    inventory_item_id UUID NOT NULL,
    supplier_id       UUID,
    price_type        VARCHAR(20),
    price             DECIMAL(15,2) NOT NULL,
    currency          VARCHAR(5) DEFAULT 'INR',
    effective_from    TIMESTAMPTZ NOT NULL,
    effective_to      TIMESTAMPTZ,
    created_at        TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_inventory_item_prices_inventory_item_id ON inventory.inventory_item_prices(inventory_item_id);
CREATE INDEX idx_inventory_item_prices_supplier_id ON inventory.inventory_item_prices(supplier_id);

-- --------------------------------------
-- inventory.inventory_item_histories
-- --------------------------------------
CREATE TABLE inventory.inventory_item_histories (
    id                UUID PRIMARY KEY,
    inventory_item_id UUID NOT NULL,
    change_type       VARCHAR(20) NOT NULL,
    reference_type    VARCHAR(20),
    reference_id      UUID,
    quantity_before   DECIMAL(15,4),
    quantity_change   DECIMAL(15,4),
    quantity_after    DECIMAL(15,4),
    unit_cost         DECIMAL(15,2),
    total_cost        DECIMAL(15,2),
    performed_by      UUID,
    notes             TEXT,
    created_at        TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_inventory_item_histories_inventory_item_id ON inventory.inventory_item_histories(inventory_item_id);

COMMIT;
