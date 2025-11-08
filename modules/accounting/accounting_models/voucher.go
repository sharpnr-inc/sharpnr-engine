package accounting_models

import (
  "github.com/google/uuid"
  "time"
)

type Voucher struct {
  ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  OrgID       uuid.UUID
  BookID      uuid.UUID
  VoucherType string
  VoucherNo   string
  VoucherDate time.Time
  Narration   string
  CreatedBy   uuid.UUID
  CreatedAt   time.Time
}
