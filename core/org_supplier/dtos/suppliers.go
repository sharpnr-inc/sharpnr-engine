package suppliers_dtos

type CreateSuppliersAPIRequestdto struct {
	SuppliersName string `json:"suppliers_name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`

	AddressLine string `json:"address_line"`
	City        string `json:"city"`
	State       string `json:"state"`
	Pincode     string `json:"pincode"`
	Country     string `json:"country"`
}
