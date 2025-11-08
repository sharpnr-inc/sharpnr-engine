package accounting_models

import (
  "github.com/google/uuid"
  "time"
)

type AccountGroup struct {
  ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  BookID    uuid.UUID `gorm:"type:uuid;index"`
  Name      string    `gorm:"not null"`
  ParentID  *uuid.UUID
  Type      string
  CreatedAt time.Time
  UpdatedAt time.Time
}
