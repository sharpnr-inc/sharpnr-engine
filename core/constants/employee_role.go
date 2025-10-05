package constants

// This is ENUMS for the Employee Roles in the Organisations

type EmployeeRoles string

const (
	Org_Owners  EmployeeRoles = "ORG_ADMIN"
	Super_Admin EmployeeRoles = "SUPER_ADMIN"
	Admin       EmployeeRoles = "ADMIN"
	Manager     EmployeeRoles = "MANAGER"
	Employee    EmployeeRoles = "EMPLOYEE"
)

type EmployeePermissions string

const (
	READ_ACCESS  EmployeePermissions = "READ_ACCESS"
	WRITE_ACCESS EmployeePermissions = "WRITE_ACCESS"
)
