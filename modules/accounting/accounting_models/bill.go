package accounting_models

import (
  "github.com/google/uuid"
  "time"
)

type Bill struct {
  ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  OrgID        uuid.UUID
  BookID       uuid.UUID
  SupplierID   uuid.UUID
  BillNo       string
  BillDate     time.Time
  DueDate      time.Time
  SubTotal     float64
  TaxAmount    float64
  TotalAmount  float64
  Status       string
  CreatedAt    time.Time
}
