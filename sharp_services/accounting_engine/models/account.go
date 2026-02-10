package accounting_models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	AccountGroupID uuid.UUID `gorm:"type:uuid;index"`

	Name string `gorm:"not null"`
	Code string `gorm:"not null"`

	OpeningBalance float64 `gorm:"default:0"`
	OpeningIsDebit bool    `gorm:"default:true"`
	Status         string  `gorm:"size:20;default:'ACTIVE'"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Account) TableName() string {
	return "accounting_engine.account"
}
