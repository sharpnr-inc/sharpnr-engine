package hospitalitydtos

import "github.com/google/uuid"

type Restaurant_GetRestaurant_ResponseDTO struct {
	RestaurantID   uuid.UUID
	ResTaurantName string
}
