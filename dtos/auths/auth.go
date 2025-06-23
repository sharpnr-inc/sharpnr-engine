package dto

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type OrganisationRegisterDTO struct {
	Name                string `json:"name" binding:"required"`
	OrganizationID      string `json:"organizationID" binding:"required,len=12"`
	Address             string `json:"address,omitempty"`
	Phone               int    `json:"phone" binding:"required"`
	Email               string `json:"email" binding:"required,email"`
	IsPhysicalVerified  *bool  `json:"isPhysicalVerified,omitempty"` // optional, pointer so omitted if nil
	ISRegistered        *bool  `json:"isRegistered,omitempty"`
	CinNumber           string `json:"cinNumber,omitempty"`
	GstNumber           string `json:"gstNumber,omitempty"`
	OrganizationRegType string `json:"organizationRegType,omitempty"`
	Banker              string `json:"banker,omitempty"`
	BankAccNumber       string `json:"bankAccNumber" binding:"required"`
	IfscCode            string `json:"ifscCode,omitempty"`
	BankBranchName      string `json:"bankBranchName,omitempty"`
	Password            string `json:"password" binding:"required,min=6"`
	AdminName           string `json:"adminName" binding:"required"`
}

type OrganizationResponseDTO struct {
	ID                  uint   `json:"id"`
	Name                string `json:"name"`
	OrganizationID      string `json:"organizationID"`
	Address             string `json:"address,omitempty"`
	Phone               int    `json:"phone"`
	Email               string `json:"email"`
	IsPhysicalVerified  bool   `json:"isPhysicalVerified"`
	ISRegistered        bool   `json:"isRegistered"`
	CinNumber           string `json:"cinNumber,omitempty"`
	GstNumber           string `json:"gstNumber,omitempty"`
	OrganizationRegType string `json:"organizationRegType,omitempty"`
	Banker              string `json:"banker,omitempty"`
	BankAccNumber       string `json:"bankAccNumber"`
	IfscCode            string `json:"ifscCode,omitempty"`
	BankBranchName      string `json:"bankBranchName,omitempty"`
	AdminName           string `json:"adminName"`
	CreatedAt           string `json:"createdAt"`
}
