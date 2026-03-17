package app

import (
	org_customer_repo "sharpnr-engine/core/customer/repo"
	org_customer_service "sharpnr-engine/core/customer/service"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	"sharpnr-engine/sero"
	invoice_engine_repos "sharpnr-engine/service_engines/invoicing_engine/repos"
	invoice_engine_services "sharpnr-engine/service_engines/invoicing_engine/services"
	tax_engine_service "sharpnr-engine/service_engines/tax_engine/service"

	"gorm.io/gorm"
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

func bootStrapCustomerEngine(
	db *gorm.DB,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) org_customer_service.OrgCustomerServiceInterface {
	repo := org_customer_repo.NewCustomerRepository(db)
	return org_customer_service.NewOrgCustomerService(
		repo,
		txManager,
		seroEngine,
	)
}

func bootStrapInvoicingEngine(
	db *gorm.DB,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) invoice_engine_services.InvoiceEngineServiceInterface {
	repo := invoice_engine_repos.NewInvoiceEngineRepo(db)
	return invoice_engine_services.NewInvoicingEngineService(
		repo,
		txManager,
		seroEngine,
	)
}
