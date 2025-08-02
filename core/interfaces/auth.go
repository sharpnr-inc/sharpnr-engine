package interfaces

import (
	dto "sharpnr-engine/core/dtos"
)

// AuthService defines the interface for authentication services.
type AuthService interface {
	RegisterOrganization(req dto.OrganizationRegistrationRequestDTO) (*dto.OrganizationRegistrationResponseDTO, string, error)
	Login(req dto.OrgAuthRequestDTO) (*dto.OrgAuthResponseDTO, error)
}
