package interfaces

import "github.com/sharpnr-inc/sharpnr-engine/core/models"

type CoreOrganizationRepository interface {
	Create(org *models.Organization) error
	Update()
	FindByUUID()
	FindBySharpnrOrgID(id string) (bool, error)
	FindByMail()
}
