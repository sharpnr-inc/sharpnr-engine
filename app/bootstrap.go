package app

import (
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	accounting_sero_event_handler "sharpnr-engine/sharp_accounting_engine/accounting_engine/sero_handler"
	inventory_engine_sero_handlers "sharpnr-engine/sharp_inventory_engine/sero_handlers"

	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"
	sse_engine "sharpnr-engine/sse_engine"

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

	// Sentinel authorization engine: shared, concurrency-safe. Currently held
	// only as a dormant placeholder for the orchestrator; it is NOT wired into
	// the Sentinel service or its routes yet (authorization comes later).
	// sentinelEngine := bootstrapSentinelEngine(db)

	// Sentinel authoring service: policy/role/group CRUD + attach/detach. No
	// engine dependency.
	sentinelService := bootstrapSentinel(db, txManager)

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

	identityService := bootstrapIdentity(employeesIdentityService, organizationService)

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
		sentinelService,
	)

	paymentsEngine := bootstrapPaymentEngine(db, mailerEngine, txManager, seroEngine)

	invoiceEngine := bootStrapInvoicingEngine(db, txManager, seroEngine, orgLegalEntityService)

	hospitality := bootstrapHospitality(
		db,
		txManager,
		seroEngine,
		sseEngine,
		orgCustomerService,
		invoiceEngine,
		paymentsEngine,
		inventoryEngine,
	)

	hospitalityOrchestration := bootstrapHospitalityDomainOrchestration(
		orgPropertyService,
		hospitality,
		orgLegalEntityService,
	)

	/* ---------------------------- Accounting Engine --------------------------- */
	accountingService := bootstrapAccounting(
		db,
		txManager,
		invoiceEngine,
		paymentsEngine,
		inventoryEngine,
		orgSupplierService,
		seroEngine,
		orgLegalEntityService,
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
		employeesHRService,
		sentinelService,
	)

	utilitiesApiService := bootstrapUtilitiesApiService(db)

	return &AppModules{
		organizationService:      organizationService,
		orgSignupService:         orgSignupService,
		orgPropertyService:       orgPropertyService,
		orgLegalEntityService:    orgLegalEntityService,
		employeesIdentityService: employeesIdentityService,
		employeesHRService:       employeesHRService,
		accountingService:        accountingService,
		sharpAuthService:         identityService,
		hospitality:              hospitality,
		hospitalityOrchestration: hospitalityOrchestration,
		orchestratorService:      orchestratorService,
		orgSetupWizard:           orgSetupWizardService,
		sseEngine:                sseEngine,
		orgCustomerService:       orgCustomerService,
		utilitiesApiService:      utilitiesApiService,
	}

}
