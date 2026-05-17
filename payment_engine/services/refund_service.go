package payment_engine_service

import (
	"context"
	"fmt"

	payment_constants "sharpnr-engine/payment_engine/constant"
	dbtx "sharpnr-engine/pkg/dbTx"

	contracts "sharpnr-engine/payment_engine/contracts"
	payment_gateway "sharpnr-engine/payment_engine/gateway"
	models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *paymentEngine) InitiateRefund(
	ctx context.Context,
	req contracts.RefundRequest,
) (*models.Refund, error) {

	payment, err := s.paymentRepo.GetPaymentByID(ctx, req.PaymentID)
	if err != nil {
		return nil, fmt.Errorf("payment not found: %w", err)
	}
	if payment.Status != payment_constants.StatusSuccess {
		return nil, fmt.Errorf("can only refund a successful payment, status: %s", payment.Status)
	}

	available := payment.AmountPaise - payment.RefundedPaise
	if req.AmountPaise > available {
		return nil, fmt.Errorf(
			"refund of %d paise exceeds refundable balance of %d paise",
			req.AmountPaise, available,
		)
	}

	refund := &models.Refund{
		ID:           uuid.New(),
		TenantID:     payment.TenantID,
		PaymentID:    payment.ID,
		RefundNumber: generateRefundNumber(),
		CurrencyCode: payment.CurrencyCode,
		AmountPaise:  req.AmountPaise,
		Status:       payment_constants.RefundStatusInitiated,
		Reason:       req.Reason,
		InitiatedBy:  req.InitiatedBy,
		SourceRef:    req.SourceRef,
	}

	// Offline payment — no gateway call
	if payment.GatewayConfigID == nil {
		return s.processOfflineRefund(ctx, refund, payment)
	}

	gateway, err := s.resolveGatewayByConfigID(ctx, *payment.GatewayConfigID)
	if err != nil {
		return nil, fmt.Errorf("resolve gateway: %w", err)
	}

	// Save before calling gateway — so audit trail exists even if gateway fails
	if err := s.paymentRepo.CreateRefund(ctx, refund); err != nil {
		return nil, fmt.Errorf("save refund: %w", err)
	}

	gatewayResp, err := gateway.Provider.InitiateRefund(ctx, payment_gateway.GatewayRefundRequest{
		GatewayTxnID:      *payment.GatewayTxnRef,
		InternalRefundID:  refund.ID.String(),
		RefundAmountPaise: req.AmountPaise,
		TotalAmountPaise:  payment.AmountPaise,
		Reason:            req.Reason,
		Config:            gateway.Config,
	})
	if err != nil {
		_ = s.paymentRepo.UpdateRefundStatus(ctx, refund.ID, payment_constants.RefundStatusFailed)
		return nil, fmt.Errorf("gateway initiate refund: %w", err)
	}

	// Mutate refund before closure — pointer is visible after WithTx
	newStatus := mapGatewayRefundStatus(gatewayResp.Status)
	rawResp := gatewayResp.RawResponse
	refund.GatewayRefundRef = &gatewayResp.GatewayRefundID
	refund.GatewayResponse = &rawResp
	refund.Status = newStatus

	err = s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.paymentRepo.UpdateRefundGatewayRefs(txCtx, refund); err != nil {
			return fmt.Errorf("update refund gateway refs: %w", err)
		}

		if newStatus == payment_constants.RefundStatusSuccess {
			if err := s.paymentRepo.IncrementPaymentRefunded(txCtx, payment.ID, req.AmountPaise); err != nil {
				return fmt.Errorf("increment payment refunded: %w", err)
			}

			newPaymentStatus := resolvePostRefundPaymentStatus(
				payment.AmountPaise,
				payment.RefundedPaise+req.AmountPaise,
			)
			if err := s.paymentRepo.UpdatePaymentStatus(txCtx, payment.ID, newPaymentStatus); err != nil {
				return fmt.Errorf("update payment status: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return refund, nil
}

func (s *paymentEngine) GetRefundsByPayment(
	ctx context.Context,
	paymentID uuid.UUID,
) ([]*models.Refund, error) {

	refunds, err := s.paymentRepo.GetRefundsByPaymentID(ctx, paymentID)
	if err != nil {
		return nil, fmt.Errorf("get refunds: %w", err)
	}
	return refunds, nil
}

func (s *paymentEngine) processOfflineRefund(
	ctx context.Context,
	refund *models.Refund,
	payment *models.Payment,
) (*models.Refund, error) {

	refund.Status = payment_constants.RefundStatusSuccess

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.paymentRepo.CreateRefund(txCtx, refund); err != nil {
			return fmt.Errorf("save offline refund: %w", err)
		}

		if err := s.paymentRepo.IncrementPaymentRefunded(txCtx, payment.ID, refund.AmountPaise); err != nil {
			return fmt.Errorf("increment payment refunded: %w", err)
		}

		newPaymentStatus := resolvePostRefundPaymentStatus(
			payment.AmountPaise,
			payment.RefundedPaise+refund.AmountPaise,
		)
		if err := s.paymentRepo.UpdatePaymentStatus(txCtx, payment.ID, newPaymentStatus); err != nil {
			return fmt.Errorf("update payment status: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return refund, nil
}

func mapGatewayRefundStatus(s string) string {
	switch s {
	case "SUCCESS":
		return payment_constants.RefundStatusSuccess
	case "FAILED":
		return payment_constants.RefundStatusFailed
	default:
		return payment_constants.RefundStatusProcessing
	}
}

func resolvePostRefundPaymentStatus(totalPaise, refundedPaise int64) string {
	if refundedPaise >= totalPaise {
		return payment_constants.StatusRefunded
	}
	return payment_constants.StatusPartiallyRefunded
}
