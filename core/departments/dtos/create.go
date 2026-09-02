package orgDepartmentsDtos

import (
	"time"

	"github.com/google/uuid"
)

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

type DepartmentViewResponseDto struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	DepartmentHead *string   `json:"department_head,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
