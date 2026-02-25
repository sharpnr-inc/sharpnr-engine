package orchestrator_dtos

import "github.com/google/uuid"

type AddLegalEntity_REQ_DTO struct {
	Type         string  `json:"type"`
	Name         string  `json:"name"`
	Address_line string  `json:"address_line"`
	Country      string  `json:"country"`
	State        string  `json:"state"`
	City         string  `json:"city"`
	Pincode      string  `json:"pincode"`
	GST          *string `json:"gst_number,omitempty"`
	CIN          *string `json:"cin_number,omitempty"`
	PAN          *string `json:"pan_number,omitempty"`
}

type LegalEntity_Response_dto struct {
	Name        string    `json:"name"`
	ID          uuid.UUID `json:"id"`
	Type        string    `json:"type"`
	AddressLine string    `json:"addressLine"`
	Country     string    `json:"country"`
	State       string    `json:"state"`
	City        string    `json:"city"`
	Pincode     string    `json:"pincode"`
	GST         *string   `json:"gst"`
	CIN_Number  *string   `json:"cinNumber"`
	PAN_Number  *string   `json:"panNumber"`
}
