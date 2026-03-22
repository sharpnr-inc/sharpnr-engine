package payment_engine_service

import (
	"context"
	payment_engine_contracts "sharpnr-engine/payment_engine/contracts"
	payment_engine_models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

type PaymentServiceInterface interface {

	// ─── Account ──────────────────────────────────────────────────────────────
	CreatePaymentAccount(ctx context.Context, orgID, emplID uuid.UUID, req payment_engine_contracts.CreatePaymentAccountRequest) (*payment_engine_models.PaymentAccount, error)
	ListPaymentAccounts(ctx context.Context, tenantID uuid.UUID) ([]*payment_engine_models.PaymentAccount, error)
	TogglePaymentAccount(ctx context.Context, id uuid.UUID, active bool) error

	// ─── Gateway Config ───────────────────────────────────────────────────────
	AddGatewayConfig(ctx context.Context, req payment_engine_contracts.AddGatewayConfigRequest) (*payment_engine_models.PaymentGatewayConfig, error)
	ListGatewayConfigs(ctx context.Context, paymentAccountID uuid.UUID) ([]*payment_engine_models.PaymentGatewayConfig, error)
	UpdateGatewayConfig(ctx context.Context, id uuid.UUID, req payment_engine_contracts.UpdateGatewayConfigRequest) error
	ToggleGatewayConfig(ctx context.Context, id uuid.UUID, active bool) error
	RemoveGatewayConfig(ctx context.Context, id uuid.UUID) error

	// ─── Order ────────────────────────────────────────────────────────────────
	CreateOrder(ctx context.Context, req payment_engine_contracts.CreateOrderRequest) (*payment_engine_models.PaymentOrder, error)
	GetOrder(ctx context.Context, orderID uuid.UUID) (*payment_engine_models.PaymentOrder, error)
	CancelOrder(ctx context.Context, orderID uuid.UUID, reason string, by string) error
	ExpireStaleOrders(ctx context.Context) error

	// ─── Payment ──────────────────────────────────────────────────────────────
	Pay(ctx context.Context, req payment_engine_contracts.PayRequest) (*payment_engine_models.Payment, error)
	ConfirmPayment(ctx context.Context, req payment_engine_contracts.ConfirmPaymentRequest) (*payment_engine_models.Payment, error)
	GetPayment(ctx context.Context, paymentID uuid.UUID) (*payment_engine_models.Payment, error)
	GetPaymentsByOrder(ctx context.Context, orderID uuid.UUID) ([]*payment_engine_models.Payment, error)

	// ─── Allocation ───────────────────────────────────────────────────────────
	AllocatePayment(ctx context.Context, req payment_engine_contracts.AllocateRequest) (*payment_engine_models.PaymentAllocation, error)
	ReverseAllocation(ctx context.Context, req payment_engine_contracts.ReverseAllocationRequest) error
	GetAllocationsByTarget(ctx context.Context, targetType string, targetID string) ([]*payment_engine_models.PaymentAllocation, error)
	GetOutstandingAmount(ctx context.Context, targetType string, targetID string) (int64, error)

	// ─── Refund ───────────────────────────────────────────────────────────────
	InitiateRefund(ctx context.Context, req payment_engine_contracts.RefundRequest) (*payment_engine_models.Refund, error)
	GetRefundsByPayment(ctx context.Context, paymentID uuid.UUID) ([]*payment_engine_models.Refund, error)

	// ─── Webhook ──────────────────────────────────────────────────────────────
	RegisterWebhookEndpoint(ctx context.Context, req payment_engine_contracts.RegisterWebhookRequest) (*payment_engine_models.WebhookEndpoint, error)
	ListWebhookEndpoints(ctx context.Context, tenantID uuid.UUID) ([]*payment_engine_models.WebhookEndpoint, error)
	RemoveWebhookEndpoint(ctx context.Context, id uuid.UUID) error
}
