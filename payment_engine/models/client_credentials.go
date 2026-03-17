package payment_models

import (
	"time"

	"github.com/google/uuid"
)

type OrgPaymentGetawayCredentials struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OrgID          uuid.UUID `gorm:"type:uuid;index" json:"org_id"`
	Provider       string    `gorm:"size:20;index" json:"provider"`
	MerchantID     string    `gorm:"size:120" json:"merchant_id"`
	MerchantKey    string    `gorm:"type:text" json:"merchant_key"`    // encrypted
	MerchantSecret string    `gorm:"type:text" json:"merchant_secret"` // encrypted
	SaltKey        string    `gorm:"type:text" json:"salt_key"`        // encrypted (PhonePe)
	SaltIndex      string    `gorm:"size:10" json:"salt_index"`        // PhonePe salt index
	IsActive       bool      `gorm:"default:true" json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (OrgPaymentGetawayCredentials) TableName() string {
	return "payments.org_payment_gataway_credentials"
}
