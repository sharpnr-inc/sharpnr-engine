package mappers

import (
	dto "github.com/sharpnr-inc/sharpnr-engine/dtos"
	"github.com/sharpnr-inc/sharpnr-engine/models"
)

func ToOrganizationModel(req dto.OrganizationRegistrationRequestDTO) *models.Organization {
	return &models.Organization{
		Name:            req.Name,
		OrgOprationType: req.OrgOprationType,
		RootEmail:       req.RootEmail,
	}

}

func ToOrganizationRegistrationResponseDTO(org *models.Organization) dto.OrganizationRegistrationResponseDTO {
	return dto.OrganizationRegistrationResponseDTO{
		SharpnrOrgID:    org.SharpnrOrgID,
		Name:            org.Name,
		OrgOprationType: org.OrgOprationType,
		RootEmail:       org.RootEmail,
		CreatedAt:       org.CreatedAt,
	}
}
