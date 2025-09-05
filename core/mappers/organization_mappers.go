package mappers

import (
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"
)

// This is the Mapper, it takes the dto of Request for Registering the organisation and returns the pointer of Organisation model
func ToOrganizationModel(req dto.OrganizationRegistrationRequestDTO) *models.Organization {
	return &models.Organization{
		Name: req.Name,
		// OrgOprationType:     req.OrgOprationType,
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
		MasterPhoneNumber:   req.MasterContact,
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
