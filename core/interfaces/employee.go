package interfaces

type EmployeeService interface {
	AddEmployeeInOrg()
	RemoveEmployeeInOrg()
	UpdateEmployeeDetails()
}

type EmployeeRepository interface {
	CreateEmployee()
	FindEmployeeById()
	RemoveEmployee()
	UpdateEmployeeDetails()
}
