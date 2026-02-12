package hotel_constants

type HotelRoomStatusType string

const (
	Available        HotelRoomStatusType = "AVAILABLE"
	Occupied         HotelRoomStatusType = "OCCUPIED"
	Reserved         HotelRoomStatusType = "RESERVED"
	UnderMaintenance HotelRoomStatusType = "UNDER_MAINTENANCE"
	Blocked          HotelRoomStatusType = "BLOCKED"
)
