package app

import (
	org_customer_repo "sharpnr-engine/core/customer/repo"
	org_customer_service "sharpnr-engine/core/customer/service"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	payment_gateway "sharpnr-engine/payment_engine/gateway"
	"sharpnr-engine/payment_engine/gateway/cash"
	"sharpnr-engine/payment_engine/gateway/phonepe"
	payment_engine_repo "sharpnr-engine/payment_engine/repo"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	"sharpnr-engine/sero"
	inventory_engine_repo "sharpnr-engine/service_engines/inventory_engine/repo"
	inventory_engine_services "sharpnr-engine/service_engines/inventory_engine/services"
	invoice_engine_repos "sharpnr-engine/service_engines/invoicing_engine/repos"
	invoice_engine_services "sharpnr-engine/service_engines/invoicing_engine/services"
	tax_engine_service "sharpnr-engine/service_engines/tax_engine/service"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

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
	taxEngine tax_engine_service.TaxEngineServiceInterface,
	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface,
) invoice_engine_services.InvoiceEngineServiceInterface {
	repo := invoice_engine_repos.NewInvoiceEngineRepo(db)
	return invoice_engine_services.NewInvoicingEngineService(
		repo,
		txManager,
		seroEngine,
		taxEngine,
		legalEntityService,
	)
}

func bootstrapPaymentEngine(
	db *gorm.DB,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) payment_engine_service.PaymentServiceInterface {
	repo := payment_engine_repo.NewPaymentRepository(db)
	// 1. Initialize the logger (Structured, Plain Text, No Color)

	registry := payment_gateway.NewProviderRegistry(
		cash.NewCashProvider(),
		phonepe.NewPhonePeProvider(),
	)
	return payment_engine_service.NewPaymentService(
		repo,
		registry,
		mailerEngine,
		txManager,
		seroEngine,
	)
}

func bootstrapInventoryEngine(
	db *gorm.DB,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) inventory_engine_services.InventoryEngineServicesInterface {
	repo := inventory_engine_repo.NewInventoryEngineRepo(db)
	return inventory_engine_services.NewInventoryEngineService(
		repo,
		txManager,
		seroEngine,
	)
}
