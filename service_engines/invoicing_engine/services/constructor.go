package invoice_engine_services

import (
	dbtx "sharpnr-engine/core/dbTx"
	"sharpnr-engine/sero"
	invoice_engine_repos "sharpnr-engine/service_engines/invoicing_engine/repos"
)

type invoiceEngineService struct {
	txManager     dbtx.TxManager
	seroEngine    *sero.SEROEngine
	invoicingRepo invoice_engine_repos.InvoicingEngineRepositoryInterface
}

func NewInvoicingEngineService(
	invoicingRepo invoice_engine_repos.InvoicingEngineRepositoryInterface,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) *invoiceEngineService {
	return &invoiceEngineService{
		invoicingRepo: invoicingRepo,
		txManager:     txManager,
		seroEngine:    seroEngine,
	}
}
