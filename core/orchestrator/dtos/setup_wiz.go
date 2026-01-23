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

type SetupOrgPropertyReqDTO struct {
	Name        string   `json:"name"`
	AddressLine string   `json:"address_line"`
	City        string   `json:"city"`
	State       string   `json:"state"`
	Country     string   `json:"country"`
	Pincode     string   `json:"pincode"`
	Longitude   *float64 `json:"longitude,omitempty"`
	Latitude    *float64 `json:"latitude,omitempty"`
	AccuracyM   *int     `json:"accuracy_m,omitempty"`
}

type SetupLegalEntityReqDTO struct {
	Name        string   `json:"name"`
	AddressLine string   `json:"address_line"`
	City        string   `json:"city"`
	State       string   `json:"state"`
	Country     string   `json:"country"`
	Pincode     string   `json:"pincode"`
	Longitude   *float64 `json:"longitude,omitempty"`
	Latitude    *float64 `json:"latitude,omitempty"`
	AccuracyM   *int     `json:"accuracy_m,omitempty"`
}
