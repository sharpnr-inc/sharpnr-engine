package accounting_models

import (
  "github.com/google/uuid"
  "time"
)

type FiscalYear struct {
  ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  OrgID      uuid.UUID
  Name       string
  StartDate  time.Time
  EndDate    time.Time
  Status     string
  CreatedAt  time.Time
  UpdatedAt  time.Time
}
