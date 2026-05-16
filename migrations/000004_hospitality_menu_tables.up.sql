-- --------------------------------------
-- hospitality.menu_group
-- --------------------------------------
CREATE TABLE hospitality.menu_group (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    property_id     UUID,
    name            VARCHAR(100) NOT NULL,
    description     VARCHAR(255),
    owner_module    VARCHAR(50),
    owner_id        UUID,
    is_active       BOOLEAN DEFAULT true,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_menu_group_organization_id ON hospitality.menu_group(organization_id);
CREATE INDEX idx_menu_group_property_id ON hospitality.menu_group(property_id);
CREATE INDEX idx_menu_group_owner_id ON hospitality.menu_group(owner_id);

-- --------------------------------------
-- hospitality.menu_category
-- --------------------------------------
CREATE TABLE hospitality.menu_category (
    id            UUID PRIMARY KEY,
    menu_group_id UUID NOT NULL,
    name          VARCHAR(100) NOT NULL,
    description   VARCHAR(255),
    is_active     BOOLEAN DEFAULT true,
    created_at    TIMESTAMPTZ DEFAULT now(),
    updated_at    TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_menu_category_menu_group_id ON hospitality.menu_category(menu_group_id);

-- --------------------------------------
-- hospitality.menu_subcategory
-- --------------------------------------
CREATE TABLE hospitality.menu_subcategory (
    id               UUID PRIMARY KEY,
    menu_category_id UUID NOT NULL,
    name             VARCHAR(100) NOT NULL,
    description      VARCHAR(255),
    is_active        BOOLEAN DEFAULT true,
    created_at       TIMESTAMPTZ DEFAULT now(),
    updated_at       TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_menu_subcategory_menu_category_id ON hospitality.menu_subcategory(menu_category_id);

-- --------------------------------------
-- hospitality.menu_item
-- --------------------------------------
CREATE TABLE hospitality.menu_item (
    id                  UUID PRIMARY KEY,
    menu_group_id       UUID,
    menu_subcategory_id UUID,
    name                VARCHAR(100) NOT NULL,
    description         VARCHAR(255),
    base_price          DECIMAL(12,2) NOT NULL,
    is_available        BOOLEAN DEFAULT true,
    unit                VARCHAR(20) DEFAULT 'plate',
    dietary_type        VARCHAR(50),
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_menu_item_menu_subcategory_id ON hospitality.menu_item(menu_subcategory_id);

-- --------------------------------------
-- hospitality.menu_variant_group
-- --------------------------------------
CREATE TABLE hospitality.menu_variant_group (
    id              UUID PRIMARY KEY,
    menu_item_id    UUID NOT NULL,
    name            VARCHAR(50) NOT NULL,
    is_multi_select BOOLEAN DEFAULT false,
    is_active       BOOLEAN DEFAULT true,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_menu_variant_group_menu_item_id ON hospitality.menu_variant_group(menu_item_id);

-- --------------------------------------
-- hospitality.menu_variant
-- --------------------------------------
CREATE TABLE hospitality.menu_variant (
    id                   UUID PRIMARY KEY,
    menu_variant_group_id UUID NOT NULL,
    menu_item_id         UUID NOT NULL,
    name                 VARCHAR(50) NOT NULL,
    price                DECIMAL(12,2) NOT NULL,
    is_active            BOOLEAN DEFAULT true,
    created_at           TIMESTAMPTZ DEFAULT now(),
    updated_at           TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_menu_variant_menu_variant_group_id ON hospitality.menu_variant(menu_variant_group_id);
CREATE INDEX idx_menu_variant_menu_item_id ON hospitality.menu_variant(menu_item_id);

-- --------------------------------------
-- hospitality.menu_item_addon
-- --------------------------------------
CREATE TABLE hospitality.menu_item_addon (
    id           UUID PRIMARY KEY,
    menu_item_id UUID NOT NULL,
    name         VARCHAR(100) NOT NULL,
    price        DECIMAL(12,2) NOT NULL,
    is_veg       BOOLEAN DEFAULT true,
    is_active    BOOLEAN DEFAULT true,
    created_at   TIMESTAMPTZ DEFAULT now(),
    updated_at   TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_menu_item_addon_menu_item_id ON hospitality.menu_item_addon(menu_item_id);

-- --------------------------------------
-- hospitality.menu_group_kitchen_register
-- --------------------------------------
CREATE TABLE hospitality.menu_group_kitchen_register (
    id           UUID PRIMARY KEY,
    menu_group_id UUID NOT NULL,
    kitchen_id   UUID NOT NULL,
    owners_id    UUID NOT NULL,
    created_at   TIMESTAMPTZ DEFAULT now(),
    updated_at   TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_menu_group_kitchen_register_menu_group_id ON hospitality.menu_group_kitchen_register(menu_group_id);
CREATE INDEX idx_menu_group_kitchen_register_owners_id ON hospitality.menu_group_kitchen_register(owners_id);

-- --------------------------------------
-- hospitality.menu_category_kitchen_register
-- --------------------------------------
CREATE TABLE hospitality.menu_category_kitchen_register (
    id              UUID PRIMARY KEY,
    menu_category_id UUID NOT NULL,
    kitchen_id      UUID NOT NULL,
    owners_id       UUID NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_menu_category_kitchen_register_menu_category_id ON hospitality.menu_category_kitchen_register(menu_category_id);
CREATE INDEX idx_menu_category_kitchen_register_owners_id ON hospitality.menu_category_kitchen_register(owners_id);

-- --------------------------------------
-- hospitality.menu_subcategory_kitchen_register
-- --------------------------------------
CREATE TABLE hospitality.menu_subcategory_kitchen_register (
    id                 UUID PRIMARY KEY,
    menu_subcategory_id UUID NOT NULL,
    kitchen_id         UUID NOT NULL,
    owners_id          UUID NOT NULL,
    created_at         TIMESTAMPTZ DEFAULT now(),
    updated_at         TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_menu_subcategory_kitchen_reg_id ON hospitality.menu_subcategory_kitchen_register(menu_subcategory_id);
CREATE INDEX idx_menu_subcategory_kitchen_register_owners_id ON hospitality.menu_subcategory_kitchen_register(owners_id);

-- --------------------------------------
-- hospitality.menu_item_kitchen_register
-- --------------------------------------
CREATE TABLE hospitality.menu_item_kitchen_register (
    id           UUID PRIMARY KEY,
    menu_item_id UUID NOT NULL,
    kitchen_id   UUID NOT NULL,
    owners_id    UUID NOT NULL,
    owner_type   VARCHAR(50) NOT NULL,
    created_at   TIMESTAMPTZ DEFAULT now(),
    updated_at   TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_menu_item_kitchen_register_menu_item_id ON hospitality.menu_item_kitchen_register(menu_item_id);
CREATE INDEX idx_menu_item_kitchen_register_owners_id ON hospitality.menu_item_kitchen_register(owners_id);

-- --------------------------------------
-- hospitality.menu_group_tax_category
-- --------------------------------------
CREATE TABLE hospitality.menu_group_tax_category (
    id              UUID PRIMARY KEY,
    menu_group_id   UUID,
    legal_entity_id UUID,
    tax_category_id UUID,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

-- --------------------------------------
-- hospitality.menu_category_tax_category
-- --------------------------------------
CREATE TABLE hospitality.menu_category_tax_category (
    id               UUID PRIMARY KEY,
    menu_category_id UUID,
    legal_entity_id  UUID,
    tax_category_id  UUID,
    created_at       TIMESTAMPTZ DEFAULT now(),
    updated_at       TIMESTAMPTZ DEFAULT now()
);

-- --------------------------------------
-- hospitality.menu_sub_category_tax_category
-- --------------------------------------
CREATE TABLE hospitality.menu_sub_category_tax_category (
    id                  UUID PRIMARY KEY,
    menu_subcategory_id UUID,
    legal_entity_id     UUID,
    tax_category_id     UUID,
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now()
);

-- --------------------------------------
-- hospitality.menu_item_tax_category
-- --------------------------------------
CREATE TABLE hospitality.menu_item_tax_category (
    id              UUID PRIMARY KEY,
    menu_item_id    UUID,
    legal_entity_id UUID,
    tax_category_id UUID,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);