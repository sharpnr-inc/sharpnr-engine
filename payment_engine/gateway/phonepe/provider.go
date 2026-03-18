// PhonePeProvider implements GatewayProvider for the PhonePe payment gateway.
//
// REQUIRED CONFIG KEYS (from PaymentGatewayConfig.ConfigEncrypted):
//   merchant_id  - PhonePe MerchantID from dashboard
//   salt_key     - SaltKey from PhonePe dashboard (used for HMAC signing)
//   salt_index   - SaltIndex from PhonePe dashboard (usually "1")
//   webhook_url  - Injected by GatewayResolver from PaymentGatewayConfig.WebhookURL
//   is_test_mode - "true" for sandbox, "false" for production
//
// CHECKSUM:
//   PhonePe uses SHA256 HMAC for request signing and webhook verification.
//   Format: SHA256(payload + endpoint + saltKey) + "###" + saltIndex
//
// ENVIRONMENT URLs:
//   Sandbox:    https://api-preprod.phonepe.com/apis/pg-sandbox
//   Production: https://api.phonepe.com/apis/hermes
//
// WEBHOOK VERIFICATION:
//   PhonePe sends X-VERIFY header = SHA256(rawBody + saltKey) + "###" + saltIndex
//   The response body is base64 encoded JSON — must decode before parsing.

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

// API endpoints
const (
	sandboxBaseURL = "https://api-preprod.phonepe.com/apis/pg-sandbox"
	prodBaseURL    = "https://api.phonepe.com/apis/hermes"

	payEndpoint    = "/pg/v1/pay"
	statusEndpoint = "/pg/v1/status/%s/%s" // /status/{merchantId}/{merchantTransactionId}
	refundEndpoint = "/pg/v1/refund"
)

// PhonePe state → your internal PaymentStatus
var stateMap = map[string]string{
	"COMPLETED": "SUCCESS",
	"FAILED":    "FAILED",
	"ERROR":     "FAILED",
	"PENDING":   "PROCESSING",
}

type PhonePeProvider struct {
	httpClient *http.Client
}

