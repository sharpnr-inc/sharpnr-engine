-- --------------------------------------
-- hospitality.restaurant
-- --------------------------------------
CREATE TABLE hospitality.restaurant (
    id              UUID PRIMARY KEY,
    name            VARCHAR(100) NOT NULL,
    organization_id UUID NOT NULL,
    property_id     UUID NOT NULL,
    legal_entity_id UUID,
    is_pure_veg     BOOLEAN DEFAULT false,
    opening_time    VARCHAR(5) NOT NULL,
    closing_time    VARCHAR(5) NOT NULL,
    is_overnight    BOOLEAN DEFAULT false NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_restaurant_organization_id ON hospitality.restaurant(organization_id);
CREATE INDEX idx_restaurant_property_id ON hospitality.restaurant(property_id);
CREATE INDEX idx_restaurant_legal_entity_id ON hospitality.restaurant(legal_entity_id);

-- --------------------------------------
-- hospitality.restaurant_floor
-- --------------------------------------
CREATE TABLE hospitality.restaurant_floor (
    id            UUID PRIMARY KEY,
    restaurant_id UUID NOT NULL,
    name          VARCHAR(100),
    floor_number  INT,
    created_at    TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_restaurant_floor_restaurant_id ON hospitality.restaurant_floor(restaurant_id);

-- --------------------------------------
-- hospitality.restaurant_table
-- --------------------------------------
CREATE TABLE hospitality.restaurant_table (
    id            UUID PRIMARY KEY,
    restaurant_id UUID NOT NULL,
    floor_id      UUID NOT NULL,
    code          VARCHAR(30) NOT NULL,
    name          VARCHAR(100),
    capacity      INT,
    status        VARCHAR(20) DEFAULT 'AVAILABLE',
    created_at    TIMESTAMPTZ DEFAULT now(),
    updated_at    TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_restaurant_table_restaurant_id ON hospitality.restaurant_table(restaurant_id);
CREATE INDEX idx_restaurant_table_floor_id ON hospitality.restaurant_table(floor_id);

-- --------------------------------------
-- hospitality.restaurant_settings
-- --------------------------------------
CREATE TABLE hospitality.restaurant_settings (
    id            UUID PRIMARY KEY,
    restaurant_id UUID NOT NULL,
    data          JSONB NOT NULL DEFAULT '{}'
);

CREATE INDEX idx_restaurant_settings_restaurant_id ON hospitality.restaurant_settings(restaurant_id);

-- --------------------------------------
-- hospitality.restaurant_order_numbering_settings
-- --------------------------------------
CREATE TABLE hospitality.restaurant_order_numbering_settings (
    id                      UUID PRIMARY KEY,
    restaurant_settings_id  UUID NOT NULL,
    prefix                  VARCHAR(20) DEFAULT 'ORD',
    start_from              BIGINT DEFAULT 1,
    current_number          BIGINT DEFAULT 1,
    reset_type              VARCHAR(20) DEFAULT 'daily',
    created_at              TIMESTAMPTZ DEFAULT now(),
    updated_at              TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_restaurant_order_numbering_settings_id ON hospitality.restaurant_order_numbering_settings(restaurant_settings_id);

-- --------------------------------------
-- hospitality.restaurant_kitchen_settings
-- --------------------------------------
CREATE TABLE hospitality.restaurant_kitchen_settings (
    id                      UUID PRIMARY KEY,
    restaurant_settings_id  UUID NOT NULL,
    auto_accept_orders      BOOLEAN DEFAULT true,
    auto_print              BOOLEAN DEFAULT false,
    group_by_course         BOOLEAN DEFAULT false,
    created_at              TIMESTAMPTZ DEFAULT now(),
    updated_at              TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_restaurant_kitchen_settings_id ON hospitality.restaurant_kitchen_settings(restaurant_settings_id);

-- --------------------------------------
-- hospitality.restaurant_pos_settings
-- --------------------------------------
CREATE TABLE hospitality.restaurant_pos_settings (
    id                      UUID PRIMARY KEY,
    restaurant_settings_id  UUID NOT NULL,
    enable_table_selection  BOOLEAN DEFAULT true,
    enable_split_bill       BOOLEAN DEFAULT true,
    enable_discount         BOOLEAN DEFAULT true,
    require_customer_info   BOOLEAN DEFAULT false,
    created_at              TIMESTAMPTZ DEFAULT now(),
    updated_at              TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_restaurant_pos_settings_id ON hospitality.restaurant_pos_settings(restaurant_settings_id);

-- --------------------------------------
-- hospitality.restaurant_tax_settings
-- --------------------------------------
CREATE TABLE hospitality.restaurant_tax_settings (
    id                      UUID PRIMARY KEY,
    restaurant_settings_id  UUID NOT NULL,
    default_tax_rate        NUMERIC(5,2) DEFAULT 0,
    is_tax_inclusive        BOOLEAN DEFAULT false,
    created_at              TIMESTAMPTZ DEFAULT now(),
    updated_at              TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_restaurant_tax_settings_id ON hospitality.restaurant_tax_settings(restaurant_settings_id);

-- --------------------------------------
-- hospitality.restaurant_invoice_settings
-- --------------------------------------
CREATE TABLE hospitality.restaurant_invoice_settings (
    id                      UUID PRIMARY KEY,
    restaurant_settings_id  UUID NOT NULL,
    invoice_prefix          VARCHAR(20) DEFAULT 'INV',
    show_logo               BOOLEAN DEFAULT true,
    show_tax_breakup        BOOLEAN DEFAULT true,
    footer_note             TEXT,
    created_at              TIMESTAMPTZ DEFAULT now(),
    updated_at              TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_restaurant_invoice_settings_id ON hospitality.restaurant_invoice_settings(restaurant_settings_id);

-- --------------------------------------
-- hospitality.restaurant_prefered_kitchen
-- --------------------------------------
CREATE TABLE hospitality.restaurant_prefered_kitchen (
    id            UUID PRIMARY KEY,
    restaurant_id UUID NOT NULL,
    kitchen_id    UUID NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT now(),
    updated_at    TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_restaurant_prefered_kitchen_restaurant_id ON hospitality.restaurant_prefered_kitchen(restaurant_id);

-- --------------------------------------
-- hospitality.restaurant_configuration
-- --------------------------------------
CREATE TABLE hospitality.restaurant_configuration (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    restaurant_id   UUID NOT NULL,
    is_active       BOOLEAN DEFAULT false NOT NULL,
    is_configured   BOOLEAN DEFAULT false NOT NULL,
    current_status  VARCHAR(50) NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_restaurant_configuration_organization_id ON hospitality.restaurant_configuration(organization_id);
CREATE UNIQUE INDEX idx_restaurant_configuration_restaurant_id ON hospitality.restaurant_configuration(restaurant_id);

-- --------------------------------------
-- hospitality.restaurant_dining_sessions
-- --------------------------------------
CREATE TABLE hospitality.restaurant_dining_sessions (
    id            UUID PRIMARY KEY,
    restaurant_id UUID NOT NULL,
    table_id      UUID NOT NULL,
    status        VARCHAR(20) NOT NULL,
    started_at    TIMESTAMPTZ,
    started_by    UUID,
    closed_at     TIMESTAMPTZ,
    cancelled_at  TIMESTAMPTZ,
    created_at    TIMESTAMPTZ DEFAULT now(),
    updated_at    TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_restaurant_dining_sessions_restaurant_id ON hospitality.restaurant_dining_sessions(restaurant_id);
CREATE INDEX idx_restaurant_dining_sessions_table_id ON hospitality.restaurant_dining_sessions(table_id);

-- --------------------------------------
-- hospitality.restaurant_dining_session_customers
-- --------------------------------------
CREATE TABLE hospitality.restaurant_dining_session_customers (
    id                           UUID PRIMARY KEY,
    restaurant_dining_session_id UUID NOT NULL,
    customer_id                  UUID,
    customer_name                VARCHAR(200),
    customer_phone               VARCHAR(20),
    customer_gst                 VARCHAR(20),
    customer_state_code          VARCHAR(3),
    created_at                   TIMESTAMPTZ DEFAULT now(),
    updated_at                   TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_dining_session_customers_session_id ON hospitality.restaurant_dining_session_customers(restaurant_dining_session_id);
CREATE INDEX idx_dining_session_customers_customer_id ON hospitality.restaurant_dining_session_customers(customer_id);

-- --------------------------------------
-- hospitality.restaurant_dining_session_payments
-- --------------------------------------
CREATE TABLE hospitality.restaurant_dining_session_payments (
    id                           UUID PRIMARY KEY,
    restaurant_dining_session_id UUID NOT NULL,
    payment_method               VARCHAR(30) NOT NULL,
    invoice_id                   UUID,
    payment_order_id             UUID,
    payment_id                   UUID,
    qr_data                      TEXT,
    status                       VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    confirmed_at                 TIMESTAMPTZ,
    created_at                   TIMESTAMPTZ DEFAULT now(),
    updated_at                   TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_dining_session_payments_session_id ON hospitality.restaurant_dining_session_payments(restaurant_dining_session_id);
CREATE INDEX idx_dining_session_payments_invoice_id ON hospitality.restaurant_dining_session_payments(invoice_id);
CREATE INDEX idx_dining_session_payments_payment_order_id ON hospitality.restaurant_dining_session_payments(payment_order_id);
CREATE INDEX idx_dining_session_payments_payment_id ON hospitality.restaurant_dining_session_payments(payment_id);

-- --------------------------------------
-- hospitality.restaurant_orders
-- --------------------------------------
CREATE TABLE hospitality.restaurant_orders (
    id                           UUID PRIMARY KEY,
    organization_id              UUID NOT NULL,
    restaurant_id                UUID NOT NULL,
    legal_entity_id              UUID NOT NULL,
    restaurant_dining_session_id UUID,
    restaurant_pick_up_orders_id UUID,
    order_type                   VARCHAR(50),
    order_number                 VARCHAR(50) NOT NULL,
    status                       VARCHAR(30) NOT NULL,
    created_at                   TIMESTAMPTZ DEFAULT now(),
    updated_at                   TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_restaurant_orders_organization_id ON hospitality.restaurant_orders(organization_id);
CREATE INDEX idx_restaurant_orders_restaurant_id ON hospitality.restaurant_orders(restaurant_id);
CREATE INDEX idx_restaurant_orders_legal_entity_id ON hospitality.restaurant_orders(legal_entity_id);
CREATE INDEX idx_restaurant_orders_dining_session_id ON hospitality.restaurant_orders(restaurant_dining_session_id);
CREATE INDEX idx_restaurant_orders_pickup_order_id ON hospitality.restaurant_orders(restaurant_pick_up_orders_id);
CREATE INDEX idx_restaurant_orders_order_number ON hospitality.restaurant_orders(order_number);
CREATE INDEX idx_restaurant_orders_status ON hospitality.restaurant_orders(status);

-- --------------------------------------
-- hospitality.restaurant_order_items
-- --------------------------------------
CREATE TABLE hospitality.restaurant_order_items (
    id                  UUID PRIMARY KEY,
    restaurant_order_id UUID NOT NULL,
    menu_item_id        UUID NOT NULL,
    menu_item_name      VARCHAR(100),
    quantity            INT NOT NULL,
    unit_price          NUMERIC(12,2) NOT NULL,
    status              VARCHAR(30) NOT NULL,
    notes               TEXT,
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_restaurant_order_items_order_id ON hospitality.restaurant_order_items(restaurant_order_id);
CREATE INDEX idx_restaurant_order_items_menu_item_id ON hospitality.restaurant_order_items(menu_item_id);
CREATE INDEX idx_restaurant_order_items_status ON hospitality.restaurant_order_items(status);

-- --------------------------------------
-- hospitality.restaurant_order_items_variants
-- --------------------------------------
CREATE TABLE hospitality.restaurant_order_items_variants (
    id                      UUID PRIMARY KEY,
    restaurant_order_item_id UUID NOT NULL,
    variant_id              UUID NOT NULL,
    variant_group_id        UUID NOT NULL,
    name                    VARCHAR(100) NOT NULL,
    price_at_order          NUMERIC(12,2) NOT NULL,
    created_at              TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_restaurant_order_items_variants_order_item_id ON hospitality.restaurant_order_items_variants(restaurant_order_item_id);

-- --------------------------------------
-- hospitality.restaurant_order_items_addons
-- --------------------------------------
CREATE TABLE hospitality.restaurant_order_items_addons (
    id                      UUID PRIMARY KEY,
    restaurant_order_item_id UUID NOT NULL,
    add_on_id               UUID NOT NULL,
    name                    VARCHAR(100) NOT NULL,
    quantity                INT NOT NULL DEFAULT 1,
    price_at_order          NUMERIC(12,2) NOT NULL,
    created_at              TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_restaurant_order_items_addons_order_item_id ON hospitality.restaurant_order_items_addons(restaurant_order_item_id);

-- --------------------------------------
-- hospitality.restaurant_pickup_orders
-- --------------------------------------
CREATE TABLE hospitality.restaurant_pickup_orders (
    id                UUID PRIMARY KEY,
    organization_id   UUID NOT NULL,
    restaurant_id     UUID NOT NULL,
    legal_entity_id   UUID NOT NULL,
    order_number      VARCHAR(50) NOT NULL,
    type              VARCHAR(20) NOT NULL DEFAULT 'TAKEAWAY',
    status            VARCHAR(30) NOT NULL DEFAULT 'PENDING_PAYMENT',
    is_paid           BOOLEAN DEFAULT false,
    confirmed_at      TIMESTAMPTZ,
    prepared_at       TIMESTAMPTZ,
    ready_at          TIMESTAMPTZ,
    completed_at      TIMESTAMPTZ,
    cancelled_at      TIMESTAMPTZ,
    created_by        VARCHAR(100),
    created_by_empl_id UUID,
    created_at        TIMESTAMPTZ DEFAULT now(),
    updated_at        TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_restaurant_pickup_orders_organization_id ON hospitality.restaurant_pickup_orders(organization_id);
CREATE INDEX idx_restaurant_pickup_orders_restaurant_id ON hospitality.restaurant_pickup_orders(restaurant_id);
CREATE INDEX idx_restaurant_pickup_orders_legal_entity_id ON hospitality.restaurant_pickup_orders(legal_entity_id);
CREATE UNIQUE INDEX idx_restaurant_pickup_orders_order_number ON hospitality.restaurant_pickup_orders(order_number);
CREATE INDEX idx_restaurant_pickup_orders_status ON hospitality.restaurant_pickup_orders(status);

-- --------------------------------------
-- hospitality.restaurant_pickup_order_payments
-- --------------------------------------
CREATE TABLE hospitality.restaurant_pickup_order_payments (
    id                          UUID PRIMARY KEY,
    restaurant_pick_up_order_id UUID NOT NULL,
    payment_method              VARCHAR(30) NOT NULL,
    invoice_id                  UUID,
    payment_order_id            UUID,
    payment_id                  UUID,
    qr_data                     TEXT,
    status                      VARCHAR(30) NOT NULL DEFAULT 'PENDING',
    created_at                  TIMESTAMPTZ DEFAULT now(),
    confirmed_at                TIMESTAMPTZ
);

CREATE UNIQUE INDEX idx_pickup_order_payments_order_id ON hospitality.restaurant_pickup_order_payments(restaurant_pick_up_order_id);
CREATE INDEX idx_pickup_order_payments_invoice_id ON hospitality.restaurant_pickup_order_payments(invoice_id);
CREATE INDEX idx_pickup_order_payments_payment_order_id ON hospitality.restaurant_pickup_order_payments(payment_order_id);
CREATE INDEX idx_pickup_order_payments_payment_id ON hospitality.restaurant_pickup_order_payments(payment_id);

-- --------------------------------------
-- hospitality.restaurant_pickup_order_customers
-- --------------------------------------
CREATE TABLE hospitality.restaurant_pickup_order_customers (
    id                          UUID PRIMARY KEY,
    restaurant_pick_up_order_id UUID NOT NULL,
    customer_id                 UUID,
    customer_name               VARCHAR(200) NOT NULL,
    customer_phone              VARCHAR(20) NOT NULL,
    customer_email              VARCHAR(200),
    delivery_address            TEXT,
    delivery_note               VARCHAR(500),
    delivery_latitude           DECIMAL(10,8),
    delivery_longitude          DECIMAL(11,8),
    created_at                  TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_pickup_order_customers_order_id ON hospitality.restaurant_pickup_order_customers(restaurant_pick_up_order_id);
CREATE INDEX idx_pickup_order_customers_customer_id ON hospitality.restaurant_pickup_order_customers(customer_id);