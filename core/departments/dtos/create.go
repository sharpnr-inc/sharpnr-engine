package orgDepartmentsDtos

import "github.com/google/uuid"

type CreateDepartmentRequestDto struct {
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description" validate:"required"`
	DepartmentHead string `json:"department_head,omitempty"`
}

type UpdateDepartmentRequestDto struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
}
