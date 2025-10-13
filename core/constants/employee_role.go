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
