package interfaces

import dto "github.com/sharpnr-inc/sharpnr-engine/core/dtos"

type CoreOrganisationService interface {
	RegisterOrganization(req dto.OrganizationRegistrationRequestDTO) (*dto.OrganizationRegistrationResponseDTO, string, error)
}
