package restaurant_models

import (
	"time"

	"github.com/google/uuid"
)

// ─── 1. MAIN ORDER ───────────────────────────────────────────────────────────

type RestaurantPickUpOrders struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	RestaurantID   uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityID  uuid.UUID `gorm:"type:uuid;not null;index"`

	OrderNumber string `gorm:"type:varchar(50);not null;uniqueIndex"`

	// TAKEAWAY, DELIVERY
	Type string `gorm:"type:varchar(20);not null;default:'TAKEAWAY'"`

	// PENDING_PAYMENT, CONFIRMED, PREPARING, READY, COMPLETED, CANCELLED
	Status string `gorm:"type:varchar(30);not null;default:'PENDING_PAYMENT';index"`

	IsPaid bool `gorm:"default:false"`

	// Timestamps for each state transition — useful for kitchen display + analytics
	ConfirmedAt *time.Time `gorm:"type:timestamp"`
	PreparedAt  *time.Time `gorm:"type:timestamp"`
	ReadyAt     *time.Time `gorm:"type:timestamp"`
	CompletedAt *time.Time `gorm:"type:timestamp"`
	CancelledAt *time.Time `gorm:"type:timestamp"`

	CreatedBy       string    `gorm:"type:varchar(100)"`
	CreatedByEmplID uuid.UUID `gorm:"type:uuid"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (RestaurantPickUpOrders) TableName() string {
	return "hospitality.restaurant_pickup_orders"
}

/* ---------------------------------- <***> --------------------------------- */

type RestaurantPickUpOrderPayment struct {
	ID                      uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantPickUpOrderID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"` // 1:1

	PaymentMethod string `gorm:"type:varchar(30);not null"` // CASH, PHONEPE, RAZORPAY

	// Set when payment engine confirms
	InvoiceID      *uuid.UUID `gorm:"type:uuid;index"`
	PaymentOrderID *uuid.UUID `gorm:"type:uuid;index"`
	PaymentID      *uuid.UUID `gorm:"type:uuid;index"` // ← actual payment record

	// Opaque refs for online gateways
	QRData *string `gorm:"type:text"` // ← replaces GatewayOrderRef
	Status string  `gorm:"type:varchar(30);not null;default:'PENDING'"`

	CreatedAt   time.Time
	ConfirmedAt *time.Time `gorm:"type:timestamp"`
}

func (RestaurantPickUpOrderPayment) TableName() string {
	return "hospitality.restaurant_pickup_order_payments"
}

/* ---------------------------------- <***> --------------------------------- */

type RestaurantPickUpOrderCustomer struct {
	ID                      uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantPickUpOrderID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"` // 1:1

	// FK to customer table — nullable for walk-ins
	CustomerID *uuid.UUID `gorm:"type:uuid;index"`

	// Snapshot
	CustomerName  string  `gorm:"type:varchar(200);not null"`
	CustomerPhone string  `gorm:"type:varchar(20);not null"`
	CustomerEmail *string `gorm:"type:varchar(200)"`

	// Delivery only — null for takeaway
	DeliveryAddress   *string  `gorm:"type:text"`
	DeliveryNote      *string  `gorm:"type:varchar(500)"`
	DeliveryLatitude  *float64 `gorm:"type:decimal(10,8)"`
	DeliveryLongitude *float64 `gorm:"type:decimal(11,8)"`

	CreatedAt time.Time
}

func (RestaurantPickUpOrderCustomer) TableName() string {
	return "hospitality.restaurant_pickup_order_customers"
}
