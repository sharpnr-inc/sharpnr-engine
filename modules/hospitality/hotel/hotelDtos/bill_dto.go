package hoteldtos

import "github.com/google/uuid"

type BillPreviewDTO struct {
	BookingID uuid.UUID

	BaseAmount  float64
	TaxAmount   float64
	TotalAmount float64

	TaxBreakup []TaxComponentPreviewDTO

	PaidAmount float64
	Balance    float64
}

type TaxComponentPreviewDTO struct {
	ComponentType string // CGST / SGST / IGST
	Rate          float64
	Amount        float64
}
