package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type Voucher struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID       uuid.UUID
	BookID      uuid.UUID
	VoucherType string
	VoucherNo   string
	VoucherDate time.Time
	Narration   string
	CreatedBy   uuid.UUID
	CreatedAt   time.Time
}
