package tax_engine_service

import (
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
)

type taxEngineService struct {
	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface
	txManager          dbtx.TxManager
	seroEngine         *sero.SEROEngine
}

func NewTaxEngine(
	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) *taxEngineService {
	return &taxEngineService{
		legalEntityService: legalEntityService,
		txManager:          txManager,
		seroEngine:         seroEngine,
	}
}
