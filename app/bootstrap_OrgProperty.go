package app

import (
	org_property_repo "sharpnr-engine/core/org_property/repo"
	org_property_services "sharpnr-engine/core/org_property/services"

	"gorm.io/gorm"
)

func bootstrapOrgProperty(
	db *gorm.DB,
) org_property_services.OrganizationPropertyServicesInterfaces {
	orgPropertyRepo := org_property_repo.NewOrganizationPropertyRepo(db)
	return org_property_services.NewOrganizationPropertyService(orgPropertyRepo)

}
