package hospitalitydtos

import (
	"time"

	"github.com/google/uuid"
)

type Restaurant_GetRestaurant_ResponseDTO struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	OrganizationID uuid.UUID `json:"organization_id"`

	IsPureVeg   bool   `json:"is_pure_veg"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`

	IsKitchen bool        `json:"is_kitchen"`
	KitchenID *uuid.UUID  `json:"kitchen_id"`
	Kitchen   *KitchenDTO `json:"kitchen"`

	LegalName   string `json:"legal_name"`
	DisplayName string `json:"display_name"`
	GSTNumber   string `json:"gst_number"`
	CINNumber   string `json:"cin_number"`
	FSSAINumber string `json:"fssai_number"`

	ContactEmail  string `json:"contact_email"`
	ContactNumber string `json:"contact_number"`
	ManagerName   string `json:"manager_name"`

	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
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

// This is the struct fir the Updating the Restaurant
type Restaurant_UpdateOrEdit_General_Details_REQ struct {
	RestaurantId          string  `json:"restaurant_id" binding:"required"`
	RestaurantLegalName   *string `json:"restaurant_legal_name,omitempty"`
	RestaurantDisplayName *string `json:"restaurant_display_name,omitempty"`
	FSSAINumber           *string `json:"fssai_number,omitempty"`
	GSTNumber             *string `json:"gst_number,omitempty"`
	CINNumber             *string `json:"cin_number,omitempty"`
	AddressLine           *string `json:"address_line,omitempty"`
	City                  *string `json:"city,omitempty"`
	State                 *string `json:"state,omitempty"`
	Country               *string `json:"country,omitempty"`
	Pincode               *string `json:"pincode,omitempty"`
	EmailAddress          *string `json:"email_address,omitempty"`
	ContactNumber         *string `json:"contact_number,omitempty"`
	ManagerName           *string `json:"manager_name,omitempty"`
}

type Restaurant_GetRestaurantLayout_Response struct {
	RestaurantId uuid.UUID            `json:"restaurant_id"`
	Floor        []RestaurantFloorDTO `json:"floors"`
}

type Restaurant_CreateOrAddTables_Request struct {
	RestaurantId uuid.UUID                      `json:"restaurant_id"`
	FloorId      uuid.UUID                      `json:"floor_id"`
	Tables       []Restaurant_TableCreatingList `json:"tables"`
}

type Restaurant_TableCreatingList struct {
	TableName     string `json:"table_name"`
	TableCapacity int    `json:"table_capacity"`
}
