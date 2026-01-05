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

type Privileges struct {
}

type EmployeeAccess struct {
	EmployeeID string             `json:"EmployeeID"`
	Privileges EmployeePrivileges `json:"Privileges"`
}

type EmployeePrivileges struct {
	SharpAccounting *ModulePrivileges `json:"SharpAccounting,omitempty"`
	SharpHR         *ModulePrivileges `json:"SharpHR,omitempty"`
	SharpHotel      *ModulePrivileges `json:"SharpHotel,omitempty"`
}

type ModulePrivileges struct {
	Access        string   `json:"Access"`
	Transactions  []string `json:"transactions,omitempty"`
	Invoice       []string `json:"invoice,omitempty"`
	BalanceSheets []string `json:"balanceSheets,omitempty"`
}

type Policy struct {
	SuperAdmins     []string `json:"SUPER_ADMINS"`
	OrgOwners       []string `json:"ORG_OWNERS"`
	EMPLOYEE_ACCESS []EmployeeAccess
}

type OrgDefaultPolicy struct {
	Policy   Policy   `json:"Policy"`
	Metadata Metadata `json:"Metadata"`
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
)

// EmployeePolicy is the structure you want to return
type EmployeePolicy struct {
	EmployeeID string                 `json:"EmployeeID"`
	Privileges map[string]interface{} `json:"Privileges"`
}

type OrgPolicy struct {
	OrgOwners      []string         `json:"ORG_OWNERS"`
	SuperAdmins    []string         `json:"SUPER_ADMINS"`
	EmployeeAccess []EmployeeAccess `json:"EMPLOYEE_ACCESS"`
}
