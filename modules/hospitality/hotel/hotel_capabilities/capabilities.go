package hotel_capabilities

type HotelCapabilities string

const (
	HotelCreateBooking        HotelCapabilities = "HOTEL:CREATE_BOOKING"
	HotelAddPaymentsInBooking HotelCapabilities = "HOTEl:ADD_PAYMENT_IN_BOOKING"
)
