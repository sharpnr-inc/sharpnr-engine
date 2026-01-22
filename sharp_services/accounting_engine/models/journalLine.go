package accounting_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type JournalLine struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	JournalEntryID uuid.UUID    `gorm:"type:uuid;not null;index"`
	JournalEntry   JournalEntry `gorm:"foreignKey:JournalEntryID"`

	AccountID uuid.UUID `gorm:"type:uuid;not null;index"`
	Account   Account   `gorm:"foreignKey:AccountID"`

	Debit  float64 `gorm:"default:0"`
	Credit float64 `gorm:"default:0"`

	// Optional but VERY useful
	PartyID *uuid.UUID     // customer / supplier
	Meta    datatypes.JSON // room_id, stay_id, order_id

	CreatedAt time.Time
}

func (JournalLine) TableName() string {
	return "accounting_engine.journal_line"
}
