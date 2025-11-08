package accounting_models

import (
  "github.com/google/uuid"
  "time"
)

type Account struct {
  ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  BookID      uuid.UUID `gorm:"type:uuid;index;not null"`
  Name        string    `gorm:"not null"`
  Code        string    `gorm:"uniqueIndex;not null"`
  GroupID     uuid.UUID `gorm:"type:uuid;index"`
  Type        string
  OpeningBal  float64
  CurrentBal  float64
  IsDefault   bool
  Status      string
  CreatedAt   time.Time
  UpdatedAt   time.Time
}
