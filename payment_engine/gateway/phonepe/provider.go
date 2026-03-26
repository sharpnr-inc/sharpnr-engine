package phonepe

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	payment_gateway "sharpnr-engine/payment_engine/gateway"
)

// API endpoints
const (
	sandboxBaseURL = "https://api-preprod.phonepe.com/apis/pg-sandbox"
	prodBaseURL    = "https://api.phonepe.com/apis/hermes"

	payEndpoint    = "/pg/v1/pay"
	statusEndpoint = "/pg/v1/status/%s/%s"
	refundEndpoint = "/pg/v1/refund"
)

// PhonePe state → internal PaymentStatus
var stateMap = map[string]string{
	"COMPLETED": "SUCCESS",
	"FAILED":    "FAILED",
	"ERROR":     "FAILED",
	"PENDING":   "PROCESSING",
}

type PhonePeProvider struct {
	httpClient *http.Client
	logger     *slog.Logger
}

func NewPhonePeProvider(logger *slog.Logger) payment_gateway.GatewayProvider {
	return &PhonePeProvider{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		logger:     logger.With("provider", "PHONEPE"),
	}
}

func (p *PhonePeProvider) Code() string { return "PHONEPE" }

/* ---------------------------------- <***> --------------------------------- */

func (p *PhonePeProvider) CreateOrder(
	ctx context.Context,
	req payment_gateway.GatewayOrderRequest,
) (*payment_gateway.GatewayOrderResponse, error) {
	p.logger.Info("Initiating CreateOrder",
		slog.String("internal_id", req.InternalOrderID))

	merchantID := req.Config["client_id"]
	saltKey := req.Config["client_secret"]
	saltIndex := req.Config["client_version"]

	webhookURL := req.Config["webhook_url"]
	baseURL := p.baseURL(req.Config)

	payload := phonePePayRequest{
		MerchantID:            merchantID,
		MerchantTransactionID: req.InternalOrderID,
		Amount:                req.AmountPaise,
		CallbackURL:           webhookURL,
		PaymentInstrument:     phonePeInstrument{Type: "UPI_QR"},
	}

	respBytes, err := p.signAndPost(ctx, baseURL+payEndpoint, payEndpoint, payload, saltKey, saltIndex)
	if err != nil {
		return nil, err
	}

	var apiResp phonePeAPIResponse
	if err := json.Unmarshal(respBytes, &apiResp); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	if !apiResp.Success {
		p.logger.Warn("PhonePe order creation failed", slog.String("msg", apiResp.Message))
		return nil, fmt.Errorf("phonepe failed: %s", apiResp.Message)
	}

	p.logger.Info("Order created successfully", slog.String("gateway_id", apiResp.Data.TransactionID))

	result := &payment_gateway.GatewayOrderResponse{
		ProviderOrderID: apiResp.Data.TransactionID,
		RawResponse:     string(respBytes),
	}

	if apiResp.Data.InstrumentResponse != nil && apiResp.Data.InstrumentResponse.QRData != nil {
		result.QRData = apiResp.Data.InstrumentResponse.QRData
	}

	expiry := time.Now().Add(10 * time.Minute)
	result.ExpiresAt = &expiry

	return result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (p *PhonePeProvider) VerifyPayment(
	ctx context.Context,
	req payment_gateway.GatewayVerifyRequest,
) (*payment_gateway.GatewayVerifyResponse, error) {
	p.logger.Info("Verifying webhook")

	saltKey := req.Config["salt_key"]
	saltIndex := req.Config["salt_index"]

	incomingChecksum := req.Headers["x-verify"]
	expectedChecksum := p.checksum(req.RawCallbackBody, saltKey, saltIndex, "")

	if incomingChecksum != expectedChecksum {
		p.logger.Error("Checksum mismatch", slog.String("received", incomingChecksum))
		return nil, fmt.Errorf("invalid signature")
	}

	var wrapper phonePeWebhookPayload
	json.Unmarshal([]byte(req.RawCallbackBody), &wrapper)

	decoded, _ := base64.StdEncoding.DecodeString(wrapper.Response)
	var data phonePeData
	json.Unmarshal(decoded, &data)

	status := p.mapState(data.State)
	p.logger.Info("Webhook verified", slog.String("status", status), slog.String("txn_id", data.TransactionID))

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
	p.logger.Info("Initiating refund", slog.String("orig_txn", req.GatewayTxnID))

	merchantID := req.Config["client_id"]
	saltKey := req.Config["client_secret"]
	saltIndex := req.Config["salt_index"]
	baseURL := p.baseURL(req.Config)

	payload := phonePeRefundRequest{
		MerchantID:            merchantID,
		MerchantTransactionID: req.InternalRefundID,
		OriginalTransactionID: req.GatewayTxnID,
		Amount:                req.RefundAmountPaise,
		CallbackURL:           req.Config["webhook_url"],
	}

	respBytes, err := p.signAndPost(ctx, baseURL+refundEndpoint, refundEndpoint, payload, saltKey, saltIndex)
	if err != nil {
		return nil, err
	}

	var apiResp phonePeAPIResponse
	json.Unmarshal(respBytes, &apiResp)

	status := p.mapState(apiResp.Data.State)
	return &payment_gateway.GatewayRefundResponse{
		GatewayRefundID: apiResp.Data.TransactionID,
		Status:          status,
		RawResponse:     string(respBytes),
	}, nil
}

func (p *PhonePeProvider) FetchStatus(
	ctx context.Context,
	gatewayTxnID string,
	config map[string]string,
) (*payment_gateway.GatewayStatusResponse, error) {
	p.logger.Info("Polling status", slog.String("txn_id", gatewayTxnID))

	merchantID := config["merchant_id"]
	saltKey := config["salt_key"]
	saltIndex := config["salt_index"]
	baseURL := p.baseURL(config)

	endpoint := fmt.Sprintf(statusEndpoint, merchantID, gatewayTxnID)
	url := baseURL + endpoint
	checksum := p.checksum("", saltKey, saltIndex, endpoint)

	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	httpReq.Header.Set("X-VERIFY", checksum)
	httpReq.Header.Set("X-MERCHANT-ID", merchantID)

	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)
	var statusResp phonePeStatusResponse
	json.Unmarshal(respBytes, &statusResp)

	return &payment_gateway.GatewayStatusResponse{
		Status:          p.mapState(statusResp.Data.State),
		GatewayTxnID:    statusResp.Data.TransactionID,
		GatewayOrderID:  statusResp.Data.MerchantTransactionID,
		PaidAmountPaise: statusResp.Data.Amount,
		RawResponse:     string(respBytes),
	}, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// INTERNAL HELPERS (Now with full implementation and specific logs)
// ─────────────────────────────────────────────────────────────────────────────

func (p *PhonePeProvider) getToken(ctx context.Context) ([]byte, error) {

	tokenURL := "https://api-preprod.phonepe.com/apis/pg-sandbox/v1/oauth/token"

	// 1. Prepare form data (x-www-form-urlencoded)
	data := url.Values{}
	data.Set("client_id", "M23PWFUKZO053_2511101913")
	data.Set("client_version", "1")
	data.Set("client_secret", "MDFmMjQ2NGUtYzVmNy00YTc4LTkzOTQtYzdmMDAxZDY1YTIw")
	data.Set("grant_type", "client_credentials")

	// 2. Create request with the encoded string body
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	// 3. Set the specific header required for OAuth
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 4. Execute
	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (p *PhonePeProvider) signAndPost(
	ctx context.Context,
	url string,
	endpoint string,
	payload interface{},
	saltKey string,
	saltIndex string,
) ([]byte, error) {
	payloadBytes, _ := json.Marshal(payload)
	encoded := base64.StdEncoding.EncodeToString(payloadBytes)
	checksum := p.checksum(encoded, saltKey, saltIndex, endpoint)

	body, _ := json.Marshal(map[string]string{"request": encoded})
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-VERIFY", checksum)

	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		p.logger.Error("HTTP request failed", slog.String("url", url), slog.Any("err", err))
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		p.logger.Error("PhonePe API error", slog.Int("code", resp.StatusCode))
	}

	return respBytes, nil
}

func (p *PhonePeProvider) checksum(data, saltKey, saltIndex, endpoint string) string {
	// Formula: SHA256(data + endpoint + saltKey) + "###" + saltIndex
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
	// Important: Log if PhonePe sends a state we haven't mapped yet
	p.logger.Warn("Unknown PhonePe state encountered", slog.String("state", state))
	return "PROCESSING"
}

/* -------------------------------------------------------------------------- */
/*                     PhonePe V2 API Integrations Helpers                    */
/* -------------------------------------------------------------------------- */

// This is the step one here we get the tenants gateway config
// put them in the Header and the After Getting response we extract the token
func (p *PhonePeProvider) getAuthToken() {}

// Step 2 this will take the token out inside the header and then
// 
func (p *PhonePeProvider) getThePaymentURLEnpoint() {

}
