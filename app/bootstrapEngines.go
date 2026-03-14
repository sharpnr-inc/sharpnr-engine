package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	"sharpnr-engine/sero"
	tax_engine_service "sharpnr-engine/service_engines/tax_engine/service"
)

func bootstrapTaxEngine(
	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) tax_engine_service.TaxEngineServiceInterface {
	return tax_engine_service.NewTaxEngine(
		legalEntityService,
		txManager,
		seroEngine,
	)
}
