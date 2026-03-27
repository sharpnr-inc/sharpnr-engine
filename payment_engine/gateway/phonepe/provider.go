package phonepe

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	payment_gateway "sharpnr-engine/payment_engine/gateway"
)

// --- Logging Colors ---
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
)

const (
	uatBaseURL  = "https://api-preprod.phonepe.com/apis/pg-sandbox"
	prodBaseURL = "https://api.phonepe.com/apis/pg"

	uatAuthURL  = "https://api-preprod.phonepe.com/apis/pg-sandbox/v1/oauth/token"
	prodAuthURL = "https://api.phonepe.com/apis/identity-manager/v1/oauth/token"

	checkoutEndpoint    = "/checkout/v2/pay"
	orderStatusEndpoint = "/checkout/v2/order/%s/status"
	refundEndpoint      = "/checkout/v2/refund"
)

// --- Provider Implementation ---

type PhonePeProvider struct {
	httpClient *http.Client
}

func NewPhonePeProvider() payment_gateway.GatewayProvider {
	return &PhonePeProvider{
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (p *PhonePeProvider) Code() string { return "PHONEPE" }

// CreateOrder handles the initiation of a payment
func (p *PhonePeProvider) CreateOrder(
	ctx context.Context,
	req payment_gateway.GatewayOrderRequest,
) (*payment_gateway.GatewayOrderResponse, error) {
	log.Printf("%s[PHONEPE-INIT]%s Starting order creation for InternalID: %s%s%s", blue, reset, yellow, req.InternalOrderID, reset)

	isTest := req.Config["is_test_mode"] == "true"

	// 1. Get OAuth token
	log.Printf("%s[STEP 1]%s Authenticating with PhonePe...", cyan, reset)
	token, err := p.getAuthToken(ctx, req.Config, isTest)
	if err != nil {
		log.Printf("%s[FATAL]%s Auth failed: %v", red, reset, err)
		return nil, fmt.Errorf("phonepe: get auth token: %w", err)
	}
	log.Printf("%s[SUCCESS]%s Token acquired", green, reset)

	// 2. Build payload
	payload := createPaymentRequest{
		MerchantOrderID: req.InternalOrderID,
		Amount:          req.AmountPaise,
		ExpireAfter:     1200,
		MetaInfo:        req.Meta,
		PaymentFlow: paymentFlow{
			Type:    "PG_CHECKOUT",
			Message: "Payment",
		},
	}
	log.Printf("%s[STEP 2]%s Prepared payload: %s%+v%s", cyan, reset, gray, payload, reset)

	// 3. Execution
	baseURL := p.resolveBaseURL(isTest)
	log.Printf("%s[STEP 3]%s Dispatching request to PhonePe API...", cyan, reset)
	respBytes, err := p.doRequest(ctx, http.MethodPost, baseURL+checkoutEndpoint, payload, token, req.Config["merchant_id"])
	if err != nil {
		log.Printf("%s[ERROR]%s API Request Failed: %v", red, reset, err)
		return nil, fmt.Errorf("phonepe: create order: %w", err)
	}

	var phonePeResp createPaymentResponse
	if err := json.Unmarshal(respBytes, &phonePeResp); err != nil {
		log.Printf("%s[ERROR]%s Response unmarshal failed: %v", red, reset, err)
		return nil, fmt.Errorf("phonepe: unmarshal create order response: %w", err)
	}

	log.Printf("%s[COMPLETED]%s Order created. PhonePe OrderID: %s%s%s", green, reset, yellow, phonePeResp.OrderID, reset)

	result := &payment_gateway.GatewayOrderResponse{
		ProviderOrderID: phonePeResp.OrderID,
		RawResponse:     string(respBytes),
	}

	if phonePeResp.RedirectURL != "" {
		log.Printf("%s[INFO]%s QR/Redirect URL extracted: %s", gray, reset, phonePeResp.RedirectURL)
		result.QRData = &phonePeResp.RedirectURL
	}

	if phonePeResp.ExpireAt > 0 {
		t := time.Unix(phonePeResp.ExpireAt, 0)
		result.ExpiresAt = &t
	}

	return result, nil
}

/* ---------------------------------- <***> --------------------------------- */

// VerifyPayment handles incoming webhooks
func (p *PhonePeProvider) VerifyPayment(
	ctx context.Context,
	req payment_gateway.GatewayVerifyRequest,
) (*payment_gateway.GatewayVerifyResponse, error) {
	log.Printf("%s[WEBHOOK-RECEIVED]%s Processing verification...", purple, reset)

	webhookUsername := req.Config["webhook_username"]
	webhookPassword := req.Config["webhook_password"]

	// 1. Verify Authorization header
	incomingAuth := req.Headers["authorization"]
	expectedAuth := p.generateWebhookSignature(webhookUsername, webhookPassword)

	log.Printf("%s[AUTH-CHECK]%s Comparing signatures...", cyan, reset)
	if incomingAuth != expectedAuth {
		log.Printf("%s[SECURITY ALERT]%s Invalid webhook signature!", red, reset)
		return nil, fmt.Errorf("phonepe: invalid webhook signature")
	}
	log.Printf("%s[AUTH-OK]%s Signature verified", green, reset)

	// 2. Parse payload
	var webhookData webhookPayload
	if err := json.Unmarshal([]byte(req.RawCallbackBody), &webhookData); err != nil {
		log.Printf("%s[ERROR]%s Webhook data corrupted: %v", red, reset, err)
		return nil, fmt.Errorf("phonepe: unmarshal webhook payload: %w", err)
	}

	status := p.mapState(webhookData.State)
	log.Printf("%s[STATUS]%s Order %s set to %s%s%s", cyan, reset, webhookData.MerchantOrderID, yellow, status, reset)

	gatewayTxnID := webhookData.OrderID
	if len(webhookData.PaymentDetails) > 0 {
		gatewayTxnID = webhookData.PaymentDetails[0].TransactionID
	}

	return &payment_gateway.GatewayVerifyResponse{
		Status:         status,
		GatewayTxnID:   gatewayTxnID,
		GatewayOrderID: webhookData.MerchantOrderID,
		PaidAt:         time.Now(),
		RawResponse:    req.RawCallbackBody,
	}, nil
}

// InitiateRefund processes a refund request
func (p *PhonePeProvider) InitiateRefund(
	ctx context.Context,
	req payment_gateway.GatewayRefundRequest,
) (*payment_gateway.GatewayRefundResponse, error) {
	log.Printf("%s[REFUND-START]%s InternalID: %s", blue, reset, req.InternalRefundID)

	isTest := req.Config["is_test_mode"] == "true"
	token, err := p.getAuthToken(ctx, req.Config, isTest)
	if err != nil {
		return nil, err
	}

	payload := refundRequest{
		MerchantRefundID: req.InternalRefundID,
		OriginalOrderID:  req.GatewayTxnID,
		Amount:           req.RefundAmountPaise,
	}

	baseURL := p.resolveBaseURL(isTest)
	respBytes, err := p.doRequest(ctx, http.MethodPost, baseURL+refundEndpoint, payload, token, req.Config["merchant_id"])
	if err != nil {
		log.Printf("%s[REFUND-FAILED]%s %v", red, reset, err)
		return nil, err
	}

	var refundResp refundResponse
	json.Unmarshal(respBytes, &refundResp)

	log.Printf("%s[REFUND-SUCCESS]%s PhonePe RefundID: %s", green, reset, refundResp.RefundID)

	return &payment_gateway.GatewayRefundResponse{
		GatewayRefundID: refundResp.RefundID,
		Status:          p.mapState(refundResp.State),
		RawResponse:     string(respBytes),
	}, nil
}

// FetchStatus reconciles status via polling
func (p *PhonePeProvider) FetchStatus(
	ctx context.Context,
	gatewayTxnID string,
	config map[string]string,
) (*payment_gateway.GatewayStatusResponse, error) {
	log.Printf("%s[FETCH-STATUS]%s Polling for Txn: %s", cyan, reset, gatewayTxnID)

	isTest := config["is_test_mode"] == "true"
	token, err := p.getAuthToken(ctx, config, isTest)
	if err != nil {
		return nil, err
	}

	baseURL := p.resolveBaseURL(isTest)
	url := baseURL + fmt.Sprintf(orderStatusEndpoint, gatewayTxnID)

	log.Printf("%s[GET]%s %s", purple, reset, url)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header.Set("Authorization", "O-Bearer "+token)
	req.Header.Set("X-MERCHANT-ID", config["merchant_id"])

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)
	var statusResp orderStatusResponse
	json.Unmarshal(respBytes, &statusResp)

	log.Printf("%s[RESULT]%s PhonePe State: %s%s%s", cyan, reset, yellow, statusResp.State, reset)

	gatewayTxnRef := statusResp.OrderID
	if len(statusResp.PaymentDetails) > 0 {
		gatewayTxnRef = statusResp.PaymentDetails[0].TransactionID
	}

	return &payment_gateway.GatewayStatusResponse{
		Status:          p.mapState(statusResp.State),
		GatewayTxnID:    gatewayTxnRef,
		GatewayOrderID:  statusResp.MerchantOrderID,
		PaidAmountPaise: statusResp.Amount,
		RawResponse:     string(respBytes),
	}, nil
}

// --- Helpers with Logging ---

func (p *PhonePeProvider) getAuthToken(ctx context.Context, config map[string]string, isTest bool) (string, error) {
	authURL := prodAuthURL
	if isTest {
		authURL = uatAuthURL
	}

	url := fmt.Sprintf(
		"%s?client_id=%s&client_version=%s&client_secret=%s&grant_type=client_credentials",
		authURL, config["client_id"], config["client_version"], config["client_secret"],
	)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)
	var tokenResp authTokenResponse
	if err := json.Unmarshal(respBytes, &tokenResp); err != nil {
		return "", err
	}

	if tokenResp.AccessToken == "" {
		log.Printf("%s[AUTH-FAIL]%s No token in response", red, reset)
		return "", fmt.Errorf("empty access token")
	}

	return tokenResp.AccessToken, nil
}

func (p *PhonePeProvider) doRequest(ctx context.Context, method, url string, payload interface{}, token, mID string) ([]byte, error) {
	log.Printf("%s[HTTP-REQ]%s %s %s", purple, reset, method, url)

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "O-Bearer "+token)
	req.Header.Set("X-MERCHANT-ID", mID)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		log.Printf("%s[API-WARN]%s Status: %d | Body: %s", yellow, reset, resp.StatusCode, string(respBytes))
		return respBytes, fmt.Errorf("phonepe error: %d", resp.StatusCode)
	}

	return respBytes, nil
}

func (p *PhonePeProvider) resolveBaseURL(isTest bool) string {
	if isTest {
		return uatBaseURL
	}
	return prodBaseURL
}

func (p *PhonePeProvider) generateWebhookSignature(username, password string) string {
	input := fmt.Sprintf("%s:%s", username, password)
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}

func (p *PhonePeProvider) mapState(state string) string {
	switch state {
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
