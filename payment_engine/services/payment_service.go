package payment_engine_service

import (
	"context"
	"fmt"
	"time"

	dbtx "sharpnr-engine/core/dbTx"
	payment_constants "sharpnr-engine/payment_engine/constant"
	contracts "sharpnr-engine/payment_engine/contracts"
	payment_gateway "sharpnr-engine/payment_engine/gateway"
	models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *paymentEngine) Pay(
	ctx context.Context,
	req contracts.PayRequest,
) (*models.Payment, error) {

	order, err := s.paymentRepo.GetOrderByID(ctx, req.PaymentOrderID)
	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}
	// helper functions
	if err := validateOrderPayable(order); err != nil {
		return nil, err
	}
	if req.AmountPaise <= 0 {
		return nil, fmt.Errorf("payment amount must be greater than zero")
	}

	payment := &models.Payment{
		ID:               uuid.New(),
		TenantID:         order.TenantID,
		LegalEntityID:    order.LegalEntityID,
		PaymentAccountID: order.PaymentAccountID,
		PaymentOrderID:   order.ID,
		PaymentNumber:    generatePaymentNumber(),
		CurrencyCode:     order.CurrencyCode,
		AmountPaise:      req.AmountPaise,
		Method:           req.Method,
		Status:           payment_constants.StatusInitiated,
		Notes:            req.Notes,
		CreatedBy:        req.CreatedBy,
	}

	if isOfflineMethod(req.Method) {
		payment.ReferenceNumber = req.ReferenceNumber
		return s.confirmOfflinePayment(ctx, payment, order)
	}

	gateway, err := s.resolveGateway(ctx, order.PaymentAccountID, req.Method)
	if err != nil {
		return nil, fmt.Errorf("resolve gateway: %w", err)
	}

	payment.GatewayConfigID = &gateway.ConfigID
	payment.Status = payment_constants.StatusProcessing

	if err := s.paymentRepo.CreatePayment(ctx, payment); err != nil {
		return nil, fmt.Errorf("save payment: %w", err)
	}

	gatewayResp, err := gateway.Provider.CreateOrder(ctx, payment_gateway.GatewayOrderRequest{
		InternalOrderID: payment.ID.String(),
		AmountPaise:     req.AmountPaise,
		CurrencyCode:    order.CurrencyCode,
		Config:          gateway.Config,
		Meta: map[string]string{
			"payer_name": order.PayerName,
		},
	})
	if err != nil {
		_ = s.paymentRepo.UpdatePaymentStatus(ctx, payment.ID, payment_constants.StatusFailed)
		return nil, fmt.Errorf("gateway create order: %w", err)
	}

	rawResp := gatewayResp.RawResponse
	payment.GatewayOrderRef = &gatewayResp.ProviderOrderID
	payment.GatewayResponse = &rawResp

	if err := s.paymentRepo.UpdatePaymentGatewayRefs(ctx, payment); err != nil {
		return nil, fmt.Errorf("update gateway refs: %w", err)
	}

	return payment, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *paymentEngine) ConfirmPayment(
	ctx context.Context,
	req contracts.ConfirmPaymentRequest,
) (*models.Payment, error) {

	payment, err := s.paymentRepo.GetPaymentByGatewayOrderRef(ctx, req.GatewayOrderRef)
	if err != nil {
		return nil, fmt.Errorf("payment not found for gateway ref '%s': %w", req.GatewayOrderRef, err)
	}

	// Idempotent — already confirmed
	if payment.Status == payment_constants.StatusSuccess {
		return payment, nil
	}

	gateway, err := s.resolveGatewayByConfigID(ctx, *payment.GatewayConfigID)
	if err != nil {
		return nil, fmt.Errorf("resolve gateway: %w", err)
	}

	verifyResp, err := gateway.Provider.VerifyPayment(ctx, payment_gateway.GatewayVerifyRequest{
		GatewayTxnID:    req.GatewayTxnID,
		GatewayOrderID:  req.GatewayOrderRef,
		AmountPaise:     payment.AmountPaise,
		RawCallbackBody: req.RawCallbackBody,
		Headers:         req.Headers,
		Config:          gateway.Config,
	})
	if err != nil {
		return nil, fmt.Errorf("verify payment: %w", err)
	}

	newStatus := mapGatewayStatus(verifyResp.Status)
	paidAt := verifyResp.PaidAt
	rawResp := verifyResp.RawResponse

	// Mutate payment before the closure so it's available after
	payment.Status = newStatus
	payment.GatewayTxnRef = &verifyResp.GatewayTxnID
	payment.GatewayResponse = &rawResp
	payment.PaidAt = &paidAt

	err = s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.paymentRepo.UpdatePaymentConfirmation(txCtx, payment); err != nil {
			return fmt.Errorf("update payment confirmation: %w", err)
		}

		if newStatus == payment_constants.StatusSuccess {
			if err := s.syncOrderStatus(txCtx, payment.PaymentOrderID, payment.AmountPaise); err != nil {
				return fmt.Errorf("sync order status: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return payment, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *paymentEngine) PublishPaymentConfirmed(
	ctx context.Context,
	payment *models.Payment,
) {
	// Fetch order to get SourceMeta — contains session_id and invoice_id
	order, err := s.paymentRepo.GetOrderByID(ctx, payment.PaymentOrderID)
	if err != nil {
		fmt.Printf("WARNING: PublishPaymentConfirmed — fetch order failed: %v\n", err)
		return
	}

	if order.SourceMeta == nil {
		fmt.Printf("WARNING: PublishPaymentConfirmed — order has no source meta\n")
		return
	}

	// Publish — restaurant service handler picks this up
	// err = s.seroEngine.Publish(ctx, payment_constants.PaymentConfirmedEvent,
	// 	payment_events.PaymentConfirmedPayload{
	// 		PaymentID:      payment.ID.String(),
	// 		PaymentOrderID: payment.PaymentOrderID.String(),
	// 		TenantID:       payment.TenantID.String(),
	// 		LegalEntityID:  payment.LegalEntityID.String(),
	// 		Method:         payment.Method,
	// 		AmountPaise:    payment.AmountPaise,
	// 		PaidAt:         payment.PaidAt.Format(time.RFC3339),
	// 		SourceMeta:     *order.SourceMeta, // echoed back — restaurant reads session_id from here
	// 	},
	// )
	// if err != nil {
	// 	fmt.Printf("WARNING: PublishPaymentConfirmed — publish failed: %v\n", err)
	// }

	// TODO: uncomment when ready to wire sero event
	fmt.Printf(
		"[PaymentEngine] Payment confirmed — payment_id: %s | order_id: %s | method: %s | amount: %d paise | source_meta: %s\n",
		payment.ID.String(),
		payment.PaymentOrderID.String(),
		payment.Method,
		payment.AmountPaise,
		*order.SourceMeta,
	)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *paymentEngine) GetPayment(
	ctx context.Context,
	paymentID uuid.UUID,
) (*models.Payment, error) {

	payment, err := s.paymentRepo.GetPaymentByID(ctx, paymentID)
	if err != nil {
		return nil, fmt.Errorf("payment not found: %w", err)
	}
	return payment, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *paymentEngine) GetPaymentsByOrder(
	ctx context.Context,
	orderID uuid.UUID,
) ([]*models.Payment, error) {

	payments, err := s.paymentRepo.GetPaymentsByOrderID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("get payments by order: %w", err)
	}
	return payments, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *paymentEngine) confirmOfflinePayment(
	ctx context.Context,
	payment *models.Payment,
	order *models.PaymentOrder,
) (*models.Payment, error) {

	now := time.Now()
	payment.Status = payment_constants.StatusSuccess
	payment.PaidAt = &now

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.paymentRepo.CreatePayment(txCtx, payment); err != nil {
			return fmt.Errorf("save offline payment: %w", err)
		}

		if err := s.syncOrderStatus(txCtx, order.ID, payment.AmountPaise); err != nil {
			return fmt.Errorf("sync order status: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return payment, nil
}

/* ---------------------------------- <***> --------------------------------- */

func validateOrderPayable(order *models.PaymentOrder) error {
	switch order.Status {
	case payment_constants.OrderStatusPaid:
		return fmt.Errorf("order is already fully paid")
	case payment_constants.OrderStatusCancelled:
		return fmt.Errorf("cannot pay against a cancelled order")
	case payment_constants.OrderStatusExpired:
		return fmt.Errorf("order has expired")
	}
	if order.ExpiresAt != nil && time.Now().After(*order.ExpiresAt) {
		return fmt.Errorf("order has expired")
	}
	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func isOfflineMethod(method string) bool {
	switch method {
	case "CASH", "CHEQUE", "BANK_TRANSFER":
		return true
	}
	return false
}

/* ---------------------------------- <***> --------------------------------- */

func mapGatewayStatus(s string) string {
	switch s {
	case "SUCCESS":
		return payment_constants.StatusSuccess
	case "FAILED":
		return payment_constants.StatusFailed
	default:
		return payment_constants.StatusProcessing
	}
}

/* ---------------------------------- <***> --------------------------------- */
