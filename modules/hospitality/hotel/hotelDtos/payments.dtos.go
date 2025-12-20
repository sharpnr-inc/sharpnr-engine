package hoteldtos

import "github.com/google/uuid"

type AddPaymentsForBookingRequest_dto struct {
	BookingId   uuid.UUID
	RoomId      uuid.UUID
	AmountId    uuid.UUID
	Method      string
	ReferenceId string
}
