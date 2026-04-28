package employee_contracts

import "github.com/google/uuid"

type EmployeeProfileOutput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	DOB       string `json:"dob"`
	Gender    string `json:"gender"`
}

type EmployeeDetailsOutput struct {
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	ID        uuid.UUID `json:"id"`
	Role      string    `json:"role"`
}
