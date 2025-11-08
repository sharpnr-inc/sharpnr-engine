package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID        uuid.UUID
	BookID       uuid.UUID
	VoucherID    uuid.UUID
	AccountID    uuid.UUID
	Date         time.Time
	Type         string
	Amount       float64
	BalanceAfter float64
}
