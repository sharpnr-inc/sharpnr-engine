package dto

type CreateEmployeeInOrganisationRequestDTO struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=6"`
	OrganizationID uint   `json:"organizationId" binding:"required"`
}

type CreateEmployeeInOrganisationResponseDTO struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	OrganizationID uint   `json:"organizationId"`
	CreatedAt      string `json:"createdAt"`
}

type UpdateEmployeeInOrganisationRequestDTO struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	Password *string `json:"password,omitempty" binding:"omitempty,min=6"`
}

type UpdateEmployeeInOrganisationResponseDTO struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	OrganizationID uint   `json:"organizationId"`
	UpdatedAt      string `json:"updatedAt"`
}
