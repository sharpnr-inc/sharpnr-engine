package hospitalitydtos

import (
	"time"

	"github.com/google/uuid"
)

type Restaurant_GetRestaurant_ResponseDTO struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	OrganizationID uuid.UUID `json:"organization_id"`

	IsPureVeg   bool      `json:"is_pure_veg"`
	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`

	// Floors + tables inside floors
	Floors []RestaurantFloorDTO `json:"floors"`

	// Direct table list (optional)
	Tables []RestaurantTableDTO `json:"tables"`

	IsKitchen bool        `json:"is_kitchen"`
	KitchenID *uuid.UUID  `json:"kitchen_id"`
	Kitchen   *KitchenDTO `json:"kitchen"`

	LegalName   string `json:"legal_name"`
	DisplayName string `json:"display_name"`
	GSTNumber   string `json:"gst_number"`
	FSSAINumber string `json:"fssai_number"`

	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// -------- FLOORS DTO --------

type RestaurantFloorDTO struct {
	ID             uuid.UUID `json:"id"`
	RestaurantID   uuid.UUID `json:"restaurant_id"`
	OrganizationID uuid.UUID `json:"organization_id"`

	Name        string `json:"name"`
	Description string `json:"description"`
	FloorNumber int    `json:"floor_number"`

	Tables []RestaurantTableDTO `json:"tables"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// -------- TABLE DTO --------

type RestaurantTableDTO struct {
	ID           uuid.UUID `json:"id"`
	RestaurantID uuid.UUID `json:"restaurant_id"`

	FloorID *uuid.UUID `json:"floor_id"`

	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Status   string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// -------- KITCHEN DTO --------
// If your Kitchen model has more fields, adjust here.

type KitchenDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
