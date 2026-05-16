-- --------------------------------------
-- hospitality.hotel
-- --------------------------------------
CREATE TABLE hospitality.hotel (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    property_id     UUID NOT NULL,
    legal_entity_id UUID NOT NULL,
    name            VARCHAR(100),
    is_active       BOOLEAN DEFAULT false,
    is_configured   BOOLEAN DEFAULT false,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_property_id ON hospitality.hotel(property_id);
CREATE INDEX idx_hotel_legal_entity_id ON hospitality.hotel(legal_entity_id);

-- --------------------------------------
-- hospitality.hotel_setup
-- --------------------------------------
CREATE TABLE hospitality.hotel_setup (
    id           UUID PRIMARY KEY,
    hotel_id     UUID NOT NULL,
    setup_status VARCHAR(50),
    setup_step   VARCHAR(100),
    created_at   TIMESTAMPTZ DEFAULT now(),
    updated_at   TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_setup_hotel_id ON hospitality.hotel_setup(hotel_id);

-- --------------------------------------
-- hospitality.hotel_operations_timing
-- --------------------------------------
CREATE TABLE hospitality.hotel_operations_timing (
    id              UUID PRIMARY KEY,
    hotel_id        UUID,
    is_24x7         BOOLEAN DEFAULT false,
    check_in_time   VARCHAR(10),
    check_out_time  VARCHAR(10)
);

-- --------------------------------------
-- hospitality.hotel_amenities
-- --------------------------------------
CREATE TABLE hospitality.hotel_amenities (
    id          UUID PRIMARY KEY,
    hotel_id    UUID,
    property_id UUID,
    name        VARCHAR(255),
    description TEXT,
    created_at  TIMESTAMPTZ DEFAULT now(),
    updated_at  TIMESTAMPTZ DEFAULT now()
);

-- --------------------------------------
-- hospitality.hotel_floor
-- --------------------------------------
CREATE TABLE hospitality.hotel_floor (
    id           UUID PRIMARY KEY,
    hotel_id     UUID NOT NULL,
    name         VARCHAR(100),
    floor_number INT NOT NULL,
    created_at   TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_floor_hotel_id ON hospitality.hotel_floor(hotel_id);

-- --------------------------------------
-- hospitality.hotel_room_type
-- --------------------------------------
CREATE TABLE hospitality.hotel_room_type (
    id               UUID PRIMARY KEY,
    hotel_id         UUID NOT NULL,
    code             VARCHAR(50) NOT NULL,
    name             VARCHAR(100),
    description      TEXT,
    base_occupancy   INT,
    max_occupancy    INT,
    is_active        BOOLEAN DEFAULT false,
    created_at       TIMESTAMPTZ DEFAULT now(),
    updated_at       TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_room_type_hotel_id ON hospitality.hotel_room_type(hotel_id);

-- --------------------------------------
-- hospitality.hotel_room_addons
-- --------------------------------------
CREATE TABLE hospitality.hotel_room_addons (
    id                       UUID PRIMARY KEY,
    hotel_id                 UUID,
    organization_property_id UUID,
    name                     VARCHAR(255),
    description              TEXT,
    code                     VARCHAR(50),
    is_complimentary         BOOLEAN DEFAULT false,
    created_at               TIMESTAMPTZ DEFAULT now(),
    updated_at               TIMESTAMPTZ DEFAULT now()
);

-- --------------------------------------
-- hospitality.hotel_room_feature
-- --------------------------------------
CREATE TABLE hospitality.hotel_room_feature (
    id                       UUID PRIMARY KEY,
    hotel_id                 UUID,
    organization_property_id UUID,
    code                     VARCHAR(50),
    name                     VARCHAR(255),
    description              TEXT,
    created_at               TIMESTAMPTZ DEFAULT now(),
    updated_at               TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_hotel_room_feature_code ON hospitality.hotel_room_feature(code);

-- --------------------------------------
-- hospitality.hotel_room_type_addon
-- --------------------------------------
CREATE TABLE hospitality.hotel_room_type_addon (
    id                  UUID PRIMARY KEY,
    hotel_id            UUID,
    hotel_room_addons_id UUID,
    hotel_room_type_id  UUID,
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now()
);

-- --------------------------------------
-- hospitality.hotel_room_type_feature
-- --------------------------------------
CREATE TABLE hospitality.hotel_room_type_feature (
    id                   UUID PRIMARY KEY,
    hotel_id             UUID,
    hotel_room_feature_id UUID,
    hotel_room_type_id   UUID,
    created_at           TIMESTAMPTZ DEFAULT now(),
    updated_at           TIMESTAMPTZ DEFAULT now()
);

-- --------------------------------------
-- hospitality.hotel_room_type_rate
-- --------------------------------------
CREATE TABLE hospitality.hotel_room_type_rate (
    id             UUID PRIMARY KEY,
    hotel_id       UUID NOT NULL,
    room_type_id   UUID NOT NULL,
    base_price     DECIMAL(12,2) NOT NULL,
    effective_from TIMESTAMPTZ NOT NULL,
    effective_to   TIMESTAMPTZ,
    created_at     TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_room_type_rate_hotel_id ON hospitality.hotel_room_type_rate(hotel_id);
CREATE INDEX idx_hotel_room_type_rate_room_type_id ON hospitality.hotel_room_type_rate(room_type_id);

-- --------------------------------------
-- hospitality.hotel_meta_data
-- --------------------------------------
CREATE TABLE hospitality.hotel_meta_data (
    id         UUID PRIMARY KEY,
    hotel_id   UUID,
    meta       JSONB,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- --------------------------------------
-- hospitality.hotel_room
-- --------------------------------------
CREATE TABLE hospitality.hotel_room (
    id                UUID PRIMARY KEY,
    hotel_id          UUID NOT NULL,
    hotel_floor_id    UUID NOT NULL,
    hotel_room_type_id UUID,
    room_code         VARCHAR(30),
    name              VARCHAR(100),
    status            VARCHAR(20) NOT NULL DEFAULT 'AVAILABLE',
    bed_type          VARCHAR(50),
    max_guests        INT,
    created_at        TIMESTAMPTZ DEFAULT now(),
    updated_at        TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_room_hotel_id ON hospitality.hotel_room(hotel_id);
CREATE INDEX idx_hotel_room_hotel_floor_id ON hospitality.hotel_room(hotel_floor_id);

-- --------------------------------------
-- hospitality.hotel_room_detail
-- --------------------------------------
CREATE TABLE hospitality.hotel_room_detail (
    id                UUID PRIMARY KEY,
    hotel_room_id     UUID NOT NULL,
    hotel_room_type_id UUID,
    description       TEXT,
    bed_type          VARCHAR(50),
    created_at        TIMESTAMPTZ DEFAULT now(),
    updated_at        TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_hotel_room_detail_hotel_room_id ON hospitality.hotel_room_detail(hotel_room_id);

-- --------------------------------------
-- hospitality.hotel_room_rate
-- --------------------------------------
CREATE TABLE hospitality.hotel_room_rate (
    id             UUID PRIMARY KEY,
    hotel_id       UUID NOT NULL,
    hotel_room_id  UUID NOT NULL,
    base_price     DECIMAL(12,2) NOT NULL,
    effective_from TIMESTAMPTZ,
    effective_to   TIMESTAMPTZ,
    created_at     TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_room_rate_hotel_id ON hospitality.hotel_room_rate(hotel_id);
CREATE INDEX idx_hotel_room_rate_hotel_room_id ON hospitality.hotel_room_rate(hotel_room_id);

-- --------------------------------------
-- hospitality.hotel_room_rate_history
-- --------------------------------------
CREATE TABLE hospitality.hotel_room_rate_history (
    id             UUID PRIMARY KEY,
    hotel_id       UUID NOT NULL,
    hotel_room_id  UUID NOT NULL,
    last_price     DECIMAL(12,2) NOT NULL,
    updated_price  DECIMAL(12,2) NOT NULL,
    reason         TEXT,
    executed_by    VARCHAR(100),
    created_at     TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_room_rate_history_hotel_id ON hospitality.hotel_room_rate_history(hotel_id);
CREATE INDEX idx_hotel_room_rate_history_hotel_room_id ON hospitality.hotel_room_rate_history(hotel_room_id);

-- --------------------------------------
-- hospitality.hotel_stay
-- --------------------------------------
CREATE TABLE hospitality.hotel_stay (
    id               UUID PRIMARY KEY,
    hotel_id         UUID NOT NULL,
    room_id          UUID NOT NULL,
    primary_guest_id UUID NOT NULL,
    status           VARCHAR(20) NOT NULL,
    check_in_date    TIMESTAMPTZ,
    check_out_date   TIMESTAMPTZ,
    created_at       TIMESTAMPTZ DEFAULT now(),
    updated_at       TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_stay_hotel_id ON hospitality.hotel_stay(hotel_id);
CREATE INDEX idx_hotel_stay_room_id ON hospitality.hotel_stay(room_id);

-- --------------------------------------
-- hospitality.hotel_charge
-- --------------------------------------
CREATE TABLE hospitality.hotel_charge (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    property_id     UUID NOT NULL,
    hotel_id        UUID NOT NULL,
    legal_entity_id UUID NOT NULL,
    stay_id         UUID NOT NULL,
    charge_type     VARCHAR(50),
    source_module   VARCHAR(50),
    source_ref_id   UUID,
    description     TEXT,
    quantity        DECIMAL(10,2) DEFAULT 1,
    unit_rate       DECIMAL(12,2),
    amount          DECIMAL(12,2),
    charge_date     TIMESTAMPTZ,
    created_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_charge_organization_id ON hospitality.hotel_charge(organization_id);
CREATE INDEX idx_hotel_charge_property_id ON hospitality.hotel_charge(property_id);
CREATE INDEX idx_hotel_charge_hotel_id ON hospitality.hotel_charge(hotel_id);
CREATE INDEX idx_hotel_charge_legal_entity_id ON hospitality.hotel_charge(legal_entity_id);
CREATE INDEX idx_hotel_charge_stay_id ON hospitality.hotel_charge(stay_id);
CREATE INDEX idx_hotel_charge_charge_type ON hospitality.hotel_charge(charge_type);
CREATE INDEX idx_hotel_charge_charge_date ON hospitality.hotel_charge(charge_date);
CREATE INDEX idx_hotel_charge_source_ref_id ON hospitality.hotel_charge(source_ref_id);

-- --------------------------------------
-- hospitality.hotel_stay_payment
-- --------------------------------------
CREATE TABLE hospitality.hotel_stay_payment (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    property_id     UUID NOT NULL,
    hotel_id        UUID NOT NULL,
    legal_entity_id UUID NOT NULL,
    stay_id         UUID NOT NULL,
    amount          DECIMAL(12,2) NOT NULL,
    method          VARCHAR(20) NOT NULL,
    status          VARCHAR(20) NOT NULL DEFAULT 'RECEIVED',
    reference       VARCHAR(100),
    received_at     TIMESTAMPTZ,
    created_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_stay_payment_organization_id ON hospitality.hotel_stay_payment(organization_id);
CREATE INDEX idx_hotel_stay_payment_property_id ON hospitality.hotel_stay_payment(property_id);
CREATE INDEX idx_hotel_stay_payment_hotel_id ON hospitality.hotel_stay_payment(hotel_id);
CREATE INDEX idx_hotel_stay_payment_legal_entity_id ON hospitality.hotel_stay_payment(legal_entity_id);
CREATE INDEX idx_hotel_stay_payment_stay_id ON hospitality.hotel_stay_payment(stay_id);
CREATE INDEX idx_hotel_stay_payment_status ON hospitality.hotel_stay_payment(status);
CREATE INDEX idx_hotel_stay_payment_received_at ON hospitality.hotel_stay_payment(received_at);

-- --------------------------------------
-- hospitality.hotel_guest
-- --------------------------------------
CREATE TABLE hospitality.hotel_guest (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    property_id     UUID NOT NULL,
    hotel_id        UUID NOT NULL,
    legal_entity_id UUID NOT NULL,
    name            VARCHAR(150) NOT NULL,
    mobile          VARCHAR(20),
    email           VARCHAR(150),
    gender          VARCHAR(10),
    dob             VARCHAR(20),
    aadhaar_id      VARCHAR(20),
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_guest_organization_id ON hospitality.hotel_guest(organization_id);
CREATE INDEX idx_hotel_guest_property_id ON hospitality.hotel_guest(property_id);
CREATE INDEX idx_hotel_guest_hotel_id ON hospitality.hotel_guest(hotel_id);
CREATE INDEX idx_hotel_guest_legal_entity_id ON hospitality.hotel_guest(legal_entity_id);
CREATE INDEX idx_hotel_guest_mobile ON hospitality.hotel_guest(mobile);
CREATE INDEX idx_hotel_guest_email ON hospitality.hotel_guest(email);
CREATE INDEX idx_hotel_guest_aadhaar_id ON hospitality.hotel_guest(aadhaar_id);

-- --------------------------------------
-- hospitality.hotel_stay_guest
-- --------------------------------------
CREATE TABLE hospitality.hotel_stay_guest (
    id              UUID PRIMARY KEY,
    hotel_stay_id   UUID NOT NULL,
    hotel_id        UUID NOT NULL,
    hotel_guest_id  UUID NOT NULL,
    name            VARCHAR(150),
    mobile          VARCHAR(20),
    email           VARCHAR(150),
    dob             VARCHAR(20),
    created_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_hotel_stay_guest_hotel_stay_id ON hospitality.hotel_stay_guest(hotel_stay_id);
CREATE INDEX idx_hotel_stay_guest_hotel_id ON hospitality.hotel_stay_guest(hotel_id);
CREATE INDEX idx_hotel_stay_guest_hotel_guest_id ON hospitality.hotel_stay_guest(hotel_guest_id);