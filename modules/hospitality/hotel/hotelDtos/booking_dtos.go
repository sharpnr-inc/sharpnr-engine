package hoteldtos

import (
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	"time"

	"github.com/google/uuid"
)

type BookingGuest_dto struct {
	Name   string  `json:"guestName"`
	Email  *string `json:"guestEmail"`
	Mobile string  `json:"guestMobile"`
}

type RoomBookingRequest_dto struct {
	RoomId         uuid.UUID `json:"roomId"`
	PrimaryGuestID uuid.UUID `json:"primaryGuestId"`

	CheckinDate  time.Time `json:"checkinDate"`
	CheckoutDate time.Time `json:"checkoutDate"`

	BookingGuests []BookingGuest_dto `json:"bookingGuests"`
	BookingDate   string             `json:"bookingDate"`
}

type BookingDetailsDTO struct {
	Booking BookingInfoDTO `json:"booking"`
	Room    RoomInfoDTO    `json:"room"`
	Guest   GuestInfoDTO   `json:"guest"`

	Charges  []ChargeDTO  `json:"charges"`
	Payments []PaymentDTO `json:"payments"`

	Bill    *BillDTO `json:"bill,omitempty"`
	Balance float64  `json:"balance"`
}

type BookingInfoDTO struct {
	ID uuid.UUID `json:"id"`

	CheckInDate  time.Time `json:"checkInDate"`
	CheckOutDate time.Time `json:"checkOutDate"`
	Status       string    `json:"status"`

	CreatedAt time.Time `json:"createdAt"`
}

type RoomInfoDTO struct {
	ID       uuid.UUID `json:"id"`
	RoomName string    `json:"roomName"`
	RoomType string    `json:"roomType"`
}

type GuestInfoDTO struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Mobile string    `json:"mobile"`
}

type ChargeDTO struct {
	ID          uuid.UUID `json:"id"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"createdAt"`
}

type PaymentDTO struct {
	ID        uuid.UUID                          `json:"id"`
	Amount    float64                            `json:"amount"`
	Method    hospitalityconstants.PaymentMethod `json:"method"`
	Status    hospitalityconstants.PaymentStatus `json:"status"`
	CreatedAt time.Time                          `json:"createdAt"`
}

type BillDTO struct {
	ID          uuid.UUID `json:"id"`
	BaseAmount  float64   `json:"baseAmount"`
	TaxAmount   float64   `json:"taxAmount"`
	TotalAmount float64   `json:"totalAmount"`
	Status      string    `json:"status"`
}
