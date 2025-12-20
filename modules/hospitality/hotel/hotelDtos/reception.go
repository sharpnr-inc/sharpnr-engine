package hoteldtos

// Response for the /hotel/reception
type ReceptionOrFrontdesk_Response struct {
	TotalRooms          int `json:"totalRooms"`
	UpcomingReservation int `json:"upcomingReservation"`
	RoomOccupied        int `json:"roomOccupied"`
	AvailableRooms      int `json:"availableRooms"`
	PendingRequest      int `json:"pendingRequest"`
}
