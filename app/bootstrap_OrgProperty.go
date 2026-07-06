package app

import (
	org_property_repo "sharpnr-engine/core/org_property/repo"
	org_property_services "sharpnr-engine/core/org_property/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	sentinelEngine "sharpnr-engine/sentinel_engine"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

func bootstrapOrgProperty(
	db *gorm.DB,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	sentinelAuthzEngine sentinelEngine.SentinelEngineInterface,
) org_property_services.OrganizationPropertyServicesInterfaces {
	orgPropertyRepo := org_property_repo.NewOrganizationPropertyRepo(db)
	return org_property_services.NewOrganizationPropertyService(orgPropertyRepo, txManager, mailerEngine, sentinelAuthzEngine)

}
