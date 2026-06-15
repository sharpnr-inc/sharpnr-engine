package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type JournalLine struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey"`
	JournalEntryID uuid.UUID  `gorm:"type:uuid;not null;index"`
	AccountID      uuid.UUID  `gorm:"type:uuid;not null;index"` // this is the
	Debit          float64    `gorm:"default:0"`
	Credit         float64    `gorm:"default:0"`
	PartyID        *uuid.UUID // customer / supplier
	CreatedAt time.Time
}

func (JournalLine) TableName() string {
	return "accounting_engine.journal_line"
}
