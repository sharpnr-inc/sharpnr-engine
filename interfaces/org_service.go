package interfaces

import (
	"github.com/google/uuid"
	dto "github.com/sharpnr-inc/sharpnr-engine/dtos"
)

// OrganizationService defines business logic for managing organizations.
type OrganizationService interface {
	// RegisterOrganization handles the creation of a new organization and returns DTO + token.
	RegisterOrganization(req dto.OrganizationRegistrationRequestDTO) (*dto.OrganizationRegistrationResponseDTO, string, error)

	// UpdateOrganization updates an existing organization based on UUID.
	UpdateOrganization(id uuid.UUID, req dto.OrganizationUpdateRequestDTO) (*dto.OrganizationUpdateResponseDTO, error)
}
