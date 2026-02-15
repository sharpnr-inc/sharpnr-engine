package restaurant_contracts

import (
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"

	"github.com/google/uuid"
)

type RestaurantFloorLayout_Response_dto struct {
	Floor_Name   string
	Floor_Number int
	Tables       []RestaurantTableLayout_Response_dto
}

type RestaurantLayout_Response_dto struct {
	Layout []RestaurantFloorLayout_Response_dto
}

type RestaurantTableLayout_Response_dto struct {
	TableName string
	TableCode string
	FloorID   string
	TableID   string
	Status    restaurant_constants.RestaurantTableStatusType
	Capacity  int
}

type OrdersResquest_dto struct {
	PropertyID             uuid.UUID
	OwnersID               uuid.UUID
	TableID                *uuid.UUID
	RoomID                 *uuid.UUID
	OrderProcessEmployeeID *uuid.UUID // this is the ID for the employeeID who took the order
	OrderType              string     // DINE_IN | TAKEAWAY | ROOM_SERVICE
	Orders                 []OrdersItem
}

type OrdersItem struct {
	ItemID             uuid.UUID
	Qty                int
	SpecialInstruction string
	VariantID          *uuid.UUID
	AddonIDs           []uuid.UUID
}
