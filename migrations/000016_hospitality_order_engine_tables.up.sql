-- --------------------------------------
-- hospitality.hospitality_order
-- --------------------------------------
CREATE TABLE hospitality.hospitality_order (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    property_id     UUID NOT NULL,
    owners_id       UUID NOT NULL,
    module          VARCHAR(50),
    status          VARCHAR(20) DEFAULT 'CREATED',
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hospitality_order_organization_id ON hospitality.hospitality_order(organization_id);
CREATE INDEX idx_hospitality_order_property_id ON hospitality.hospitality_order(property_id);
CREATE INDEX idx_hospitality_order_owners_id ON hospitality.hospitality_order(owners_id);
CREATE INDEX idx_hospitality_order_module ON hospitality.hospitality_order(module);
CREATE INDEX idx_hospitality_order_status ON hospitality.hospitality_order(status);

-- --------------------------------------
-- hospitality.hospitality_order_item
-- --------------------------------------
CREATE TABLE hospitality.hospitality_order_item (
    id           UUID PRIMARY KEY,
    order_id     UUID NOT NULL,
    menu_item_id UUID NOT NULL,
    base_price   DECIMAL(12,2) NOT NULL,
    quantity     INT NOT NULL DEFAULT 1,
    notes        TEXT,
    status       VARCHAR(20) DEFAULT 'PENDING',
    created_at   TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hospitality_order_item_order_id ON hospitality.hospitality_order_item(order_id);
CREATE INDEX idx_hospitality_order_item_status ON hospitality.hospitality_order_item(status);

-- --------------------------------------
-- hospitality.hospitality_order_item_variant
-- --------------------------------------
CREATE TABLE hospitality.hospitality_order_item_variant (
    id               UUID PRIMARY KEY,
    order_item_id    UUID NOT NULL,
    menu_variant_id  UUID NOT NULL,
    price            DECIMAL(12,2) NOT NULL,
    created_at       TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hospitality_order_item_variant_order_item_id ON hospitality.hospitality_order_item_variant(order_item_id);

-- --------------------------------------
-- hospitality.hospitality_order_item_addon
-- --------------------------------------
CREATE TABLE hospitality.hospitality_order_item_addon (
    id            UUID PRIMARY KEY,
    order_item_id UUID NOT NULL,
    addon_id      UUID NOT NULL,
    price         DECIMAL(12,2) NOT NULL,
    quantity      INT NOT NULL DEFAULT 1,
    created_at    TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hospitality_order_item_addon_order_item_id ON hospitality.hospitality_order_item_addon(order_item_id);

-- --------------------------------------
-- hospitality.order_kitchen_routing
-- --------------------------------------
CREATE TABLE hospitality.order_kitchen_routing (
    id                UUID PRIMARY KEY,
    order_item_id     UUID,
    kitchen_id        UUID,
    kitchen_ticket_id UUID,
    status            VARCHAR(20),
    created_at        TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_order_kitchen_routing_order_item_id ON hospitality.order_kitchen_routing(order_item_id);
CREATE INDEX idx_order_kitchen_routing_kitchen_id ON hospitality.order_kitchen_routing(kitchen_id);