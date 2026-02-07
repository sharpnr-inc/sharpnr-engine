package restaurant_contracts

import "github.com/google/uuid"

type CreateFloorInput struct {
	RestaurantID uuid.UUID
	FloorName    string
	FloorNumber  int
}

type CreateTablesInFloorInput struct {
	RestaurantID uuid.UUID
	FloorID      uuid.UUID
	Capacity     *int
	Name         *string
	Code         *string
}

type SetupRestaurantLayoutInput struct {
	Floor_Name       string
	Floor_Number     int
	Number_of_tables int
}

type UpdateGeneralDetailsInput struct {
	Name         string
	IsConfigured bool
	IsPureVeg    bool
	OpeningTime  string
	ClosingTime  string
}
