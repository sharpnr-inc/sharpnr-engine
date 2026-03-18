package phonepe

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	payment_gateway "sharpnr-engine/payment_engine/gateway"
)

const (
	prodBaseURL    = "https://api.phonepe.com/apis/hermes"
	sandboxBaseURL = "https://api-preprod.phonepe.com/apis/pg-sandbox"
	payEndpoint    = "/pg/v1/pay"
	statusEndpoint = "/pg/v1/status"
	refundEndpoint = "/pg/v1/refund"
)

type PhonePeProvider struct {
	httpClient *http.Client
}

func NewPhonePeProvider() payment_gateway.GatewayProvider {
	return &PhonePeProvider{
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (p *PhonePeProvider) Code() string { return "PHONEPE" }

// Config keys expected in req.Config:
// "merchant_id", "salt_key", "salt_index", "is_test_mode"
func (p *PhonePeProvider) CreateOrder(
	ctx context.Context,
	req payment_gateway.GatewayOrderRequest,
) (*payment_gateway.GatewayOrderResponse, error) {

	merchantID := req.Config["merchant_id"]
	saltKey     := req.Config["salt_key"]
	saltIndex   := req.Config["salt_index"]
	baseURL     := p.resolveBaseURL(req.Config)

	payload := phonePePayRequest{
		MerchantID:            merchantID,
		MerchantTransactionID: req.InternalOrderID,
		Amount:                req.AmountPaise,
		CallbackURL:           req.CallbackURL,
		PaymentInstrument:     phonePeInstrument{Type: "PAY_PAGE"},
	}

	if req.RedirectURL != nil {
		payload.RedirectURL  = *req.RedirectURL
		payload.RedirectMode = "POST"
	}
	if req.PayerPhone != nil {
		payload.MobileNumber = *req.PayerPhone
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("phonepe: marshal payload: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(payloadBytes)
	checksum := p.computeChecksum(encoded, saltKey, saltIndex, payEndpoint)

	body, _ := json.Marshal(map[string]string{"request": encoded})
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+payEndpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("phonepe: build request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-VERIFY", checksum)

	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("phonepe: http call failed: %w", err)
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)
	var phonePeResp phonePeResponse
	if err := json.Unmarshal(respBytes, &phonePeResp); err != nil {
		return nil, fmt.Errorf("phonepe: unmarshal response: %w", err)
	}

	if !phonePeResp.Success {
		return nil, fmt.Errorf("phonepe: order creation failed: %s - %s", phonePeResp.Code, phonePeResp.Message)
	}

	result := &payment_gateway.GatewayOrderResponse{
		ProviderOrderID: phonePeResp.Data.TransactionID,
		RawResponse:     string(respBytes),
	}

	if phonePeResp.Data.InstrumentResponse != nil {
		if phonePeResp.Data.InstrumentResponse.RedirectInfo != nil {
			result.PaymentURL = &phonePeResp.Data.InstrumentResponse.RedirectInfo.URL
		}
		if phonePeResp.Data.InstrumentResponse.QRData != nil {
			result.QRCode = phonePeResp.Data.InstrumentResponse.QRData
		}
	}

	return result, nil
}

// Called from webhook_handlers.go when PhonePe POSTs to your callback URL
func (p *PhonePeProvider) VerifyPayment(
	ctx context.Context,
	req payment_gateway.GatewayVerifyRequest,
) (*payment_gateway.GatewayVerifyResponse, error) {

	saltKey   := req.Config["salt_key"]
	saltIndex := req.Config["salt_index"]

	// Verify the X-VERIFY header signature
	incomingChecksum := req.Headers["x-verify"]
	expectedChecksum := p.computeChecksum(req.RawCallbackBody, saltKey, saltIndex, "")
	if incomingChecksum != expectedChecksum {
		return nil, fmt.Errorf("phonepe: invalid webhook signature")
	}

	// Decode base64 response body
	var callbackPayload phonePeCallbackPayload
	if err := json.Unmarshal([]byte(req.RawCallbackBody), &callbackPayload); err != nil {
		return nil, fmt.Errorf("phonepe: unmarshal callback: %w", err)
	}

	decoded, err := base64.StdEncoding.DecodeString(callbackPayload.Response)
	if err != nil {
		return nil, fmt.Errorf("phonepe: decode callback response: %w", err)
	}

	var data phonePeData
	if err := json.Unmarshal(decoded, &data); err != nil {
		return nil, fmt.Errorf("phonepe: unmarshal decoded callback: %w", err)
	}

	status := p.mapStatus(data.State)

	return &payment_gateway.GatewayVerifyResponse{
		Status:          status,
		GatewayTxnID:    data.TransactionID,
		GatewayOrderID:  data.MerchantTransactionID,
		PaidAmountPaise: data.Amount,
		PaidAt:          time.Now(),
		RawResponse:     string(decoded),
	}, nil
}

func (p *PhonePeProvider) InitiateRefund(
	ctx context.Context,
	req payment_gateway.GatewayRefundRequest,
) (*payment_gateway.GatewayRefundResponse, error) {

	saltKey   := req.Config["salt_key"]
	saltIndex := req.Config["salt_index"]
	merchantID := req.Config["merchant_id"]
	baseURL    := p.resolveBaseURL(req.Config)

	payload := phonePeRefundRequest{
		MerchantID:            merchantID,
		MerchantTransactionID: req.InternalRefundID,
		OriginalTransactionID: req.GatewayTxnID,
		Amount:                req.RefundAmountPaise,
	}

	payloadBytes, _ := json.Marshal(payload)
	encoded := base64.StdEncoding.EncodeToString(payloadBytes)
	checksum := p.computeChecksum(encoded, saltKey, saltIndex, refundEndpoint)

	body, _ := json.Marshal(map[string]string{"request": encoded})
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+refundEndpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("phonepe: build refund request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-VERIFY", checksum)

	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("phonepe: refund http call: %w", err)
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)
	var phonePeResp phonePeResponse
	json.Unmarshal(respBytes, &phonePeResp)

	if !phonePeResp.Success {
		return nil, fmt.Errorf("phonepe: refund failed: %s - %s", phonePeResp.Code, phonePeResp.Message)
	}

	return &payment_gateway.GatewayRefundResponse{
		GatewayRefundID: phonePeResp.Data.TransactionID,
		Status:          p.mapStatus(phonePeResp.Data.State),
		RawResponse:     string(respBytes),
	}, nil
}

func (p *PhonePeProvider) FetchStatus(
	ctx context.Context,
	gatewayTxnID string,
) (*payment_gateway.GatewayStatusResponse, error) {
	// PhonePe status check requires config — not available here
	// Use VerifyPayment for webhook-driven flow
	// This is available for polling if needed — requires config injection
	return nil, fmt.Errorf("phonepe: FetchStatus requires config — use VerifyPayment via webhook")
}

// -----------------------------------------------
// HELPERS
// -----------------------------------------------

func (p *PhonePeProvider) computeChecksum(payload, saltKey, saltIndex, endpoint string) string {
	data := payload + endpoint + saltKey
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x###%s", hash, saltIndex)
}

func (p *PhonePeProvider) resolveBaseURL(config map[string]string) string {
	if config["is_test_mode"] == "true" {
		return sandboxBaseURL
	}
	return prodBaseURL
}

func (p *PhonePeProvider) mapStatus(phonePeState string) string {
	switch phonePeState {
	case "COMPLETED":
		return "SUCCESS"
	case "FAILED", "ERROR":
		return "FAILED"
	case "PENDING":
		return "PROCESSING"
	default:
		return "PROCESSING"
	}
}