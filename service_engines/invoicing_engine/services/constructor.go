package invoice_engine_services

import (
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	invoice_engine_repos "sharpnr-engine/service_engines/invoicing_engine/repos"
	tax_engine_service "sharpnr-engine/service_engines/tax_engine/service"
)

type invoiceEngineService struct {
	txManager          dbtx.TxManager
	seroEngine         *sero.SEROEngine
	invoicingRepo      invoice_engine_repos.InvoicingEngineRepositoryInterface
	taxEngine          tax_engine_service.TaxEngineServiceInterface
	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface
}

func NewInvoicingEngineService(
	invoicingRepo invoice_engine_repos.InvoicingEngineRepositoryInterface,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
	taxEngine tax_engine_service.TaxEngineServiceInterface,
	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface,
) *invoiceEngineService {
	return &invoiceEngineService{
		invoicingRepo:      invoicingRepo,
		txManager:          txManager,
		seroEngine:         seroEngine,
		taxEngine:          taxEngine,
		legalEntityService: legalEntityService,
	}
}
