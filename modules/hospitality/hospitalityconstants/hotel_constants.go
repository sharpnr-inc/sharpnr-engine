package hospitalityconstants

type HotelRoomStatusType string

const (
	Available         HotelRoomStatusType = "AVAILABLE"
	Booked            HotelRoomStatusType = "BOOKED"
	Under_maintenance HotelRoomStatusType = "UNDER_MAINTENANCE"
	Cleaning          HotelRoomStatusType = "ON_Cleaning"
	Enabled           HotelRoomStatusType = "ENABLED"
	Disabled          HotelRoomStatusType = "DISABLED"
)

type HotelChargeType string

const (
	// -------------------------
	// Room related
	// -------------------------
	RoomRent           HotelChargeType = "ROOM_RENT"
	EarlyCheckInCharge HotelChargeType = "EARLY_CHECKIN"
	LateCheckOutCharge HotelChargeType = "LATE_CHECKOUT"

	// -------------------------
	// Food & Beverage
	// -------------------------
	FoodOrder HotelChargeType = "FOOD_ORDER"
	BarOrder  HotelChargeType = "BAR_ORDER"
	Minibar   HotelChargeType = "MINIBAR"

	// -------------------------
	// Services
	// -------------------------
	LaundryService HotelChargeType = "LAUNDRY"
	SpaService     HotelChargeType = "SPA"
	HouseKeeping   HotelChargeType = "HOUSE_KEEPING"
	RoomService    HotelChargeType = "ROOM_SERVICE"

	// -------------------------
	// Misc
	// -------------------------
	DamageCharge  HotelChargeType = "DAMAGE"
	ServiceCharge HotelChargeType = "SERVICE_CHARGE"
)

type HotelBookingStatusType string

const (
	HotelRoomCheckedId  HotelBookingStatusType = "CHECKED_IN"
	HotelRoomCheckedOut HotelBookingStatusType = "CHECKED_OUT"
)
