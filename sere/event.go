package sharpnreventbus

import (
	"time"

	"github.com/google/uuid"
)

// SharpEvent represents a generic event published by any module.
type SharpEvent struct {
	ID             uuid.UUID   `json:"id"`
	Name           string      `json:"name"`    // e.g. "EventPaymentReceived"
	Source         string      `json:"source"`  // e.g. "Accounting", "Hospitality"
	Payload        interface{} `json:"payload"` // Can be any struct (Invoice, Payment, etc.)
	Timestamp      time.Time   `json:"timestamp"`
	OrganizationID uuid.UUID   `json:"organization_id"`
}
