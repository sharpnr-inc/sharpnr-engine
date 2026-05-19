-- --------------------------------------
-- hospitality.bar
-- --------------------------------------
CREATE TABLE hospitality.bar (
    id              UUID PRIMARY KEY,
    name            VARCHAR(100) NOT NULL,
    organization_id UUID NOT NULL,
    property_id     UUID NOT NULL,
    legal_entity_id UUID NOT NULL,
    is_active       BOOLEAN DEFAULT false,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_bar_property_id ON hospitality.bar(property_id);
CREATE INDEX idx_bar_legal_entity_id ON hospitality.bar(legal_entity_id);

-- --------------------------------------
-- hospitality.bar_floor
-- --------------------------------------
CREATE TABLE hospitality.bar_floor (
    id           UUID PRIMARY KEY,
    bar_id       UUID NOT NULL,
    property_id  UUID NOT NULL,
    name         VARCHAR(100),
    floor_number INT,
    created_at   TIMESTAMPTZ DEFAULT now(),
    updated_at   TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_bar_floor_bar_id ON hospitality.bar_floor(bar_id);
CREATE INDEX idx_bar_floor_property_id ON hospitality.bar_floor(property_id);

-- --------------------------------------
-- hospitality.bar_table
-- --------------------------------------
CREATE TABLE hospitality.bar_table (
    id          UUID PRIMARY KEY,
    bar_id      UUID NOT NULL,
    floor_id    UUID,
    property_id UUID NOT NULL,
    code        VARCHAR(50) NOT NULL,
    capacity    INT NOT NULL,
    status      VARCHAR(20),
    created_at  TIMESTAMPTZ DEFAULT now(),
    updated_at  TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_bar_table_bar_id ON hospitality.bar_table(bar_id);
CREATE INDEX idx_bar_table_property_id ON hospitality.bar_table(property_id);