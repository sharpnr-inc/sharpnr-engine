package employee

// CreateEmployeeDTO represents the payload to create a new employee.
//
// All fields are required and validated using Gin binding tags.
type CreateEmployeeDTO struct {
	FullName string `json:"full_name" binding:"required"`   // Full name of the employee
	Email    string `json:"email" binding:"required,email"` // Email address (must be valid)
	Phone    string `json:"phone" binding:"required"`       // Phone number
	Role     string `json:"role" binding:"required"`        // Role assigned to the employee (e.g., admin, manager)
}

// UpdateEmployeeDTO represents the data for updating an employee.
//
// Fields are optional and only provided fields will be updated.
type UpdateEmployeeDTO struct {
	FullName *string `json:"full_name,omitempty"` // New full name
	Phone    *string `json:"phone,omitempty"`     // New phone number
	Role     *string `json:"role,omitempty"`      // New role
}

// EmployeeResponseDTO is the shape of the employee data returned in API responses.
type EmployeeResponseDTO struct {
	ID       uint   `json:"id"`        // Unique employee ID
	FullName string `json:"full_name"` // Full name of the employee
	Email    string `json:"email"`     // Email address
	Phone    string `json:"phone"`     // Phone number
	Role     string `json:"role"`      // Role assigned to the employee
}
