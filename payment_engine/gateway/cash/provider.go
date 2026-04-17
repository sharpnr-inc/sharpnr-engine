// CashProvider handles Cash and Card payment methods.
// These are "offline" methods — no external gateway API is called.
//
// CASH:
//   Staff receives physical cash and records it.
//   CreateOrder echoes back the internal order ID.
//   VerifyPayment self-confirms immediately.
//   Refund is a manual physical handback — marked SUCCESS immediately.
//
// CARD:
//   POS terminal is external hardware not integrated via API.
//   Staff confirms after terminal approves.
//   Same self-confirming behavior as Cash from the engine's perspective.
//   The distinction (CASH vs CARD) is only meaningful for reporting.

package cash

import (
	"context"
	"fmt"
	"time"

	payment_gateway "sharpnr-engine/payment_engine/gateway"
)

type CashProvider struct{}

func NewCashProvider() payment_gateway.GatewayProvider {
	return &CashProvider{}
}

func (c *CashProvider) Code() string { return "CASH" }

// CreateOrder — no API call, echo back internal ID
func (c *CashProvider) CreateOrder(
	_ context.Context,
	req payment_gateway.GatewayOrderRequest,
) (*payment_gateway.GatewayOrderResponse, error) {
	return &payment_gateway.GatewayOrderResponse{
		ProviderOrderID: req.InternalOrderID,
		QRData:          nil, // no QR for cash
		RawResponse: fmt.Sprintf(
			`{"method":"CASH","internal_order_id":"%s","auto_confirmed":true}`,
			req.InternalOrderID,
		),
	}, nil
}

// VerifyPayment — self-confirming, no signature check needed
func (c *CashProvider) VerifyPayment(
	_ context.Context,
	req payment_gateway.GatewayVerifyRequest,
) (*payment_gateway.GatewayVerifyResponse, error) {
	now := time.Now()
	return &payment_gateway.GatewayVerifyResponse{
		Status:          "SUCCESS",
		GatewayTxnID:   req.GatewayOrderID,
		GatewayOrderID: req.GatewayOrderID,
		PaidAmountPaise: req.AmountPaise,
		PaidAt:          now,
		RawResponse:     `{"method":"CASH","auto_confirmed":true}`,
	}, nil
}

// InitiateRefund — manual physical handback, mark success immediately
func (c *CashProvider) InitiateRefund(
	_ context.Context,
	req payment_gateway.GatewayRefundRequest,
) (*payment_gateway.GatewayRefundResponse, error) {
	return &payment_gateway.GatewayRefundResponse{
		GatewayRefundID: fmt.Sprintf("CASH-REFUND-%s", req.InternalRefundID),
		Status:          "SUCCESS",
		RawResponse: fmt.Sprintf(
			`{"method":"CASH","refund_id":"%s","manual":true}`,
			req.InternalRefundID,
		),
	}, nil
}

// FetchStatus — always success for cash, no remote state to check
func (c *CashProvider) FetchStatus(
	_ context.Context,
	gatewayTxnID string,
	_ map[string]string,
) (*payment_gateway.GatewayStatusResponse, error) {
	return &payment_gateway.GatewayStatusResponse{
		Status:       "SUCCESS",
		GatewayTxnID: gatewayTxnID,
		RawResponse:  `{"method":"CASH","status":"SUCCESS"}`,
	}, nil
}