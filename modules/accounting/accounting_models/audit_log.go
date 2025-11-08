package accounting_models

import (
  "github.com/google/uuid"
  "time"
)

type AuditLog struct {
  ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  OrgID     uuid.UUID
  UserID    uuid.UUID
  Action    string
  Entity    string
  EntityID  uuid.UUID
  Timestamp time.Time
}
