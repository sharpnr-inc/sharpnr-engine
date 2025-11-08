package accounting_models

import "github.com/google/uuid"

type VoucherEntry struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	VoucherID   uuid.UUID `gorm:"type:uuid;index"`
	AccountID   uuid.UUID `gorm:"type:uuid;index"`
	Type        string    // DR or CR
	Amount      float64
	Description string
}
