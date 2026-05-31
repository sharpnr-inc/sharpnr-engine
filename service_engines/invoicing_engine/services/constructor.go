package invoice_engine_services

import (
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	invoice_engine_repos "sharpnr-engine/service_engines/invoicing_engine/repos"
)

type invoiceEngineService struct {
	txManager     dbtx.TxManager
	seroEngine    *sero.SEROEngine
	invoicingRepo invoice_engine_repos.InvoicingEngineRepositoryInterface

	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface
}

func NewInvoicingEngineService(
	invoicingRepo invoice_engine_repos.InvoicingEngineRepositoryInterface,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,

	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface,
) *invoiceEngineService {
	return &invoiceEngineService{
		invoicingRepo:      invoicingRepo,
		txManager:          txManager,
		seroEngine:         seroEngine,
		legalEntityService: legalEntityService,
	}
}
