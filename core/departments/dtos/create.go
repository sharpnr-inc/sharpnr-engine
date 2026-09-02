package orgDepartmentsDtos

import (
	"time"

	"github.com/google/uuid"
)

// CreateDepartmentRequestDto is the payload for a new org-level department.
//
// Code is optional: when it is empty the service derives one from the name.
// DepartmentHead is the employee who runs it, as a UUID string; empty means
// nobody is in charge yet.
type CreateDepartmentRequestDto struct {
	Name           string `json:"name" validate:"required" binding:"required"`
	Description    string `json:"description" validate:"required" binding:"required"`
	Code           string `json:"code,omitempty"`
	DepartmentHead string `json:"department_head,omitempty"`
}

// UpdateDepartmentRequestDto is a partial update: only the fields present are
// written. ID is taken from the request path by the controller, so a body that
// carries a different id cannot retarget the update.
//
// DepartmentHead set to an empty string clears the department head.
type UpdateDepartmentRequestDto struct {
	ID             uuid.UUID `json:"id" validate:"required"`
	Name           *string   `json:"name,omitempty"`
	Description    *string   `json:"description,omitempty"`
	Code           *string   `json:"code,omitempty"`
	DepartmentHead *string   `json:"department_head,omitempty"`
}

type DepartmentViewResponseDto struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Code           string    `json:"code"`
	DepartmentHead *string   `json:"department_head,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
