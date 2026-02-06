package restaurant_contracts

import "github.com/google/uuid"

type CreateFloorInput struct {
	RestaurantID uuid.UUID
	FloorName    string
	FloorNumber  int
}

type CreateTablesInFloor struct {
	RestaurantID uuid.UUID
	FloorID      uuid.UUID
	Capacity     *int
	Name         *string
}
