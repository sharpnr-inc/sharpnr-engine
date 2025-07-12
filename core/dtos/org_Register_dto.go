package dto

type OrganizationRegistrationRequestDTO struct {
	Name                string  `json:"name" validate:"required,min=3,max=100"`
	Description         string  `json:"description" validate:"required,min=10,max=500"`
	OrgOprationType     string  `json:"org_operation_type" validate:"required"`
	Address             string  `json:"address" validate:"required"`
	State               string  `json:"state" validate:"required"`
	City                string  `json:"city" validate:"required"`
	Pincode             string  `json:"pincode" validate:"required"`
	Country             string  `json:"country" validate:"required"`
	PublicContact       string  `json:"public_contact" validate:"required"`
	ContactNumber       string  `json:"contact_number" validate:"required"`
	ContactNumber2      *string `json:"contact_number2,omitempty"`
	MasterContact       *string `json:"master_contact,omitempty"`
	RootEmail           string  `json:"root_email" validate:"required,email"`
	SecondaryEmail      *string `json:"secondary_email,omitempty"`
	CinNumber           *string `json:"cin_number,omitempty"`
	GstNumber           *string `json:"gst_number,omitempty"`
	OrganizationRegType *string `json:"organization_reg_type,omitempty"`
	Banker              *string `json:"banker,omitempty"`
	BankAccNumber       string  `json:"bank_acc_number" validate:"required"`
	IfscCode            *string `json:"ifsc_code,omitempty"`
	BankBranchName      *string `json:"bank_branch_name,omitempty"`
	UPIID               *string `json:"upi_id,omitempty"`
	Password            string  `json:"password" validate:"required,min=6"`
	MasterPassword      string  `json:"master_password" validate:"required,min=6"`
}

type OrganizationRegistrationResponseDTO struct {
	SharpnrOrgID    string `json:"sharpnr_org_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	OrgOprationType string `json:"org_operation_type"`
	RootEmail       string `json:"root_email"`
	CreatedAt       int64  `json:"created_at"`
}

type OrganizationUpdateRequestDTO struct {
	Name                *string `json:"name,omitempty"`
	Address             *string `json:"address,omitempty"`
	State               *string `json:"state,omitempty"`
	City                *string `json:"city,omitempty"`
	Pincode             *string `json:"pincode,omitempty"`
	Country             *string `json:"country,omitempty"`
	ContactNumber       *string `json:"contact_number,omitempty"`
	ContactNumber2      *string `json:"contact_number2,omitempty"`
	MasterContact       *string `json:"master_contact,omitempty"`
	SecondaryEmail      *string `json:"secondary_email,omitempty"`
	GstNumber           *string `json:"gst_number,omitempty"`
	OrganizationRegType *string `json:"organization_reg_type,omitempty"`
	Banker              *string `json:"banker,omitempty"`
	BankAccNumber       *string `json:"bank_acc_number,omitempty"`
	IfscCode            *string `json:"ifsc_code,omitempty"`
	BankBranchName      *string `json:"bank_branch_name,omitempty"`
	UPIID               *string `json:"upi_id,omitempty"`
	MasterPassword      string
}

type OrganizationUpdateResponseDTO struct {
	Message string `json:"message"`
	Updated bool   `json:"updated"`
}
