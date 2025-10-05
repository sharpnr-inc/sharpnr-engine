package policyengine

// __METADATA__ section
type Metadata struct {
	OrgID     string  `json:"ORG_ID"`
	SharpnrID string  `json:"SHARPNR_ID"`
	PolicyID  string  `json:"POLICY_ID"`
	Version   float64 `json:"VERSION"`
	OrgType   string  `json:"ORG_TYPE"`
	CreatedAt string  `json:"CREATED_AT"`
	UpdatedAt string  `json:"UPDATED_AT"`
	// Modules   []interface{} `json:"MODULES"` // mixed: string + object
}

type Policy struct {
	SuperAdmins []string `json:"SUPER_ADMINS"`
	OrgOwners   []string `json:"ORG_OWNERS"`
}

type OrgDefaultPolicy struct {
	Metadata Metadata `json:"__METADATA__"`
	Policy   Policy   `json:"ORG_POLICY"`
}

// Role types
type Role string

const (
	RoleSuperAdmin Role = "SUPER_ADMIN"
	RoleOrgOwner   Role = "ORG_OWNER"
	RoleAdmin      Role = "ADMIN"
	RoleManager    Role = "MANAGER"
	RoleEmployee   Role = "EMPLOYEE"
)

// Permission types
type Permission string

const (
	PermissionRead   Permission = "READ"
	PermissionWrite  Permission = "WRITE"
	PermissionDelete Permission = "DELETE"
	PermissionManage Permission = "MANAGE"
)

func CreatePolicyFromTemplate() {

}
