package hoteldtos

import (
	"sharpnr-engine/modules/hospitality/hospitalityconstants"

	"github.com/google/uuid"
)

type AddPaymentsForBookingRequest_dto struct {
	BookingID uuid.UUID                          `json:"booking_id" binding:"required"`
	Amount    float64                            `json:"amount" binding:"required,gt=0"`
	Method    hospitalityconstants.PaymentMethod `json:"method" binding:"required"`
	Reference *string                            `json:"reference,omitempty"`
}
