package mappers

import (
	"time"

	dto "github.com/sharpnr-inc/sharpnr-engine/dtos/auths"
	"github.com/sharpnr-inc/sharpnr-engine/models"
)

func ToOrganizationResponseDTO(org *models.Organization) dto.OrganizationResponseDTO {
	return dto.OrganizationResponseDTO{
		ID:                  org.ID,
		Name:                org.Name,
		OrganizationID:      org.OrganizationID,
		Address:             org.Address,
		Phone:               org.Phone,
		Email:               org.Email,
		IsPhysicalVerified:  org.IsPhysicalVerified,
		ISRegistered:        org.ISRegistered,
		CinNumber:           org.CinNumber,
		GstNumber:           org.GstNumber,
		OrganizationRegType: org.OrganizationRegType,
		Banker:              org.Banker,
		BankAccNumber:       org.BankAccNumber,
		IfscCode:            org.IfscCode,
		BankBranchName:      org.BankBranchName,
		AdminName:           org.AdminName,
		CreatedAt:           org.CreatedAt.Format(time.RFC3339),
	}
}
