package mappers

import (
	dto "github.com/sharpnr-inc/sharpnr-engine/dtos"
	"github.com/sharpnr-inc/sharpnr-engine/models"
)

func ToOrganizationResponseDTO(org *models.Organization) dto.OrganizationRegistrationResponseDTO {
	return dto.OrganizationRegistrationResponseDTO{
		SharpnrOrgID:    org.SharpnrOrgID,
		Name:            org.Name,
		OrgOprationType: org.OrgOprationType,
		RootEmail:       org.RootEmail,
		CreatedAt:       org.CreatedAt,
	}
}
