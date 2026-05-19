-- --------------------------------------
-- payments.payment_accounts
-- --------------------------------------
CREATE TABLE payments.payment_accounts (
    id              UUID PRIMARY KEY,
    tenant_id       UUID NOT NULL,
    account_ref     VARCHAR(100) NOT NULL,
    legal_entity_id UUID NOT NULL,
    currency_code   CHAR(3) NOT NULL DEFAULT 'INR',
    is_active       BOOLEAN DEFAULT true,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_payment_accounts_tenant_id ON payments.payment_accounts(tenant_id);
CREATE INDEX idx_payment_accounts_legal_entity_id ON payments.payment_accounts(legal_entity_id);

-- --------------------------------------
-- payments.gateway_configs
-- --------------------------------------
CREATE TABLE payments.gateway_configs (
    id                 UUID PRIMARY KEY,
    tenant_id          UUID NOT NULL,
    legal_entity_id    UUID NOT NULL,
    payment_account_id UUID NOT NULL,
    provider_code      VARCHAR(30) NOT NULL,
    display_name       VARCHAR(100) NOT NULL,
    config_encrypted   TEXT NOT NULL,
    is_test_mode       BOOLEAN DEFAULT false,
    is_active          BOOLEAN DEFAULT true,
    created_at         TIMESTAMPTZ DEFAULT now(),
    updated_at         TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_gateway_configs_tenant_id ON payments.gateway_configs(tenant_id);
CREATE INDEX idx_gateway_configs_legal_entity_id ON payments.gateway_configs(legal_entity_id);
CREATE INDEX idx_gateway_configs_payment_account_id ON payments.gateway_configs(payment_account_id);

-- --------------------------------------
-- payments.payment_orders
-- --------------------------------------
CREATE TABLE payments.payment_orders (
    id                 UUID PRIMARY KEY,
    tenant_id          UUID NOT NULL,
    legal_entity_id    UUID NOT NULL,
    payment_account_id UUID NOT NULL,
    order_number       VARCHAR(50) NOT NULL,
    payer_type         VARCHAR(30) NOT NULL,
    payer_ref          VARCHAR(100),
    payer_name         VARCHAR(200) NOT NULL,
    currency_code      CHAR(3) NOT NULL DEFAULT 'INR',
    amount_paise       BIGINT NOT NULL,
    paid_paise         BIGINT DEFAULT 0,
    status             VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    source             VARCHAR(50),
    source_ref_id      VARCHAR(100),
    source_meta        TEXT,
    expires_at         TIMESTAMPTZ,
    created_by         VARCHAR(100) NOT NULL,
    created_at         TIMESTAMPTZ DEFAULT now(),
    updated_at         TIMESTAMPTZ DEFAULT now(),
    deleted_at         TIMESTAMPTZ
);

CREATE INDEX idx_payment_orders_tenant_id ON payments.payment_orders(tenant_id);
CREATE INDEX idx_payment_orders_legal_entity_id ON payments.payment_orders(legal_entity_id);
CREATE INDEX idx_payment_orders_payment_account_id ON payments.payment_orders(payment_account_id);
CREATE INDEX idx_payment_orders_status ON payments.payment_orders(status);
CREATE INDEX idx_payment_orders_deleted_at ON payments.payment_orders(deleted_at);
CREATE UNIQUE INDEX idx_order_number_tenant ON payments.payment_orders(order_number, tenant_id);

-- --------------------------------------
-- payments.payments
-- --------------------------------------
CREATE TABLE payments.payments (
    id                 UUID PRIMARY KEY,
    tenant_id          UUID NOT NULL,
    legal_entity_id    UUID NOT NULL,
    payment_account_id UUID NOT NULL,
    payment_order_id   UUID NOT NULL,
    payment_number     VARCHAR(50) NOT NULL,
    currency_code      CHAR(3) NOT NULL DEFAULT 'INR',
    amount_paise       BIGINT NOT NULL,
    method             VARCHAR(30) NOT NULL,
    status             VARCHAR(30) NOT NULL DEFAULT 'INITIATED',
    gateway_config_id  UUID,
    gateway_order_ref  VARCHAR(200),
    gateway_txn_ref    VARCHAR(200),
    gateway_response   TEXT,
    reference_number   VARCHAR(100),
    paid_at            TIMESTAMPTZ,
    refunded_paise     BIGINT DEFAULT 0,
    notes              TEXT,
    created_by         VARCHAR(100) NOT NULL,
    created_at         TIMESTAMPTZ DEFAULT now(),
    updated_at         TIMESTAMPTZ DEFAULT now(),
    deleted_at         TIMESTAMPTZ
);

CREATE INDEX idx_payments_tenant_id ON payments.payments(tenant_id);
CREATE INDEX idx_payments_legal_entity_id ON payments.payments(legal_entity_id);
CREATE INDEX idx_payments_payment_account_id ON payments.payments(payment_account_id);
CREATE INDEX idx_payments_payment_order_id ON payments.payments(payment_order_id);
CREATE INDEX idx_payments_method ON payments.payments(method);
CREATE INDEX idx_payments_status ON payments.payments(status);
CREATE INDEX idx_payments_gateway_config_id ON payments.payments(gateway_config_id);
CREATE INDEX idx_payments_gateway_txn_ref ON payments.payments(gateway_txn_ref);
CREATE INDEX idx_payments_deleted_at ON payments.payments(deleted_at);
CREATE UNIQUE INDEX idx_pay_number_tenant ON payments.payments(payment_number, tenant_id);

-- --------------------------------------
-- payments.payment_allocations
-- --------------------------------------
CREATE TABLE payments.payment_allocations (
    id               UUID PRIMARY KEY,
    tenant_id        UUID NOT NULL,
    payment_id       UUID NOT NULL,
    target_type      VARCHAR(30) NOT NULL,
    target_id        VARCHAR(100) NOT NULL,
    allocated_paise  BIGINT NOT NULL,
    status           VARCHAR(20) NOT NULL DEFAULT 'ACTIVE',
    created_by       VARCHAR(100) NOT NULL,
    created_at       TIMESTAMPTZ DEFAULT now(),
    reversed_at      TIMESTAMPTZ,
    is_reversed      BOOLEAN DEFAULT false,
    reversed_by      VARCHAR(100),
    reversal_note    TEXT
);

CREATE INDEX idx_payment_allocations_tenant_id ON payments.payment_allocations(tenant_id);
CREATE INDEX idx_payment_allocations_payment_id ON payments.payment_allocations(payment_id);
CREATE INDEX idx_payment_allocations_status ON payments.payment_allocations(status);

-- --------------------------------------
-- payments.refunds
-- --------------------------------------
CREATE TABLE payments.refunds (
    id                  UUID PRIMARY KEY,
    tenant_id           UUID NOT NULL,
    payment_id          UUID NOT NULL,
    refund_number       VARCHAR(50) NOT NULL,
    currency_code       CHAR(3) NOT NULL DEFAULT 'INR',
    amount_paise        BIGINT NOT NULL,
    status              VARCHAR(20) NOT NULL DEFAULT 'INITIATED',
    reason              TEXT NOT NULL,
    gateway_refund_ref  VARCHAR(200),
    gateway_response    TEXT,
    refunded_at         TIMESTAMPTZ,
    initiated_by        VARCHAR(100) NOT NULL,
    source_ref          VARCHAR(100),
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_refunds_tenant_id ON payments.refunds(tenant_id);
CREATE INDEX idx_refunds_payment_id ON payments.refunds(payment_id);
CREATE INDEX idx_refunds_status ON payments.refunds(status);
CREATE UNIQUE INDEX idx_refund_number_tenant ON payments.refunds(refund_number, tenant_id);

-- --------------------------------------
-- payments.webhook_endpoints
-- --------------------------------------
CREATE TABLE payments.webhook_endpoints (
    id                UUID PRIMARY KEY,
    tenant_id         UUID NOT NULL,
    url               TEXT NOT NULL,
    description       VARCHAR(200),
    signing_secret    TEXT NOT NULL,
    subscribed_events TEXT DEFAULT '[]',
    is_active         BOOLEAN DEFAULT true,
    created_at        TIMESTAMPTZ DEFAULT now(),
    updated_at        TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_webhook_endpoints_tenant_id ON payments.webhook_endpoints(tenant_id);

-- --------------------------------------
-- payments.webhook_delivery_logs
-- --------------------------------------
CREATE TABLE payments.webhook_delivery_logs (
    id               UUID PRIMARY KEY,
    tenant_id        UUID NOT NULL,
    endpoint_id      UUID NOT NULL,
    event_type       VARCHAR(50) NOT NULL,
    event_id         VARCHAR(100) NOT NULL,
    payload          TEXT NOT NULL,
    status           VARCHAR(20) NOT NULL DEFAULT 'PENDING',
    attempt_num      INT DEFAULT 1,
    response_code    INT,
    response_body    TEXT,
    response_latency INT,
    error_message    TEXT,
    attempted_at     TIMESTAMPTZ NOT NULL,
    next_retry_at    TIMESTAMPTZ
);

CREATE INDEX idx_webhook_delivery_logs_tenant_id ON payments.webhook_delivery_logs(tenant_id);
CREATE INDEX idx_webhook_delivery_logs_endpoint_id ON payments.webhook_delivery_logs(endpoint_id);
CREATE INDEX idx_webhook_delivery_logs_event_type ON payments.webhook_delivery_logs(event_type);
CREATE INDEX idx_webhook_delivery_logs_event_id ON payments.webhook_delivery_logs(event_id);
CREATE INDEX idx_webhook_delivery_logs_status ON payments.webhook_delivery_logs(status);
CREATE INDEX idx_webhook_delivery_logs_next_retry_at ON payments.webhook_delivery_logs(next_retry_at);