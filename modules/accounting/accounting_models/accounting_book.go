package accounting_models

import (
  "github.com/google/uuid"
  "time"
)

type AccountingBook struct {
  ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  OrgID        uuid.UUID `gorm:"type:uuid;not null;index"`
  Name         string     `gorm:"not null"`
  Department   string
  FiscalYearID uuid.UUID  `gorm:"type:uuid"`
  Status       string
  CreatedAt    time.Time
  UpdatedAt    time.Time
}
