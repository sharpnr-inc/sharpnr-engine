package mappers

import (
	dto "github.com/sharpnr-inc/sharpnr-engine/dtos"
	"github.com/sharpnr-inc/sharpnr-engine/models"
)

func ToOrganizationModel(req dto.OrganizationRegistrationRequestDTO) *models.Organization {
	return &models.Organization{
		Name:                req.Name,
		OrgOprationType:     req.OrgOprationType,
		RootEmail:           req.RootEmail,
		OrgDescription:      req.Description,
		Address:             req.Address,
		State:               req.State,
		City:                req.City,
		Pincode:             req.Pincode,
		Country:             req.Country,
		PublicContact:       req.PublicContact,
		ContactNumber:       req.ContactNumber,
		ContactNumber2:      req.ContactNumber2,
		MasterContact:       req.MasterContact,
		SecondaryEmail:      req.SecondaryEmail,
		CinNumber:           req.CinNumber,
		GstNumber:           req.GstNumber,
		OrganizationRegType: req.OrganizationRegType,
		Banker:              req.Banker,
		BankAccNumber:       req.BankAccNumber,
		IfscCode:            req.IfscCode,
		BankBranchName:      req.BankBranchName,
		UPIID:               req.UPIID,
		Password:            req.Password,
		MasterPassword:      req.MasterPassword,
	}

}

func ToOrganizationRegistrationResponseDTO(org *models.Organization) dto.OrganizationRegistrationResponseDTO {
	return dto.OrganizationRegistrationResponseDTO{
		SharpnrOrgID:    org.SharpnrOrgID,
		Name:            org.Name,
		Description:     org.OrgDescription,
		OrgOprationType: org.OrgOprationType,
		RootEmail:       org.RootEmail,
		CreatedAt:       org.CreatedAt,
	}
}
