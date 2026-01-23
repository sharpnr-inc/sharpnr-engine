package orchestrator_dtos

type SetupOrganizationHQAddressReqDTO struct {
	AddressLine string   `json:"address_line" binding:"required"`
	City        string   `json:"city" binding:"required"`
	State       string   `json:"state" binding:"required"`
	Country     string   `json:"country" binding:"required"`
	Pincode     string   `json:"pincode" binding:"required"`
	Latitude    *float64 `json:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
	AccuracyM   *int     `json:"accuracy_m,omitempty"`
}
