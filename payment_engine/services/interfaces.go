package payment_engine_service

import (
	"context"
	payment_engine_contracts "sharpnr-engine/payment_engine/contracts"
	payment_engine_models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

type PaymentServiceInterface interface {
	CreateOrder(ctx context.Context, req payment_engine_contracts.CreateOrderRequest) (*payment_engine_models.PaymentOrder, error)
	GetOrder(ctx context.Context, orderID uuid.UUID) (*payment_engine_models.PaymentOrder, error)
	CancelOrder(ctx context.Context, orderID uuid.UUID, reason string, by string) error
	ExpireStaleOrders(ctx context.Context) error // called by a cron job

	// Payment
	Pay(ctx context.Context, req payment_engine_contracts.PayRequest) (*payment_engine_models.Payment, error)
	ConfirmPayment(ctx context.Context, req payment_engine_contracts.ConfirmPaymentRequest) (*payment_engine_models.Payment, error)
	GetPayment(ctx context.Context, paymentID uuid.UUID) (*payment_engine_models.Payment, error)
	GetPaymentsByOrder(ctx context.Context, orderID uuid.UUID) ([]*payment_engine_models.Payment, error)

	// Allocation
	AllocatePayment(ctx context.Context, req payment_engine_contracts.AllocateRequest) (*payment_engine_models.PaymentAllocation, error)
	ReverseAllocation(ctx context.Context, allocationID uuid.UUID, by string, note string) error
	GetAllocationsByTarget(ctx context.Context, targetType string, targetID string) ([]*payment_engine_models.PaymentAllocation, error)

	// Refunds
	InitiateRefund(ctx context.Context, req payment_engine_contracts.RefundRequest) (*payment_engine_models.Refund, error)
	GetRefundsByPayment(ctx context.Context, paymentID uuid.UUID) ([]*payment_engine_models.Refund, error)

	// Queries used by invoice engine / checkout
	GetOutstandingAmount(ctx context.Context, targetType string, targetID string) (int64, error) // in paise
}
