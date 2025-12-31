package sharphrdtos

import (
	"time"

	"github.com/google/uuid"
)

type CheckIfEmployeeExist_REQ_DTO struct {
	EmployeeEmail string `json:"employee_email"`
}

type CheckIfEmployeeExist_RESPONSE_DTO struct {
	Exists string `json:"employee_exist"`
	Test   string `json:"query_email"`
	ORG    any    `json:"orguuid"`
	// Message string `json:"message"`
}

type EmployeeData struct {
	ID    string
	Name  string
	Email string
	Role  string
}

type ListAllEmployee_RESPONSE_DTO struct {
	EmployeeList []EmployeeData `json:"employee_list"`
}

type GetEmployeeDetails_RESPONSE_DTO struct {
	ID           string    `json:"id"`
	OrgID        any       `json:"org_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Designation  string    `json:"designation"`
	Department   string    `json:"department"`
	JoiningDate  string    `json:"joining_date"`
	AccessLevel  string    `json:"access_level"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	LastModified time.Time `json:"last_modified"`
}

type CreateEmployeeInOrganization_RequestDTO struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Mobile         string `json:"mobile_number"`
	OrganizationID string `json:"organizationId" binding:"required"`
	Role           string `json:"employee_role"`
}

type CreateEmployeeInOrganizationResponseDTO struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	OrganizationID string    `json:"organizationId"`
	CreatedAt      string    `json:"createdAt"`
}
