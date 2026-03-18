package restaurant_models

import (
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	"time"

	"github.com/google/uuid"
)

// ─── 1. SESSION HEADER ───────────────────────────────────────────────────────

type RestaurantDiningSession struct {
	ID           uuid.UUID                                              `gorm:"type:uuid;primaryKey"`
	RestaurantID uuid.UUID                                              `gorm:"type:uuid;not null;index"`
	TableID      uuid.UUID                                              `gorm:"type:uuid;not null;index"`
	Status       restaurant_constants.RestaurantDiningSessionStatusType `gorm:"size:20;not null"`

	// OPEN → AWAITING_PAYMENT → CLOSED / CANCELLED
	StartedAt   time.Time
	StartedBy   uuid.UUID  `gorm:"type:uuid"`
	ClosedAt    *time.Time `gorm:"type:timestamp"`
	CancelledAt *time.Time `gorm:"type:timestamp"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantDiningSession) TableName() string {
	return "hospitality.restaurant_dining_sessions"
}

// ─── 2. CUSTOMER SNAPSHOT ────────────────────────────────────────────────────
// Captured when session is opened or customer is identified mid-session
// Walk-in = CustomerID nil, name optional

type RestaurantDiningSessionCustomer struct {
	ID                        uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantDiningSessionID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"` // 1:1

	CustomerID    *uuid.UUID `gorm:"type:uuid;index"` // nil for walk-in
	CustomerName  string     `gorm:"type:varchar(200)"`
	CustomerPhone string     `gorm:"type:varchar(20)"`
	CustomerGST   *string    `gorm:"type:varchar(20)"` // for B2B GST invoices

	// Snapshotted for tax calculation — drives CGST/SGST vs IGST
	CustomerStateCode *string `gorm:"type:varchar(3)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantDiningSessionCustomer) TableName() string {
	return "hospitality.restaurant_dining_session_customers"
}

// ─── 3. PAYMENT ──────────────────────────────────────────────────────────────
// Set when CloseSession is called
// PaymentID is null until gateway confirms (online) or immediately (offline)

type RestaurantDiningSessionPayment struct {
	ID                        uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantDiningSessionID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"` // 1:1

	PaymentMethod string `gorm:"type:varchar(30);not null"` // CASH, UPI, PHONEPE, RAZORPAY

	// Payment engine refs — set progressively
	InvoiceID      *uuid.UUID `gorm:"type:uuid;index"` // set when invoice generated
	PaymentOrderID *uuid.UUID `gorm:"type:uuid;index"` // set when order created
	PaymentID      *uuid.UUID `gorm:"type:uuid;index"` // set when payment confirmed

	// For online gateways — frontend redirects here
	QRData *string `gorm:"type:text"` // ← replaces GatewayOrderRef

	// PENDING, CONFIRMED, FAILED
	Status string `gorm:"type:varchar(20);not null;default:'PENDING'"`

	ConfirmedAt *time.Time `gorm:"type:timestamp"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (RestaurantDiningSessionPayment) TableName() string {
	return "hospitality.restaurant_dining_session_payments"
}

// hospitality.restaurant_dining_sessions            ← session lifecycle + timestamps
// hospitality.restaurant_dining_session_customers   ← customer snapshot (1:1)
// hospitality.restaurant_dining_session_payments    ← payment engine refs (1:1)
// hospitality.restaurant_orders                     ← orders placed during session (1:many)
// hospitality.restaurant_order_items                ← items per order
// hospitality.restaurant_order_items_variants       ← variants per item
// hospitality.restaurant_order_items_addons         ← addons per item
