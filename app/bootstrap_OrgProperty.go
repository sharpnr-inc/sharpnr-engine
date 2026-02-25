package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	org_property_repo "sharpnr-engine/core/org_property/repo"
	org_property_services "sharpnr-engine/core/org_property/services"

	"gorm.io/gorm"
)

func bootstrapOrgProperty(
	db *gorm.DB,
	txManager dbtx.TxManager,
) org_property_services.OrganizationPropertyServicesInterfaces {
	orgPropertyRepo := org_property_repo.NewOrganizationPropertyRepo(db)
	return org_property_services.NewOrganizationPropertyService(orgPropertyRepo, txManager)

}
