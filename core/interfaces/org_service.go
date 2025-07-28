package interfaces

import dto "sharpnr-engine/core/dtos"

// This Interfaces is responsible for the Creating and Validation for the Organisations
// it takes a "req" of Type "dto.OrganizationRegistrationRequestDTO" and Returns "*dto.OrganizationRegistrationResponseDTO"
// string, error
type CoreOrganisationService interface {
	// This Will Register Organisation
	RegisterOrganization(req dto.OrganizationRegistrationRequestDTO) (*dto.OrganizationRegistrationResponseDTO, string, error)
}
