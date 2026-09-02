package orgDepartmentsDtos

type CreateDepartmentRequest struct {
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description" validate:"required"`
	OrganizationID string `json:"organization_id" validate:"required"`
}
