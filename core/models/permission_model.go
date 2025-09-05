package models

import (
	"time"

	"github.com/google/uuid"
)

type Permision struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"size:50;not null;unique"` // e.g., CREATE_INVOICE
	Description string
	Module      string // optional grouping like Accounting, HR
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
