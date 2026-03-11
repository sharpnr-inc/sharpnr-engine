package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_repo "sharpnr-engine/core/org_legal_entity/repo"
	legal_entity_service "sharpnr-engine/core/org_legal_entity/service"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

func bootstrapOrgLegalEntity(
	db *gorm.DB,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) legal_entity_service.OrgLegalEntityServiceInterface {

	orgLegalEntityService := legal_entity_repo.NewOrgLegalEntityRepo(db)
	return legal_entity_service.NewOrgLegalEntityService(orgLegalEntityService, txManager, mailerEngine)
}
