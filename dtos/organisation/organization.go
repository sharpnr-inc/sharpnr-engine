package oragnisation

type CreateOrganizationDTO struct {
	Name                string  `json:"name" binding:"required"`
	OrganizationID      string  `json:"organization_id" binding:"required,len=12"`
	Address             *string `json:"address,omitempty"`
	Phone               int     `json:"phone" binding:"required"`       // required
	Email               string  `json:"email" binding:"required,email"` // required and valid email
	IsPhysicalVerified  *bool   `json:"is_physical_verified,omitempty"` // optional, defaults false
	ISRegistered        *bool   `json:"is_registered,omitempty"`        // optional, defaults false
	CinNumber           *string `json:"cin_number,omitempty"`
	GstNumber           *string `json:"gst_number,omitempty"`
	OrganizationRegType *string `json:"organization_reg_type,omitempty"`
	Banker              *string `json:"banker,omitempty"`
	BankAccNumber       string  `json:"bank_acc_number" binding:"required"`
	IfscCode            *string `json:"ifsc_code,omitempty"`
	BankBranchName      *string `json:"bank_branch_name,omitempty"`
	Password            string  `json:"password" binding:"required"`
	AdminName           string  `json:"admin_name" binding:"required"`
}

type UpdateOrganizationDTO struct {
	Name                *string `json:"name,omitempty"`
	Address             *string `json:"address,omitempty"`
	Phone               *int    `json:"phone,omitempty"`
	Email               *string `json:"email,omitempty"`
	IsPhysicalVerified  *bool   `json:"is_physical_verified,omitempty"`
	ISRegistered        *bool   `json:"is_registered,omitempty"`
	CinNumber           *string `json:"cin_number,omitempty"`
	GstNumber           *string `json:"gst_number,omitempty"`
	OrganizationRegType *string `json:"organization_reg_type,omitempty"`
	Banker              *string `json:"banker,omitempty"`
	BankAccNumber       *string `json:"bank_acc_number,omitempty"`
	IfscCode            *string `json:"ifsc_code,omitempty"`
	BankBranchName      *string `json:"bank_branch_name,omitempty"`
	Password            *string `json:"password,omitempty"`
	AdminName           *string `json:"admin_name,omitempty"`
}
