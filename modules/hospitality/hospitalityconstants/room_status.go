package hospitalityconstants

type RoomStatus string

const (
	Available         RoomStatus = "AVAILABLE"
	Booked            RoomStatus = "BOOKED"
	Under_maintenance RoomStatus = "UNDER_MAINTENANCE"
	Cleaning          RoomStatus = "ON_Cleaning"
	Enabled           RoomStatus = "ENABLED"
	Disabled          RoomStatus = "DISABLED"
)
