-- --------------------------------------
-- accounting_engine.accounting_book
-- --------------------------------------
CREATE TABLE accounting_engine.accounting_book (
    id              UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    legal_entity_id UUID NOT NULL,
    name            VARCHAR(255) NOT NULL,
    status          VARCHAR(50),
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_accounting_book_organization_id ON accounting_engine.accounting_book(organization_id);
CREATE INDEX idx_accounting_book_legal_entity_id ON accounting_engine.accounting_book(legal_entity_id);

-- --------------------------------------
-- accounting_engine.accounting_book_setting
-- --------------------------------------
CREATE TABLE accounting_engine.accounting_book_setting (
    id                UUID PRIMARY KEY,
    org_id            UUID,
    accounting_book_id UUID NOT NULL,
    currency          VARCHAR(10),
    rounding          VARCHAR(20),
    date_format       VARCHAR(20)
);

CREATE UNIQUE INDEX idx_accounting_book_setting_book_id ON accounting_engine.accounting_book_setting(accounting_book_id);

-- --------------------------------------
-- accounting_engine.fiscal_year
-- --------------------------------------
CREATE TABLE accounting_engine.fiscal_year (
    id                UUID PRIMARY KEY,
    accounting_book_id UUID NOT NULL,
    name              VARCHAR(100),
    start_date        TIMESTAMPTZ,
    end_date          TIMESTAMPTZ,
    status            VARCHAR(50),
    created_at        TIMESTAMPTZ DEFAULT now(),
    updated_at        TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_fiscal_year_accounting_book_id ON accounting_engine.fiscal_year(accounting_book_id);

-- --------------------------------------
-- accounting_engine.account_group
-- --------------------------------------
CREATE TABLE accounting_engine.account_group (
    id                UUID PRIMARY KEY,
    name              VARCHAR(255) NOT NULL,
    accounting_book_id UUID NOT NULL,
    type              VARCHAR(20) NOT NULL,
    created_at        TIMESTAMPTZ DEFAULT now(),
    updated_at        TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_account_group_accounting_book_id ON accounting_engine.account_group(accounting_book_id);

-- --------------------------------------
-- accounting_engine.account
-- --------------------------------------
CREATE TABLE accounting_engine.account (
    id               UUID PRIMARY KEY,
    account_group_id UUID,
    name             VARCHAR(255) NOT NULL,
    code             VARCHAR(50) NOT NULL,
    opening_balance  DECIMAL(15,2) DEFAULT 0,
    opening_is_debit BOOLEAN DEFAULT true,
    status           VARCHAR(20) DEFAULT 'ACTIVE',
    created_at       TIMESTAMPTZ DEFAULT now(),
    updated_at       TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_account_account_group_id ON accounting_engine.account(account_group_id);

-- --------------------------------------
-- accounting_engine.journal_entry
-- --------------------------------------
CREATE TABLE accounting_engine.journal_entry (
    id                UUID PRIMARY KEY,
    accounting_book_id UUID NOT NULL,
    fiscal_year_id    UUID NOT NULL,
    entry_date        TIMESTAMPTZ NOT NULL,
    reference         VARCHAR(255),
    narration         TEXT,
    entry_type        VARCHAR(50),
    created_at        TIMESTAMPTZ DEFAULT now(),
    updated_at        TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_journal_entry_accounting_book_id ON accounting_engine.journal_entry(accounting_book_id);
CREATE INDEX idx_journal_entry_fiscal_year_id ON accounting_engine.journal_entry(fiscal_year_id);

-- --------------------------------------
-- accounting_engine.journal_line
-- --------------------------------------
CREATE TABLE accounting_engine.journal_line (
    id               UUID PRIMARY KEY,
    journal_entry_id UUID NOT NULL,
    account_id       UUID NOT NULL,
    debit            DECIMAL(15,2) DEFAULT 0,
    credit           DECIMAL(15,2) DEFAULT 0,
    party_id         UUID,
    created_at       TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_journal_line_journal_entry_id ON accounting_engine.journal_line(journal_entry_id);
CREATE INDEX idx_journal_line_account_id ON accounting_engine.journal_line(account_id);

-- --------------------------------------
-- accounting_engine.expense
-- --------------------------------------
CREATE TABLE accounting_engine.expense (
    id                  UUID PRIMARY KEY,
    organization_id     UUID NOT NULL,
    accounting_book_id  UUID NOT NULL,
    fiscal_year_id      UUID NOT NULL,
    expense_date        TIMESTAMPTZ NOT NULL,
    expense_category_id UUID,
    category_label      VARCHAR(100),
    description         VARCHAR(255) NOT NULL,
    payee_name          VARCHAR(255),
    supplier_id         UUID,
    sub_total           DECIMAL(15,2) DEFAULT 0,
    tax_amount          DECIMAL(15,2) DEFAULT 0,
    total_amount        DECIMAL(15,2) DEFAULT 0,
    tax_category_id     UUID,
    debit_account_id    UUID NOT NULL,
    credit_account_id   UUID NOT NULL,
    journal_entry_id    UUID,
    payment_mode        VARCHAR(20),
    reference           VARCHAR(255),
    status              VARCHAR(20) DEFAULT 'DRAFT',
    attachment_url      TEXT,
    narration           TEXT,
    created_by          UUID,
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_expense_organization_id ON accounting_engine.expense(organization_id);
CREATE INDEX idx_expense_accounting_book_id ON accounting_engine.expense(accounting_book_id);
CREATE INDEX idx_expense_fiscal_year_id ON accounting_engine.expense(fiscal_year_id);
CREATE INDEX idx_expense_expense_category_id ON accounting_engine.expense(expense_category_id);
CREATE INDEX idx_expense_supplier_id ON accounting_engine.expense(supplier_id);

-- --------------------------------------
-- accounting_engine.purchase_invoice
-- --------------------------------------
CREATE TABLE accounting_engine.purchase_invoice (
    id                  UUID PRIMARY KEY,
    accounting_book_id  UUID NOT NULL,
    fiscal_year_id      UUID NOT NULL,
    supplier_id         UUID NOT NULL,
    purchase_order_id   UUID,
    invoice_number      VARCHAR(100),
    our_reference       VARCHAR(100),
    invoice_date        TIMESTAMPTZ NOT NULL,
    due_date            TIMESTAMPTZ,
    status              VARCHAR(20) DEFAULT 'DRAFT',
    payment_status      VARCHAR(20) DEFAULT 'UNPAID',
    payable_account_id  UUID NOT NULL,
    sub_total           DECIMAL(15,2) DEFAULT 0,
    tax_amount          DECIMAL(15,2) DEFAULT 0,
    total_amount        DECIMAL(15,2) DEFAULT 0,
    amount_paid         DECIMAL(15,2) DEFAULT 0,
    amount_due          DECIMAL(15,2) DEFAULT 0,
    journal_entry_id    UUID,
    narration           TEXT,
    created_by          UUID,
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_purchase_invoice_accounting_book_id ON accounting_engine.purchase_invoice(accounting_book_id);
CREATE INDEX idx_purchase_invoice_fiscal_year_id ON accounting_engine.purchase_invoice(fiscal_year_id);
CREATE INDEX idx_purchase_invoice_supplier_id ON accounting_engine.purchase_invoice(supplier_id);
CREATE INDEX idx_purchase_invoice_purchase_order_id ON accounting_engine.purchase_invoice(purchase_order_id);

-- --------------------------------------
-- accounting_engine.purchase_invoice_line
-- --------------------------------------
CREATE TABLE accounting_engine.purchase_invoice_line (
    id                   UUID PRIMARY KEY,
    purchase_invoice_id  UUID NOT NULL,
    description          VARCHAR(255) NOT NULL,
    quantity             DECIMAL(15,4) DEFAULT 1,
    unit_price           DECIMAL(15,2) DEFAULT 0,
    debit_account_id     UUID NOT NULL,
    tax_category_id      UUID,
    line_amount          DECIMAL(15,2) DEFAULT 0,
    tax_amount           DECIMAL(15,2) DEFAULT 0,
    total_amount         DECIMAL(15,2) DEFAULT 0,
    inventory_item_id    UUID,
    inventory_id         UUID,
    inventory_account_id UUID,
    created_at           TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_purchase_invoice_line_invoice_id ON accounting_engine.purchase_invoice_line(purchase_invoice_id);

-- --------------------------------------
-- accounting_engine.purchase_payment
-- --------------------------------------
CREATE TABLE accounting_engine.purchase_payment (
    id                  UUID PRIMARY KEY,
    purchase_invoice_id UUID NOT NULL,
    supplier_id         UUID NOT NULL,
    payment_date        TIMESTAMPTZ NOT NULL,
    amount              DECIMAL(15,2) NOT NULL,
    payment_method      VARCHAR(20),
    pay_from_account_id UUID NOT NULL,
    journal_entry_id    UUID,
    reference           VARCHAR(255),
    notes               TEXT,
    created_by          UUID,
    created_at          TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_purchase_payment_purchase_invoice_id ON accounting_engine.purchase_payment(purchase_invoice_id);
CREATE INDEX idx_purchase_payment_supplier_id ON accounting_engine.purchase_payment(supplier_id);