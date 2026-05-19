-- --------------------------------------
-- core.org_suppliers
-- --------------------------------------
CREATE TABLE core.org_suppliers (
    id                UUID PRIMARY KEY,
    organization_id   UUID,
    name              VARCHAR(255) NOT NULL,
    gst_number        VARCHAR(50),
    pan_number        VARCHAR(50),
    phone             VARCHAR(20),
    email             VARCHAR(255),
    address_line      VARCHAR(255),
    city              VARCHAR(100),
    state             VARCHAR(100),
    pincode           VARCHAR(20),
    country           VARCHAR(100),
    created_by_engine VARCHAR(20),
    created_at        TIMESTAMPTZ DEFAULT now(),
    updated_at        TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_org_suppliers_organization_id ON core.org_suppliers(organization_id);

-- --------------------------------------
-- core.org_supplier_legal_entity_links
-- --------------------------------------
CREATE TABLE core.org_supplier_legal_entity_links (
    id                  UUID PRIMARY KEY,
    supplier_id         UUID,
    legal_entity_id     UUID,
    payable_account_id  UUID,
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_org_supplier_links_supplier_id ON core.org_supplier_legal_entity_links(supplier_id);
CREATE INDEX idx_org_supplier_links_legal_entity_id ON core.org_supplier_legal_entity_links(legal_entity_id);