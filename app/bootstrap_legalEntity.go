package app

import (
	legal_entity_repo "sharpnr-engine/core/legal_entity/repo"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	dbtx "sharpnr-engine/pkg/dbTx"
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
