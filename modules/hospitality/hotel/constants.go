package hotel

type HOTEL_ROOM_STATUS string

const (
	Available        HOTEL_ROOM_STATUS = "AVAILABLE"
	Occupied         HOTEL_ROOM_STATUS = "OCCUPIED"
	Unavailable      HOTEL_ROOM_STATUS = "UNAVAILABLE"
	UnderMaintenance HOTEL_ROOM_STATUS = "UNDER_MAINTENANCE"
)

type HOTEL_BILL_STATUS string

const (
	Pending HOTEL_BILL_STATUS = "PENDING"
	Paid    HOTEL_BILL_STATUS = "PAID"
)
