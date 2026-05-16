-- --------------------------------------
-- invoice_engine.invoice
-- --------------------------------------
CREATE TABLE invoice_engine.invoice (
    id                  UUID PRIMARY KEY,
    organization_id     UUID NOT NULL,
    legal_entity_id     UUID,
    invoice_number      VARCHAR(50) NOT NULL,
    invoice_date        DATE NOT NULL,
    due_date            DATE,
    customer_id         UUID,
    customer_name       VARCHAR(200) NOT NULL,
    customer_gst        VARCHAR(20),
    customer_state_code VARCHAR(3),
    legal_entity_gstin  VARCHAR(20),
    sub_total           DECIMAL(15,2) NOT NULL,
    tax_total           DECIMAL(15,2) NOT NULL,
    total               DECIMAL(15,2) NOT NULL,
    paid_paise          BIGINT DEFAULT 0,
    status              VARCHAR(20) NOT NULL DEFAULT 'DRAFT',
    payment_status      VARCHAR(20) NOT NULL DEFAULT 'UNPAID',
    source              VARCHAR(50),
    source_ref_id       UUID,
    created_by          VARCHAR(100),
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now(),
    deleted_at          TIMESTAMPTZ
);

CREATE INDEX idx_invoice_organization_id ON invoice_engine.invoice(organization_id);
CREATE INDEX idx_invoice_legal_entity_id ON invoice_engine.invoice(legal_entity_id);
CREATE INDEX idx_invoice_customer_id ON invoice_engine.invoice(customer_id);
CREATE INDEX idx_invoice_status ON invoice_engine.invoice(status);
CREATE INDEX idx_invoice_payment_status ON invoice_engine.invoice(payment_status);
CREATE INDEX idx_invoice_source ON invoice_engine.invoice(source);
CREATE INDEX idx_invoice_source_ref_id ON invoice_engine.invoice(source_ref_id);
CREATE INDEX idx_invoice_deleted_at ON invoice_engine.invoice(deleted_at);
CREATE UNIQUE INDEX idx_inv_num_org ON invoice_engine.invoice(invoice_number, organization_id);

-- --------------------------------------
-- invoice_engine.invoice_item
-- --------------------------------------
CREATE TABLE invoice_engine.invoice_item (
    id              UUID PRIMARY KEY,
    invoice_id      UUID NOT NULL,
    description     TEXT NOT NULL,
    quantity        DECIMAL(12,2) NOT NULL DEFAULT 1,
    rate            DECIMAL(15,2) NOT NULL,
    amount          DECIMAL(15,2) NOT NULL,
    tax_category_id UUID,
    hsn_code        VARCHAR(10),
    tax_rate        DECIMAL(5,2),
    tax_amount      DECIMAL(15,2),
    total_amount    DECIMAL(15,2) NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_invoice_item_invoice_id ON invoice_engine.invoice_item(invoice_id);
CREATE INDEX idx_invoice_item_tax_category_id ON invoice_engine.invoice_item(tax_category_id);

-- --------------------------------------
-- invoice_engine.invoice_tax_breakup
-- --------------------------------------
CREATE TABLE invoice_engine.invoice_tax_breakup (
    id         UUID PRIMARY KEY,
    invoice_id UUID NOT NULL,
    tax_type   VARCHAR(20) NOT NULL,
    rate       DECIMAL(5,2) NOT NULL,
    amount     DECIMAL(15,2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_invoice_tax_breakup_invoice_id ON invoice_engine.invoice_tax_breakup(invoice_id);

-- --------------------------------------
-- invoice_engine.invoice_payment_link
-- --------------------------------------
CREATE TABLE invoice_engine.invoice_payment_link (
    id               UUID PRIMARY KEY,
    invoice_id       UUID NOT NULL,
    payment_ref      VARCHAR(100) NOT NULL,
    payment_method   VARCHAR(30) NOT NULL,
    currency_code    CHAR(3) NOT NULL DEFAULT 'INR',
    allocated_paise  BIGINT NOT NULL,
    status           VARCHAR(20) NOT NULL DEFAULT 'APPLIED',
    paid_at          TIMESTAMPTZ NOT NULL,
    created_at       TIMESTAMPTZ DEFAULT now(),
    reversed_at      TIMESTAMPTZ
);

CREATE INDEX idx_invoice_payment_link_invoice_id ON invoice_engine.invoice_payment_link(invoice_id);
CREATE INDEX idx_invoice_payment_link_payment_ref ON invoice_engine.invoice_payment_link(payment_ref);
CREATE INDEX idx_invoice_payment_link_status ON invoice_engine.invoice_payment_link(status);