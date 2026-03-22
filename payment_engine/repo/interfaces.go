package payment_engine_repo

import (
	"context"
	payment_engine_models "sharpnr-engine/payment_engine/models"
	"time"

	"github.com/google/uuid"
)

// -----------------------------------------------
// PRIMARY INTERFACE
// This is the ONLY thing business services import
// Hotel, Restaurant, Bar, Orchestrator all depend
// on this — never on the concrete impl
// -----------------------------------------------
type PaymentEngineRepoInterface interface {

	// ─── Payment Account ──────────────────────────────────────────────────────

	CreatePaymentAccount(ctx context.Context, account *payment_engine_models.PaymentAccount) error
	GetPaymentAccountByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.PaymentAccount, error)
	GetPaymentAccountByLegalEntity(ctx context.Context, tenantID, legalEntityID uuid.UUID) (*payment_engine_models.PaymentAccount, error)
	ListPaymentAccounts(ctx context.Context, tenantID uuid.UUID) ([]*payment_engine_models.PaymentAccount, error)
	UpdatePaymentAccountStatus(ctx context.Context, id uuid.UUID, isActive bool) error

	// ─── Gateway Config ───────────────────────────────────────────────────────

	CreateGatewayConfig(ctx context.Context, config *payment_engine_models.PaymentGatewayConfig) error
	GetGatewayConfigByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.PaymentGatewayConfig, error)
	GetActiveGatewayConfig(ctx context.Context, paymentAccountID uuid.UUID, providerCode string) (*payment_engine_models.PaymentGatewayConfig, error)
	ListGatewayConfigsByAccount(ctx context.Context, paymentAccountID uuid.UUID) ([]*payment_engine_models.PaymentGatewayConfig, error)
	UpdateGatewayConfig(ctx context.Context, id uuid.UUID, configEncrypted string) error
	ToggleGatewayConfig(ctx context.Context, id uuid.UUID, isActive bool) error
	DeleteGatewayConfig(ctx context.Context, id uuid.UUID) error

	// ─── Payment Order ────────────────────────────────────────────────────────

	CreateOrder(ctx context.Context, order *payment_engine_models.PaymentOrder) error
	GetOrderByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.PaymentOrder, error)
	GetOrderByNumber(ctx context.Context, tenantID uuid.UUID, orderNumber string) (*payment_engine_models.PaymentOrder, error)
	ListOrdersBySource(ctx context.Context, tenantID uuid.UUID, sourceType string, sourceRefID string) ([]*payment_engine_models.PaymentOrder, error)
	UpdateOrderStatus(ctx context.Context, id uuid.UUID, status string) error
	UpdateOrderPaidAmount(ctx context.Context, id uuid.UUID, paidPaise int64, newStatus string) error
	GetExpiredPendingOrders(ctx context.Context, before time.Time) ([]*payment_engine_models.PaymentOrder, error)

	// ─── Payment ──────────────────────────────────────────────────────────────

	CreatePayment(ctx context.Context, payment *payment_engine_models.Payment) error
	GetPaymentByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.Payment, error)
	GetPaymentByGatewayOrderRef(ctx context.Context, gatewayOrderRef string) (*payment_engine_models.Payment, error)
	GetPaymentByGatewayTxnRef(ctx context.Context, gatewayTxnRef string) (*payment_engine_models.Payment, error)
	GetPaymentsByOrderID(ctx context.Context, orderID uuid.UUID) ([]*payment_engine_models.Payment, error)
	UpdatePaymentStatus(ctx context.Context, id uuid.UUID, status string) error
	UpdatePaymentGatewayRefs(ctx context.Context, payment *payment_engine_models.Payment) error
	UpdatePaymentConfirmation(ctx context.Context, payment *payment_engine_models.Payment) error
	IncrementPaymentRefunded(ctx context.Context, id uuid.UUID, amountPaise int64) error

	// ─── Payment Allocation ───────────────────────────────────────────────────

	CreateAllocation(ctx context.Context, allocation *payment_engine_models.PaymentAllocation) error
	GetAllocationByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.PaymentAllocation, error)
	GetAllocationsByPaymentID(ctx context.Context, paymentID uuid.UUID) ([]*payment_engine_models.PaymentAllocation, error)
	GetAllocationsByTarget(ctx context.Context, targetType string, targetID string) ([]*payment_engine_models.PaymentAllocation, error)
	GetTotalAllocatedByPayment(ctx context.Context, paymentID uuid.UUID) (int64, error)
	GetTotalAllocatedByTarget(ctx context.Context, targetType string, targetID string) (int64, error)
	ReverseAllocation(ctx context.Context, id uuid.UUID, by string, note string) error

	// ─── Refund ───────────────────────────────────────────────────────────────

	CreateRefund(ctx context.Context, refund *payment_engine_models.Refund) error
	GetRefundByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.Refund, error)
	GetRefundsByPaymentID(ctx context.Context, paymentID uuid.UUID) ([]*payment_engine_models.Refund, error)
	UpdateRefundStatus(ctx context.Context, id uuid.UUID, status string) error
	UpdateRefundGatewayRefs(ctx context.Context, refund *payment_engine_models.Refund) error

	// ─── Webhook Endpoint ─────────────────────────────────────────────────────

	CreateWebhookEndpoint(ctx context.Context, endpoint *payment_engine_models.WebhookEndpoint) error
	GetWebhookEndpointByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.WebhookEndpoint, error)
	GetActiveWebhookEndpoints(ctx context.Context, tenantID uuid.UUID, eventType string) ([]*payment_engine_models.WebhookEndpoint, error)
	ListWebhookEndpoints(ctx context.Context, tenantID uuid.UUID) ([]*payment_engine_models.WebhookEndpoint, error)
	UpdateWebhookEndpoint(ctx context.Context, id uuid.UUID, url string, subscribedEvents string) error
	ToggleWebhookEndpoint(ctx context.Context, id uuid.UUID, isActive bool) error
	DeleteWebhookEndpoint(ctx context.Context, id uuid.UUID) error

	// ─── Webhook Delivery Log ─────────────────────────────────────────────────

	CreateWebhookDeliveryLog(ctx context.Context, log *payment_engine_models.WebhookDeliveryLog) error
	GetWebhookDeliveryLogByID(ctx context.Context, id uuid.UUID) (*payment_engine_models.WebhookDeliveryLog, error)
	GetPendingWebhookRetries(ctx context.Context, before time.Time) ([]*payment_engine_models.WebhookDeliveryLog, error)
	GetDeliveryLogsByEndpoint(ctx context.Context, endpointID uuid.UUID) ([]*payment_engine_models.WebhookDeliveryLog, error)
	UpdateWebhookDeliveryLog(ctx context.Context, log *payment_engine_models.WebhookDeliveryLog) error
}
