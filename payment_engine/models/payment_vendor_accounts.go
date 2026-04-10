package payment_engine_models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type PaymentVendorAccount struct {
	ID               uuid.UUID
	OrganizationID   uuid.UUID // TenantID we call the Organization
	ProviderVendorID string    `gorm:"type:varchar(200);not null"`
	VendorStatus     string    `gorm:"type:varchar(30);not null;default:'PENDING'"`

	// Normalized KYC across all providers
	// PENDING, IN_REVIEW, VERIFIED, REJECTED, ON_HOLD
	KYCStatus  string  `gorm:"type:varchar(30);not null;default:'PENDING'"`
	KYCRemarks *string `gorm:"type:text"`

	// Provider-specific data — shape differs per provider
	// Cashfree:  { "schedule_option": 1, "bank_account_no": "enc:...", "ifsc": "HDFC...", "pan": "enc:...", "account_type": "BUSINESS", "doc_statuses": {"PAN": "VERIFIED"} }
	// Razorpay:  { "linked_account_id": "acc_xxx", "stakeholder_id": "sth_xxx", "dashboard_access": false }
	// PhonePe:   { "merchant_id": "M123", "merchant_group_id": "GRP1", "settlement_mode": "BANK" }
	// Pinelabs:  { "terminal_id": "T001", "store_id": "S001", "merchant_code": "MC1", "device_serial": "PL123" }
	ProviderData json.RawMessage `gorm:"type:jsonb"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
