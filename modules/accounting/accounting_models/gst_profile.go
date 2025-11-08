package accounting_models

import "github.com/google/uuid"

type GSTProfile struct {
  ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  OrgID   uuid.UUID
  GSTIN   string
  State   string
  Type    string
  Status  string
}
