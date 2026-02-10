package restaurant_contracts

import restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"

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
