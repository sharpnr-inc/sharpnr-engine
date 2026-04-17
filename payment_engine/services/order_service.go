package payment_engine_service

import (
	"context"
	"fmt"
	"time"

	payment_constants "sharpnr-engine/payment_engine/constant"
	payment_payment_constants "sharpnr-engine/payment_engine/constant"
	contracts "sharpnr-engine/payment_engine/contracts"
	models "sharpnr-engine/payment_engine/models"
	payment_engine_models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

func (s *paymentEngine) CreateOrder(
	ctx context.Context,
	req contracts.CreateOrderRequest,
) (*models.PaymentOrder, error) {

	if req.AmountPaise <= 0 {
		return nil, fmt.Errorf("order amount must be greater than zero")
	}
	if _, err := s.paymentRepo.GetPaymentAccountByID(ctx, req.PaymentAccountID); err != nil {
		return nil, fmt.Errorf("payment account not found: %w", err)
	}
	if req.CurrencyCode == "" {
		req.CurrencyCode = "INR"
	}

	order := &models.PaymentOrder{
		ID:               uuid.New(),
		TenantID:         req.TenantID,
		LegalEntityID:    req.LegalEntityID,
		PaymentAccountID: req.PaymentAccountID,
		OrderNumber:      generateOrderNumber(),
		PayerType:        req.PayerType,
		PayerRef:         req.PayerRef,
		PayerName:        req.PayerName,
		CurrencyCode:     req.CurrencyCode,
		AmountPaise:      req.AmountPaise,
		PaidPaise:        0,
		Status:           "Pending",
		Source:           req.Source,
		SourceRefID:      req.SourceRefID,
		SourceMeta:       req.SourceMeta,
		CreatedBy:        req.CreatedBy,
	}

	if req.ExpiresInMinutes > 0 {
		t := time.Now().Add(time.Duration(req.ExpiresInMinutes) * time.Minute)
		order.ExpiresAt = &t
	}

	if err := s.paymentRepo.CreateOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("create order: %w", err)
	}
	return order, nil
}

func (s *paymentEngine) GetOrder(
	ctx context.Context,
	orderID uuid.UUID,
) (*models.PaymentOrder, error) {

	order, err := s.paymentRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}
	return order, nil
}

func (s *paymentEngine) CancelOrder(
	ctx context.Context,
	orderID uuid.UUID,
	reason string,
	by string,
) error {

	order, err := s.paymentRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("order not found: %w", err)
	}
	if order.Status == payment_payment_constants.OrderStatusPaid {
		return fmt.Errorf("cannot cancel a fully paid order")
	}
	if order.Status == payment_constants.OrderStatusCancelled {
		return fmt.Errorf("order is already cancelled")
	}
	return s.paymentRepo.UpdateOrderStatus(ctx, orderID, payment_constants.OrderStatusCancelled)
}

func (s *paymentEngine) ExpireStaleOrders(ctx context.Context) error {
	orders, err := s.paymentRepo.GetExpiredPendingOrders(ctx, time.Now())
	if err != nil {
		return fmt.Errorf("fetch expired orders: %w", err)
	}
	for _, order := range orders {
		if err := s.paymentRepo.UpdateOrderStatus(ctx, order.ID, payment_constants.OrderStatusExpired); err != nil {
			fmt.Printf("WARNING: failed to expire order %s: %v\n", order.ID, err)
		}
	}
	return nil
}

func (s *paymentEngine) syncOrderStatus(
	ctx context.Context,
	orderID uuid.UUID,
	addedPaise int64,
) error {

	order, err := s.paymentRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	newPaid := order.PaidPaise + addedPaise

	var newStatus string
	switch {
	case newPaid >= order.AmountPaise:
		newStatus = payment_constants.OrderStatusPaid
	case newPaid > 0:
		newStatus = payment_constants.OrderStatusPartiallyPaid
	default:
		newStatus = payment_constants.OrderStatusPending
	}

	return s.paymentRepo.UpdateOrderPaidAmount(ctx, orderID, newPaid, newStatus)
}

func (c *paymentEngine) GetOrderByLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
) ([]*payment_engine_models.PaymentOrder, error) {

	// 1. Optional: Add authorization or validation logic here
	// e.g., Verify if emplID has access to orgID or legalEntityID

	// 2. Call the repository method we just implemented
	orders, err := c.paymentRepo.ListOrderByLegalEntityID(ctx, legalEntityID)
	if err != nil {
		// Log the error with context if you have a logger
		return nil, fmt.Errorf("failed to fetch payment orders for legal entity %s: %w", legalEntityID, err)
	}

	// 3. Return the results
	return orders, nil
}
