package orchestrator_dtos

type AddLegalEntity_REQ_DTO struct {
	Type         string
	Name         string
	Address_line string
	Country      string
	State        string
	City         string
	Pincode      string
	GST          *string
	CIN          *string
	PAN          *string
}
