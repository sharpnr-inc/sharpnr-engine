package oragnisation

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

type CreateEmployeeRequestDTO struct {
}
