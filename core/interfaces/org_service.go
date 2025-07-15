package interfaces

import dto "github.com/sharpnr-inc/sharpnr-engine/core/dtos"

type CoreOrganisationService interface {
	// This Will Register Organisation
	RegisterOrganization(req dto.OrganizationRegistrationRequestDTO) (*dto.OrganizationRegistrationResponseDTO, string, error)
}
