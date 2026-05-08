package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	"sharpnr-engine/sero"
	accounting_sero_event_handler "sharpnr-engine/service_engines/accounting_engine/sero_handler"
	inventory_engine_sero_handlers "sharpnr-engine/service_engines/inventory_engine/sero_handlers"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
	sse_engine "sharpnr-engine/sse"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BootstrapModules(
	db *gorm.DB,
	redis *redis.Client,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	seroEngine *sero.SEROEngine,
	sseEngine *sse_engine.StreamManager,
) *AppModules {

	// Sentinel Engine uses the DB and used to Evaluate the authorization
	// sentinel_engine := bootStrapSentinelEngine(db)

	organizationService := bootstrapOrganization(db, redis, mailerEngine, txManager)
	orgPropertyService := bootstrapOrgProperty(db, txManager, mailerEngine)
	orgLegalEntityService := bootstrapOrgLegalEntity(db, txManager, mailerEngine)

	/* ---------------------------- inventory engine ---------------------------- */
	inventoryEngine := bootstrapInventoryEngine(db, txManager, seroEngine)
	inventory_engine_sero_handlers.RegisterSeroHandler(seroEngine, inventoryEngine)
	/* ---------------------------- inventory engine ---------------------------- */

	employeesIdentityService := bootstrapEmployeeIdentity(db, seroEngine, txManager, mailerEngine)
	employeesHRService := bootstrapEmployeeHR(db, seroEngine, txManager, mailerEngine)

	orgCustomerService := bootStrapCustomerEngine(db, txManager, seroEngine)
	orgSupplierService := bootstrapSupplier(db, txManager, mailerEngine, seroEngine)

	identityService := bootstrapIdentity(employeesIdentityService)

	orgSetupWizardService := bootstrapOrgSetupWizard(
		db,
		orgLegalEntityService,
		orgPropertyService,
		organizationService,
		identityService,
		txManager,
		mailerEngine,
	)

	orgSignupService := bootstrapOrgSignup(
		txManager,
		mailerEngine,
		employeesIdentityService,
		organizationService,
		orgSetupWizardService,
	)

	paymentsEngine := bootstrapPaymentEngine(db, mailerEngine, txManager, seroEngine)

	taxEngineService := bootstrapTaxEngine(orgLegalEntityService, txManager, seroEngine)
	invoiceEngine := bootStrapInvoicingEngine(db, txManager, seroEngine, taxEngineService, orgLegalEntityService)

	hospitality := bootstrapHospitality(
		db,
		txManager,
		seroEngine,
		sseEngine,
		taxEngineService,
		orgCustomerService,
		invoiceEngine,
		paymentsEngine,
		inventoryEngine,
	)

	hospitalityOrchestration := bootstrapHospitalityDomainOrchestration(
		orgPropertyService,
		hospitality,
	)

	/* ---------------------------- Accounting Engine --------------------------- */
	accountingService := bootstrapAccounting(
		db,
		txManager,
		taxEngineService,
		invoiceEngine,
		paymentsEngine,
		inventoryEngine,
		orgSupplierService,
		seroEngine,
	)
	accounting_sero_event_handler.Register(seroEngine, accountingService)
	/* ---------------------------- Accounting Engine --------------------------- */

	orchestratorService := bootstrapSharpnrOrchestrator(
		organizationService,
		orgPropertyService,
		orgSetupWizardService,
		hospitality,
		accountingService,
		orgLegalEntityService,
		orgCustomerService,
		txManager,
		seroEngine,
		paymentsEngine,
		inventoryEngine,
		orgSupplierService,
	)

	return &AppModules{
		organizationService:   organizationService,
		orgSignupService:      orgSignupService,
		orgPropertyService:    orgPropertyService,
		orgLegalEntityService: orgLegalEntityService,

		employeesIdentityService: employeesIdentityService,
		employeesHRService:       employeesHRService,

		accountingService:        accountingService,
		sharpAuthService:         identityService,
		hospitality:              hospitality,
		hospitalityOrchestration: hospitalityOrchestration,
		orchestratorService:      orchestratorService,
		orgSetupWizard:           orgSetupWizardService,
		sseEngine:                sseEngine,
		taxEngineService:         taxEngineService,
		orgCustomerService:       orgCustomerService,
	}

}
