-- --------------------------------------
-- hospitality.kitchen
-- --------------------------------------
CREATE TABLE hospitality.kitchen (
    id                        UUID PRIMARY KEY,
    organization_id           UUID NOT NULL,
    organization_property_id  UUID NOT NULL,
    name                      VARCHAR(100) NOT NULL,
    description               TEXT,
    type                      VARCHAR(50),
    is_active                 BOOLEAN DEFAULT false,
    created_at                TIMESTAMPTZ DEFAULT now(),
    updated_at                TIMESTAMPTZ DEFAULT now()
);

-- --------------------------------------
-- hospitality.kitchen_ticket
-- --------------------------------------
CREATE TABLE hospitality.kitchen_ticket (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    property_id     UUID,
    kitchen_id      UUID NOT NULL,
    status          VARCHAR(20) DEFAULT 'PENDING',
    ticket_no       VARCHAR(50),
    order_id        UUID,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now(),
    completed_at    TIMESTAMPTZ
);

CREATE INDEX idx_kitchen_ticket_kitchen_id ON hospitality.kitchen_ticket(kitchen_id);
CREATE INDEX idx_kitchen_ticket_order_id ON hospitality.kitchen_ticket(order_id);

-- --------------------------------------
-- hospitality.kitchen_ticket_item
-- --------------------------------------
CREATE TABLE hospitality.kitchen_ticket_item (
    id                UUID PRIMARY KEY,
    kitchen_ticket_id UUID NOT NULL,
    menu_item_id      UUID NOT NULL,
    menu_item_name    VARCHAR(120),
    quantity          INT NOT NULL,
    remarks           VARCHAR(255),
    is_printed        BOOLEAN DEFAULT false,
    created_at        TIMESTAMPTZ DEFAULT now(),
    updated_at        TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_kitchen_ticket_item_ticket_id ON hospitality.kitchen_ticket_item(kitchen_ticket_id);

-- --------------------------------------
-- hospitality.kitchen_ticket_item_variants
-- --------------------------------------
CREATE TABLE hospitality.kitchen_ticket_item_variants (
    id                    UUID PRIMARY KEY,
    kitchen_ticket_item_id UUID,
    variant_id            UUID,
    variant_name          VARCHAR(100),
    created_at            TIMESTAMPTZ DEFAULT now()
);

-- --------------------------------------
-- hospitality.kitchen_ticket_item_addon
-- --------------------------------------
CREATE TABLE hospitality.kitchen_ticket_item_addon (
    id                    UUID PRIMARY KEY,
    kitchen_ticket_item_id UUID NOT NULL,
    add_on_id             UUID NOT NULL,
    quantity              INT NOT NULL DEFAULT 0,
    created_at            TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_kitchen_ticket_item_addon_ticket_item_id ON hospitality.kitchen_ticket_item_addon(kitchen_ticket_item_id);

-- --------------------------------------
-- hospitality.kitchen_auto_accept_orders
-- --------------------------------------
CREATE TABLE hospitality.kitchen_auto_accept_orders (
    id                      UUID PRIMARY KEY,
    organization_id         UUID NOT NULL,
    property_id             UUID,
    kitchen_id              UUID NOT NULL,
    is_auto_accept_enabled  BOOLEAN DEFAULT false,
    created_at              TIMESTAMPTZ DEFAULT now(),
    updated_at              TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_kitchen_auto_accept_kitchen_id ON hospitality.kitchen_auto_accept_orders(kitchen_id);

-- --------------------------------------
-- hospitality.kitchen_printer_configs
-- --------------------------------------
CREATE TABLE hospitality.kitchen_printer_configs (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    property_id     UUID,
    kitchen_id      UUID NOT NULL,
    printer_name    VARCHAR(100),
    ip_address      VARCHAR(45),
    printer_type    VARCHAR(50),
    print_duplicate BOOLEAN DEFAULT false,
    is_enabled      BOOLEAN DEFAULT true,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_kitchen_printer_configs_kitchen_id ON hospitality.kitchen_printer_configs(kitchen_id);

-- --------------------------------------
-- hospitality.kitchen_throttling_configs
-- --------------------------------------
CREATE TABLE hospitality.kitchen_throttling_configs (
    id                    UUID PRIMARY KEY,
    organization_id       UUID NOT NULL,
    property_id           UUID,
    kitchen_id            UUID NOT NULL,
    max_orders_per_slot   INT DEFAULT 0,
    slot_duration_minutes INT DEFAULT 15,
    is_enabled            BOOLEAN DEFAULT false,
    created_at            TIMESTAMPTZ DEFAULT now(),
    updated_at            TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_kitchen_throttling_configs_kitchen_id ON hospitality.kitchen_throttling_configs(kitchen_id);

-- --------------------------------------
-- hospitality.kitchen_operating_hours
-- --------------------------------------
CREATE TABLE hospitality.kitchen_operating_hours (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    property_id     UUID,
    kitchen_id      UUID NOT NULL,
    day_of_week     INT NOT NULL,
    opening_time    VARCHAR(5),
    closing_time    VARCHAR(5),
    is_closed       BOOLEAN DEFAULT false,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_kitchen_operating_hours_kitchen_id ON hospitality.kitchen_operating_hours(kitchen_id);