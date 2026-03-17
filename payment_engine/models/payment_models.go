package payment_models

import (
	"time"

	"github.com/google/uuid"
)

type Payments struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OrgID        uuid.UUID `gorm:"type:uuid;index" json:"org_id"`
	CustomerUPI  string    `gorm:"size:100;index" json:"customer_upi"`
	AmountPaise  int64     `json:"amount_paise"`                            // store in paise (e.g., 45600)
	Amount       float64   `gorm:"-" json:"amount"`                         // derived for responses (₹)
	ReferenceID  string    `gorm:"size:80;uniqueIndex" json:"reference_id"` // your txn id
	Provider     string    `gorm:"size:20;index" json:"provider"`           // phonepe | paytm
	Status       string    `gorm:"size:20;index" json:"status"`             // pending | success | failed | cancelled
	RequestData  string    `gorm:"type:text" json:"request_data"`           // raw request JSON
	ResponseData string    `gorm:"type:text" json:"response_data"`          // raw response JSON / callback
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Payments) TableName() string {
	return "payments.payments"
}
