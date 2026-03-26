package phonepe

// ─── Auth ────────────────────────────────────────────────────────────────────

type authTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// ─── Create Payment ───────────────────────────────────────────────────────────

type createPaymentRequest struct {
	MerchantOrderID string      `json:"merchantOrderId"`
	Amount          int64       `json:"amount"`       // paise
	ExpireAfter     int         `json:"expireAfter"`  // seconds
	MetaInfo        map[string]string `json:"metaInfo,omitempty"`
	PaymentFlow     paymentFlow `json:"paymentFlow"`
}

type paymentFlow struct {
	Type         string       `json:"type"` // "PG_CHECKOUT"
	Message      string       `json:"message,omitempty"`
	MerchantURLs *merchantURLs `json:"merchantUrls,omitempty"`
}

type merchantURLs struct {
	RedirectURL string `json:"redirectUrl,omitempty"`
}

type createPaymentResponse struct {
	OrderID     string `json:"orderId"`
	State       string `json:"state"`       // "PENDING"
	RedirectURL string `json:"redirectUrl"` // checkout page URL
	ExpireAt    int64  `json:"expireAt"`    // unix timestamp
}

// ─── Order Status ─────────────────────────────────────────────────────────────

type orderStatusResponse struct {
	OrderID        string         `json:"orderId"`
	MerchantOrderID string        `json:"merchantOrderId"`
	State          string         `json:"state"`    // "COMPLETED", "FAILED", "PENDING"
	Amount         int64          `json:"amount"`
	PaymentDetails []paymentDetail `json:"paymentDetails,omitempty"`
}

type paymentDetail struct {
	TransactionID string `json:"transactionId"`
	PaymentMode   string `json:"paymentMode"`
	ErrorCode     string `json:"errorCode,omitempty"`
	State         string `json:"state"`
}

// ─── Webhook ──────────────────────────────────────────────────────────────────

type webhookPayload struct {
	OrderID        string         `json:"orderId"`
	MerchantOrderID string        `json:"merchantOrderId"` // your internal payment ID
	State          string         `json:"state"`
	PaymentDetails []paymentDetail `json:"paymentDetails,omitempty"`
}

// ─── Refund ───────────────────────────────────────────────────────────────────

type refundRequest struct {
	MerchantRefundID string `json:"merchantRefundId"`
	OriginalOrderID  string `json:"originalMerchantOrderId"` // your internal payment ID
	Amount           int64  `json:"amount"`
}

type refundResponse struct {
	RefundID string `json:"refundId"`
	State    string `json:"state"` // "COMPLETED", "PENDING", "FAILED"
}