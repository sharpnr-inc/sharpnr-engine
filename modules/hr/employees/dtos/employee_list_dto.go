package sharphr_dtos

import "github.com/google/uuid"

type OrgEmployeeList struct {
	ID           uuid.UUID `json:"id"`
	EmployeeName string    `json:"employeeName"`
	EmployeeRole string    `json:"employeeRole"`
	EmployeeID   string    `json:"employeeID"`
	Status       string    `json:"status"`
}

type EmployeeOrgProfile struct {
	ID           uuid.UUID `json:"id"`
	EmployeeName string    `json:"employeeName"`
	EmployeeRole string    `json:"employeeRole"`
	EmployeeID   string    `json:"employeeID"`
	Status       string    `json:"status"`

	ManagerName    string `json:"managerName"`
	DepartmentName string `json:"departmentName"`
}
