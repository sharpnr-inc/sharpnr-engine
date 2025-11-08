package accounting_models

import "github.com/google/uuid"

type TaxRate struct {
  ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  Name   string
  CGST   float64
  SGST   float64
  IGST   float64
  Active bool
}
