-- --------------------------------------
-- customers.customer
-- --------------------------------------
CREATE TABLE customers.customer (
    id              UUID PRIMARY KEY,
    organization_id UUID,
    name            VARCHAR(255),
    email           VARCHAR(255),
    phone           VARCHAR(20),
    gstin           VARCHAR(50)
);

-- --------------------------------------
-- customers.customer_details
-- --------------------------------------
CREATE TABLE customers.customer_details (
    id              UUID PRIMARY KEY,
    customer_id     UUID,
    secondary_email VARCHAR(255),
    secondary_phone VARCHAR(20),
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_customer_details_customer_id ON customers.customer_details(customer_id);

-- --------------------------------------
-- customers.customer_org_meta
-- --------------------------------------
CREATE TABLE customers.customer_org_meta (
    id          UUID PRIMARY KEY,
    customer_id UUID,
    added_by    UUID,
    created_at  VARCHAR(255)
);

CREATE INDEX idx_customer_org_meta_customer_id ON customers.customer_org_meta(customer_id);