package accounting_models

import "github.com/google/uuid"

type BillItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	BillID    uuid.UUID
	ItemName  string
	Qty       float64
	Rate      float64
	TaxRateID uuid.UUID
	Amount    float64
}
