// PhonePe-specific request/response types.
// NEVER exposed outside this package.
// The rest of the engine only sees the shared gateway DTOs from provider.go.
//
// PHONEPE FLOW FOR COUNTER (UPI QR):
//   1. Your server sends a PAY_PAGE or UPI_QR request to PhonePe
//   2. PhonePe returns a QR string
//   3. You show QR on second screen
//   4. Customer scans with any UPI app
//   5. PhonePe POSTs webhook to your registered WebhookURL
//   6. You verify HMAC, update payment, close session
//
// CHECKSUM ALGORITHM:
//   SHA256(base64(payload) + "/" + endpoint + saltKey) + "###" + saltIndex
//
// DOCS: https://developer.phonepe.com/v1/reference/pay-api

package phonepe

// ─── OUTBOUND (your server → PhonePe) ────────────────────────────────────────

type phonePePayRequest struct {
	MerchantID            string            `json:"merchantId"`
	MerchantTransactionID string            `json:"merchantTransactionId"` // = your Payment.ID
	Amount                int64             `json:"amount"`                // paise
	CallbackURL           string            `json:"callbackUrl"`           // from config["webhook_url"]
	PaymentInstrument     phonePeInstrument `json:"paymentInstrument"`
	MobileNumber          string            `json:"mobileNumber,omitempty"`
}

type phonePeInstrument struct {
	// UPI_QR for counter flow — customer scans QR
	// PAY_PAGE for redirect flow (not used in counter app)
	Type string `json:"type"`
}

type phonePeRefundRequest struct {
	MerchantID            string `json:"merchantId"`
	MerchantTransactionID string `json:"merchantTransactionId"` // = your Refund.ID
	OriginalTransactionID string `json:"originalTransactionId"` // = Payment.GatewayTxnRef
	Amount                int64  `json:"amount"`
	CallbackURL           string `json:"callbackUrl"`
}

// ─── INBOUND (PhonePe → your server) ─────────────────────────────────────────

type phonePeAPIResponse struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"` // PAYMENT_INITIATED, PAYMENT_SUCCESS, PAYMENT_ERROR etc.
	Message string      `json:"message"`
	Data    phonePeData `json:"data"`
}

type phonePeData struct {
	MerchantID            string                     `json:"merchantId"`
	MerchantTransactionID string                     `json:"merchantTransactionId"`
	TransactionID         string                     `json:"transactionId"` // PhonePe's own TxnID
	Amount                int64                      `json:"amount"`
	State                 string                     `json:"state"` // PENDING, COMPLETED, FAILED
	InstrumentResponse    *phonePeInstrumentResponse `json:"instrumentResponse,omitempty"`
}

type phonePeInstrumentResponse struct {
	Type   string  `json:"type"`
	QRData *string `json:"qrData,omitempty"` // UPI QR string — shown on second screen
}

// phonePeWebhookPayload — raw POST body from PhonePe webhook
type phonePeWebhookPayload struct {
	Response string `json:"response"` // base64 encoded JSON of phonePeData
}

// phonePeStatusResponse — response from status check API
type phonePeStatusResponse struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    phonePeData `json:"data"`
}
