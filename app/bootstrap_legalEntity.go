package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	org_legal_entity_repo "sharpnr-engine/core/org_legal_entity/repo"
	org_legal_entity_service "sharpnr-engine/core/org_legal_entity/service"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

func bootstrapOrgLegalEntity(
	db *gorm.DB,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) org_legal_entity_service.OrgLegalEntityServiceInterface {

	orgLegalEntityService := org_legal_entity_repo.NewOrgLegalEntityRepo(db)
	return org_legal_entity_service.NewOrgLegalEntityService(orgLegalEntityService, txManager, mailerEngine)
}
