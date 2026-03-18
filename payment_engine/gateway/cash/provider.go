package cash

import (
    "context"
    "time"
    payment_gateway "sharpnr-engine/payment_engine/gateway"
)

type CashProvider struct{}

func NewCashProvider() payment_gateway.GatewayProvider {
    return &CashProvider{}
}

func (c *CashProvider) Code() string {
    return "CASH"
}

// Cash has no gateway — order is self-created locally
func (c *CashProvider) CreateOrder(
    ctx context.Context,
    req payment_gateway.GatewayOrderRequest,
) (*payment_gateway.GatewayOrderResponse, error) {
    return &payment_gateway.GatewayOrderResponse{
        ProviderOrderID: req.InternalOrderID, // echo back our own ID
    }, nil
}

// Cash is self-confirming — whoever received it marks it done
func (c *CashProvider) VerifyPayment(
    ctx context.Context,
    req payment_gateway.GatewayVerifyRequest,
) (*payment_gateway.GatewayVerifyResponse, error) {
    now := time.Now()
    return &payment_gateway.GatewayVerifyResponse{
        Status:       "SUCCESS",
        GatewayTxnID: req.GatewayOrderID,
        PaidAt:       now,
        RawResponse:  `{"method":"CASH","auto_confirmed":true}`,
    }, nil
}

// Cash refunds are manual — just acknowledge
func (c *CashProvider) InitiateRefund(
    ctx context.Context,
    req payment_gateway.GatewayRefundRequest,
) (*payment_gateway.GatewayRefundResponse, error) {
    return &payment_gateway.GatewayRefundResponse{
        GatewayRefundID: "CASH-REFUND-MANUAL",
        Status:          "SUCCESS",
        RawResponse:     `{"method":"CASH","refund":"manual"}`,
    }, nil
}

// Cash has no status to fetch — always confirmed
func (c *CashProvider) FetchStatus(
    ctx context.Context,
    gatewayTxnID string,
) (*payment_gateway.GatewayStatusResponse, error) {
    return &payment_gateway.GatewayStatusResponse{
        Status:       "SUCCESS",
        GatewayTxnID: gatewayTxnID,
    }, nil
}