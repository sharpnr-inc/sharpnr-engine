package payment_engine_service

import (
	"context"
	"github.com/google/uuid"
	dbtx "sharpnr-engine/core/dbTx"
	payment_engine_contracts "sharpnr-engine/payment_engine/contracts"
	payment_engine_models "sharpnr-engine/payment_engine/models"
	payment_engine_repo "sharpnr-engine/payment_engine/repo"
	"sharpnr-engine/sero"
	"sharpnr-engine/sharp_mailer_engine"
)

type paymentEngine struct {
	paymentRepo  payment_engine_repo.PaymentEngineRepoInterface
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
	txManager    dbtx.TxManager
	seroEngine   *sero.SEROEngine
}

// AllocatePayment implements [PaymentServiceInterface].
func (p *paymentEngine) AllocatePayment(ctx context.Context, req payment_engine_contracts.AllocateRequest) (*payment_engine_models.PaymentAllocation, error) {
	panic("unimplemented")
}

// CancelOrder implements [PaymentServiceInterface].
func (p *paymentEngine) CancelOrder(ctx context.Context, orderID uuid.UUID, reason string, by string) error {
	panic("unimplemented")
}

// ConfirmPayment implements [PaymentServiceInterface].
func (p *paymentEngine) ConfirmPayment(ctx context.Context, req payment_engine_contracts.ConfirmPaymentRequest) (*payment_engine_models.Payment, error) {
	panic("unimplemented")
}

// CreateOrder implements [PaymentServiceInterface].
func (p *paymentEngine) CreateOrder(ctx context.Context, req payment_engine_contracts.CreateOrderRequest) (*payment_engine_models.PaymentOrder, error) {
	panic("unimplemented")
}

// ExpireStaleOrders implements [PaymentServiceInterface].
func (p *paymentEngine) ExpireStaleOrders(ctx context.Context) error {
	panic("unimplemented")
}

// GetAllocationsByTarget implements [PaymentServiceInterface].
func (p *paymentEngine) GetAllocationsByTarget(ctx context.Context, targetType string, targetID string) ([]*payment_engine_models.PaymentAllocation, error) {
	panic("unimplemented")
}

// GetOrder implements [PaymentServiceInterface].
func (p *paymentEngine) GetOrder(ctx context.Context, orderID uuid.UUID) (*payment_engine_models.PaymentOrder, error) {
	panic("unimplemented")
}

// GetOutstandingAmount implements [PaymentServiceInterface].
func (p *paymentEngine) GetOutstandingAmount(ctx context.Context, targetType string, targetID string) (int64, error) {
	panic("unimplemented")
}

// GetPayment implements [PaymentServiceInterface].
func (p *paymentEngine) GetPayment(ctx context.Context, paymentID uuid.UUID) (*payment_engine_models.Payment, error) {
	panic("unimplemented")
}

// GetPaymentsByOrder implements [PaymentServiceInterface].
func (p *paymentEngine) GetPaymentsByOrder(ctx context.Context, orderID uuid.UUID) ([]*payment_engine_models.Payment, error) {
	panic("unimplemented")
}

// GetRefundsByPayment implements [PaymentServiceInterface].
func (p *paymentEngine) GetRefundsByPayment(ctx context.Context, paymentID uuid.UUID) ([]*payment_engine_models.Refund, error) {
	panic("unimplemented")
}

// InitiateRefund implements [PaymentServiceInterface].
func (p *paymentEngine) InitiateRefund(ctx context.Context, req payment_engine_contracts.RefundRequest) (*payment_engine_models.Refund, error) {
	panic("unimplemented")
}

// Pay implements [PaymentServiceInterface].
func (p *paymentEngine) Pay(ctx context.Context, req payment_engine_contracts.PayRequest) (*payment_engine_models.Payment, error) {
	panic("unimplemented")
}

// ReverseAllocation implements [PaymentServiceInterface].
func (p *paymentEngine) ReverseAllocation(ctx context.Context, allocationID uuid.UUID, by string, note string) error {
	panic("unimplemented")
}

func NewPaymentService(
	paymentRepo payment_engine_repo.PaymentEngineRepoInterface,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) *paymentEngine {
	return &paymentEngine{
		paymentRepo:  paymentRepo,
		mailerEngine: mailerEngine,
		txManager:    txManager,
		seroEngine:   seroEngine,
	}
}
