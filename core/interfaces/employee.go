package interfaces

// This is Interface for Service
type EmployeeService interface {
	AddEmployeeInOrg()
	RemoveEmployeeInOrg()
	UpdateEmployeeDetails()
}

// This is Interface for Employee Repo
type EmployeeRepository interface {
	CreateEmployee()
	FindEmployeeById()
	RemoveEmployee()
	UpdateEmployeeDetails()
}
