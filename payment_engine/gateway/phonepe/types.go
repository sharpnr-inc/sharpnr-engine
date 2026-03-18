package phonepe

// -----------------------------------------------
// PHONEPE-SPECIFIC TYPES
// Never exposed outside this package
// -----------------------------------------------

type phonePePayRequest struct {
	MerchantID            string            `json:"merchantId"`
	MerchantTransactionID string            `json:"merchantTransactionId"`
	MerchantUserID        string            `json:"merchantUserId,omitempty"`
	Amount                int64             `json:"amount"` // in paise
	RedirectURL           string            `json:"redirectUrl,omitempty"`
	RedirectMode          string            `json:"redirectMode,omitempty"`
	CallbackURL           string            `json:"callbackUrl,omitempty"`
	MobileNumber          string            `json:"mobileNumber,omitempty"`
	PaymentInstrument     phonePeInstrument `json:"paymentInstrument"`
}

type phonePeInstrument struct {
	Type string `json:"type"` // PAY_PAGE, UPI_COLLECT, UPI_QR
}

type phonePeResponse struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    phonePeData `json:"data"`
}

type phonePeData struct {
	MerchantID            string                     `json:"merchantId"`
	MerchantTransactionID string                     `json:"merchantTransactionId"`
	TransactionID         string                     `json:"transactionId"`
	Amount                int64                      `json:"amount"`
	State                 string                     `json:"state"` // PENDING, COMPLETED, FAILED
	InstrumentResponse    *phonePeInstrumentResponse `json:"instrumentResponse,omitempty"`
}

type phonePeInstrumentResponse struct {
	Type         string           `json:"type"`
	RedirectInfo *phonePeRedirect `json:"redirectInfo,omitempty"`
	QRData       *string          `json:"qrData,omitempty"`
}

type phonePeRedirect struct {
	URL    string `json:"url"`
	Method string `json:"method"` // GET / POST
}

type phonePeCallbackPayload struct {
	Response string `json:"response"` // base64 encoded JSON
}

type phonePeStatusResponse struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    phonePeData `json:"data"`
}

type phonePeRefundRequest struct {
	MerchantID            string `json:"merchantId"`
	MerchantTransactionID string `json:"merchantTransactionId"` // your internal refund ID
	OriginalTransactionID string `json:"originalTransactionId"` // gateway txn ID being refunded
	Amount                int64  `json:"amount"`
	CallbackURL           string `json:"callbackUrl"`
}
