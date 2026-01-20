package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type AccountMapping struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrgID uuid.UUID `gorm:"type:uuid;not null;index"`

	Module string `gorm:"size:50;not null"`
	// HOSPITALITY, HEALTHCARE, MANUFACTURING

	SubModule string `gorm:"size:50;not null"`
	// HOTEL, BAR, RESTAURANT

	SourceType string `gorm:"size:50;not null"`
	// CHARGE, TAX, PAYMENT, SYSTEM

	SourceCode string `gorm:"size:100;not null"`
	// ROOM_RENT, FOOD, CGST, CASH, GUEST_RECEIVABLE

	AccountID uuid.UUID `gorm:"type:uuid;not null"`
	Account   Account   `gorm:"foreignKey:AccountID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
