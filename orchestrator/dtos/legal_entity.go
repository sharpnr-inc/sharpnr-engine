package orchestrator_dtos

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
