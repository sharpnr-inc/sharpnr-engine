package payment_engine_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID         uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityID    uuid.UUID `gorm:"type:uuid;not null;index"`
	PaymentAccountID uuid.UUID `gorm:"type:uuid;not null;index"`
	PaymentOrderID   uuid.UUID `gorm:"type:uuid;not null;index"`
	PaymentNumber    string    `gorm:"type:varchar(50);not null;uniqueIndex:idx_pay_number_tenant"`

	CurrencyCode string  `gorm:"type:char(3);not null;default:'INR'"`
	AmountPaise  int64   `gorm:"not null"`
	Amount       float64 `gorm:"-"` // derived

	// CASH, UPI, CARD, BANK_TRANSFER, CHEQUE, WALLET, CREDIT_NOTE
	Method string `gorm:"type:varchar(30);not null;index"`

	// INITIATED, PROCESSING, SUCCESS, FAILED, REFUNDED, PARTIALLY_REFUNDED, CANCELLED
	Status string `gorm:"type:varchar(30);not null;default:'INITIATED';index"`

	// Gateway linkage — null for offline methods (Cash, Cheque)
	GatewayConfigID *uuid.UUID `gorm:"type:uuid;index"`
	GatewayOrderRef *string    `gorm:"type:varchar(200)"`       // gateway's order ID
	GatewayTxnRef   *string    `gorm:"type:varchar(200);index"` // gateway's txn ID
	GatewayResponse *string    `gorm:"type:text"`               // raw JSON snapshot — never update after set

	// ← These three were missing
	QRCode    *string    `gorm:"type:text"`      // "upi://pay?pa=..." — passed to frontend QR renderer
	ExpiresAt *time.Time `gorm:"type:timestamp"` // when this gateway session expires — drives countdown timer

	// Offline/manual methods
	ReferenceNumber *string    `gorm:"type:varchar(100)"` // UTR, UPI VPA, cheque no
	PaidAt          *time.Time `gorm:"type:timestamp"`    // actual settlement timestamp

	// Cached refund total — updated by refund service
	RefundedPaise  int64   `gorm:"default:0"`
	RefundedAmount float64 `gorm:"-"` // derived

	Notes     *string `gorm:"type:text"`
	CreatedBy string  `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Payment) TableName() string { return "payments.payments" }