func NewPhonePeProvider() payment_gateway.GatewayProvider {
	return &PhonePeProvider{
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (p *PhonePeProvider) Code() string { return "PHONEPE" }

// CreateOrder initiates a UPI QR payment request on PhonePe.
// Returns QRData — shown on second screen for customer to scan.
func (p *PhonePeProvider) CreateOrder(
	ctx context.Context,
	req payment_gateway.GatewayOrderRequest,
) (*payment_gateway.GatewayOrderResponse, error) {

	merchantID := req.Config["merchant_id"]
	saltKey := req.Config["salt_key"]
	saltIndex := req.Config["salt_index"]
	webhookURL := req.Config["webhook_url"]
	baseURL := p.baseURL(req.Config)

	// Build PhonePe pay request
	payload := phonePePayRequest{
		MerchantID:            merchantID,
		MerchantTransactionID: req.InternalOrderID, // maps back to Payment.ID on webhook
		Amount:                req.AmountPaise,
		CallbackURL:           webhookURL,
		PaymentInstrument:     phonePeInstrument{Type: "UPI_QR"},
	}

	// Encode → sign → send
	respBytes, err := p.signAndPost(ctx, baseURL+payEndpoint, payEndpoint, payload, saltKey, saltIndex)
	if err != nil {
		return nil, fmt.Errorf("phonepe CreateOrder: %w", err)
	}

	var apiResp phonePeAPIResponse
	if err := json.Unmarshal(respBytes, &apiResp); err != nil {
		return nil, fmt.Errorf("phonepe unmarshal response: %w", err)
	}
	if !apiResp.Success {
		return nil, fmt.Errorf("phonepe order failed [%s]: %s", apiResp.Code, apiResp.Message)
	}

	result := &payment_gateway.GatewayOrderResponse{
		ProviderOrderID: apiResp.Data.TransactionID,
		RawResponse:     string(respBytes),
	}

	// Extract QR string for second screen
	if apiResp.Data.InstrumentResponse != nil && apiResp.Data.InstrumentResponse.QRData != nil {
		result.QRData = apiResp.Data.InstrumentResponse.QRData
	}

	// Set expiry — PhonePe QR typically valid for 10 minutes
	expiry := time.Now().Add(10 * time.Minute)
	result.ExpiresAt = &expiry

	return result, nil
}

// VerifyPayment validates inbound webhook from PhonePe.
// Called from payment_engine/controllers/webhook_handlers.go.
//
// PhonePe webhook flow:
//  1. PhonePe POSTs to your WebhookURL with X-VERIFY header
//  2. Body contains base64 encoded JSON
//  3. Verify: SHA256(body + saltKey) + "###" + saltIndex must match X-VERIFY
//  4. Decode base64 body to get actual payment status
func (p *PhonePeProvider) VerifyPayment(
	ctx context.Context,
	req payment_gateway.GatewayVerifyRequest,
) (*payment_gateway.GatewayVerifyResponse, error) {

	saltKey := req.Config["salt_key"]
	saltIndex := req.Config["salt_index"]

	// 1. Verify HMAC signature from X-VERIFY header
	incomingChecksum := req.Headers["x-verify"]
	expectedChecksum := p.checksum(req.RawCallbackBody, saltKey, saltIndex, "")
	if incomingChecksum != expectedChecksum {
		return nil, fmt.Errorf("phonepe VerifyPayment: invalid webhook signature")
	}

	// 2. Parse wrapper — body is { "response": "<base64>" }
	var wrapper phonePeWebhookPayload
	if err := json.Unmarshal([]byte(req.RawCallbackBody), &wrapper); err != nil {
		return nil, fmt.Errorf("phonepe VerifyPayment: unmarshal wrapper: %w", err)
	}

	// 3. Decode base64 inner payload
	decoded, err := base64.StdEncoding.DecodeString(wrapper.Response)
	if err != nil {
		return nil, fmt.Errorf("phonepe VerifyPayment: decode base64: %w", err)
	}

	// 4. Parse actual payment data
	var data phonePeData
	if err := json.Unmarshal(decoded, &data); err != nil {
		return nil, fmt.Errorf("phonepe VerifyPayment: unmarshal data: %w", err)
	}

	status := p.mapState(data.State)

	return &payment_gateway.GatewayVerifyResponse{
		Status:          status,
		GatewayTxnID:    data.TransactionID,
		GatewayOrderID:  data.MerchantTransactionID, // = your Payment.ID
		PaidAmountPaise: data.Amount,
		PaidAt:          time.Now(),
		RawResponse:     string(decoded),
	}, nil
}

// InitiateRefund requests a refund via PhonePe API.
func (p *PhonePeProvider) InitiateRefund(
	ctx context.Context,
	req payment_gateway.GatewayRefundRequest,
) (*payment_gateway.GatewayRefundResponse, error) {

	merchantID := req.Config["merchant_id"]
	saltKey := req.Config["salt_key"]
	saltIndex := req.Config["salt_index"]
	webhookURL := req.Config["webhook_url"]
	baseURL := p.baseURL(req.Config)

	payload := phonePeRefundRequest{
		MerchantID:            merchantID,
		MerchantTransactionID: req.InternalRefundID, // = Refund.ID
		OriginalTransactionID: req.GatewayTxnID,     // = original Payment.GatewayTxnRef
		Amount:                req.RefundAmountPaise,
		CallbackURL:           webhookURL,
	}

	respBytes, err := p.signAndPost(ctx, baseURL+refundEndpoint, refundEndpoint, payload, saltKey, saltIndex)
	if err != nil {
		return nil, fmt.Errorf("phonepe InitiateRefund: %w", err)
	}

	var apiResp phonePeAPIResponse
	if err := json.Unmarshal(respBytes, &apiResp); err != nil {
		return nil, fmt.Errorf("phonepe unmarshal refund response: %w", err)
	}
	if !apiResp.Success {
		return nil, fmt.Errorf("phonepe refund failed [%s]: %s", apiResp.Code, apiResp.Message)
	}

	return &payment_gateway.GatewayRefundResponse{
		GatewayRefundID: apiResp.Data.TransactionID,
		Status:          p.mapState(apiResp.Data.State),
		RawResponse:     string(respBytes),
	}, nil
}

// FetchStatus polls PhonePe for current transaction status.
// Used for reconciliation when webhook was missed.
func (p *PhonePeProvider) FetchStatus(
	ctx context.Context,
	gatewayTxnID string,
	config map[string]string,
) (*payment_gateway.GatewayStatusResponse, error) {

	merchantID := config["merchant_id"]
	saltKey := config["salt_key"]
	saltIndex := config["salt_index"]
	baseURL := p.baseURL(config)

	endpoint := fmt.Sprintf(statusEndpoint, merchantID, gatewayTxnID)
	url := baseURL + endpoint

	// Status check uses GET with checksum in header
	checksum := p.checksum("", saltKey, saltIndex, endpoint)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("phonepe FetchStatus build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-VERIFY", checksum)
	httpReq.Header.Set("X-MERCHANT-ID", merchantID)

	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("phonepe FetchStatus http: %w", err)
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)

	var statusResp phonePeStatusResponse
	if err := json.Unmarshal(respBytes, &statusResp); err != nil {
		return nil, fmt.Errorf("phonepe FetchStatus unmarshal: %w", err)
	}

	return &payment_gateway.GatewayStatusResponse{
		Status:          p.mapState(statusResp.Data.State),
		GatewayTxnID:    statusResp.Data.TransactionID,
		GatewayOrderID:  statusResp.Data.MerchantTransactionID,
		PaidAmountPaise: statusResp.Data.Amount,
		RawResponse:     string(respBytes),
	}, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// INTERNAL HELPERS
// ─────────────────────────────────────────────────────────────────────────────

// signAndPost encodes payload to base64, computes checksum, POSTs to PhonePe.
func (p *PhonePeProvider) signAndPost(
	ctx context.Context,
	url string,
	endpoint string,
	payload interface{},
	saltKey string,
	saltIndex string,
) ([]byte, error) {

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal payload: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(payloadBytes)
	checksum := p.checksum(encoded, saltKey, saltIndex, endpoint)

	body, _ := json.Marshal(map[string]string{"request": encoded})

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-VERIFY", checksum)

	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("http post: %w", err)
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBytes))
	}

	return respBytes, nil
}

// checksum computes the PhonePe HMAC signature.
// Formula: SHA256(data + endpoint + saltKey) + "###" + saltIndex
// For webhooks: data = raw body, endpoint = ""
// For API calls: data = base64(payload), endpoint = "/pg/v1/pay" etc.
func (p *PhonePeProvider) checksum(data, saltKey, saltIndex, endpoint string) string {
	input := data + endpoint + saltKey
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x###%s", hash, saltIndex)
}

func (p *PhonePeProvider) baseURL(config map[string]string) string {
	if config["is_test_mode"] == "true" {
		return sandboxBaseURL
	}
	return prodBaseURL
}

func (p *PhonePeProvider) mapState(state string) string {
	if status, ok := stateMap[state]; ok {
		return status
	}
	return "PROCESSING"
}
