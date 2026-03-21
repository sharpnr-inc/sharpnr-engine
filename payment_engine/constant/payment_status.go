package payment_constants

// ─── Payment Status ───────────────────────────────────────────────────────────

const (
	StatusInitiated         = "INITIATED"
	StatusProcessing        = "PROCESSING"
	StatusSuccess           = "SUCCESS"
	StatusFailed            = "FAILED"
	StatusRefunded          = "REFUNDED"
	StatusPartiallyRefunded = "PARTIALLY_REFUNDED"
	StatusCancelled         = "CANCELLED"
)

// ─── Order Status ─────────────────────────────────────────────────────────────

const (
	OrderStatusPending       = "PENDING"
	OrderStatusPartiallyPaid = "PARTIALLY_PAID"
	OrderStatusPaid          = "PAID"
	OrderStatusExpired       = "EXPIRED"
	OrderStatusCancelled     = "CANCELLED"
)

// ─── Refund Status ────────────────────────────────────────────────────────────

const (
	RefundStatusInitiated  = "INITIATED"
	RefundStatusProcessing = "PROCESSING"
	RefundStatusSuccess    = "SUCCESS"
	RefundStatusFailed     = "FAILED"
)

// ─── Payment Method ───────────────────────────────────────────────────────────

const (
	MethodCash         = "CASH"
	MethodCheque       = "CHEQUE"
	MethodBankTransfer = "BANK_TRANSFER"
	MethodUPI          = "UPI"
	MethodCard         = "CARD"
	MethodWallet       = "WALLET"
	MethodCreditNote   = "CREDIT_NOTE"
	MethodPhonePe      = "PHONEPE"
	MethodRazorpay     = "RAZORPAY"
	MethodCashfree     = "CASHFREE"
)

// ─── Payment Source ───────────────────────────────────────────────────────────
// Where the payment order originated — set by the caller, opaque to engine

const (
	SourceInvoice     = "INVOICE"
	SourcePOS         = "POS"
	SourceAdvance     = "ADVANCE"
	SourceWalletTopup = "WALLET_TOPUP"
)

// ─── Payer Type ───────────────────────────────────────────────────────────────

const (
	PayerTypeCustomer = "CUSTOMER"
	PayerTypeGuest    = "GUEST"
	PayerTypeWalkIn   = "WALK_IN"
	PayerTypeVendor   = "VENDOR"
)

// ─── Allocation Target ────────────────────────────────────────────────────────
// TargetType values used in PaymentAllocation
// Engine stores these as opaque strings — never queries the target tables

const (
	AllocationTargetInvoice    = "INVOICE"
	AllocationTargetAdvance    = "ADVANCE"
	AllocationTargetCreditNote = "CREDIT_NOTE"
)

// ─── Allocation Status ────────────────────────────────────────────────────────

const (
	AllocationStatusActive   = "ACTIVE"
	AllocationStatusReversed = "REVERSED"
)

// ─── Gateway Provider Codes ───────────────────────────────────────────────────
// Must match the Code() return value on each GatewayProvider implementation

const (
	ProviderPhonePe  = "PHONEPE"
	ProviderRazorpay = "RAZORPAY"
	ProviderCashfree = "CASHFREE"
	ProviderCash     = "CASH"
)

// ─── Webhook Events ───────────────────────────────────────────────────────────
// EventType values used in WebhookDeliveryLog and SubscribedEvents

const (
	EventPaymentInitiated  = "PAYMENT_INITIATED"
	EventPaymentSuccess    = "PAYMENT_SUCCESS"
	EventPaymentFailed     = "PAYMENT_FAILED"
	EventPaymentProcessing = "PAYMENT_PROCESSING"
	EventPaymentCancelled  = "PAYMENT_CANCELLED"

	EventRefundInitiated  = "REFUND_INITIATED"
	EventRefundSuccess    = "REFUND_SUCCESS"
	EventRefundFailed     = "REFUND_FAILED"
	EventRefundProcessing = "REFUND_PROCESSING"

	EventOrderCreated   = "ORDER_CREATED"
	EventOrderPaid      = "ORDER_PAID"
	EventOrderExpired   = "ORDER_EXPIRED"
	EventOrderCancelled = "ORDER_CANCELLED"

	EventPaymentAllocated   = "PAYMENT_ALLOCATED"
	EventAllocationReversed = "ALLOCATION_REVERSED"
)

// ─── Webhook Delivery Status ──────────────────────────────────────────────────

const (
	WebhookDeliveryPending = "PENDING"
	WebhookDeliverySuccess = "SUCCESS"
	WebhookDeliveryFailed  = "FAILED"
)

// ─── Max Retry Attempts ───────────────────────────────────────────────────────

const WebhookMaxRetryAttempts = 5
