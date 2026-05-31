package app

import (
	org_customer_repo "sharpnr-engine/core/customer/repo"
	org_customer_service "sharpnr-engine/core/customer/service"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	payment_gateway "sharpnr-engine/payment_engine/gateway"
	"sharpnr-engine/payment_engine/gateway/cash"
	payment_engine_repo "sharpnr-engine/payment_engine/repo"
	payment_engine_service "sharpnr-engine/payment_engine/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	inventory_engine_repo "sharpnr-engine/sharp_inventory_engine/repo"
	inventory_engine_services "sharpnr-engine/sharp_inventory_engine/services"

	invoice_engine_repos "sharpnr-engine/service_engines/invoicing_engine/repos"
	invoice_engine_services "sharpnr-engine/service_engines/invoicing_engine/services"

	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"gorm.io/gorm"
)

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

	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface,
) invoice_engine_services.InvoiceEngineServiceInterface {
	repo := invoice_engine_repos.NewInvoiceEngineRepo(db)
	return invoice_engine_services.NewInvoicingEngineService(
		repo,
		txManager,
		seroEngine,
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
