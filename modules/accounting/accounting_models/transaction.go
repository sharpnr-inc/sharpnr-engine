package accounting_models

import (
  "github.com/google/uuid"
  "time"
)

type Transaction struct {
  ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  OrgID        uuid.UUID
  BookID       uuid.UUID
  VoucherID    uuid.UUID
  AccountID    uuid.UUID
  Date         time.Time
  Type         string
  Amount       float64
  BalanceAfter float64
}
