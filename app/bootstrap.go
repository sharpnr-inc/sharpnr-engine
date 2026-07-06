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

	// Audit log engine: records every tenant event. Bootstrapped first so it
	// can be injected into every other engine as they are wired below.
	auditLogService := bootstrapAuditLogEngine(db)

	// Sentinel authorization engine: shared, concurrency-safe Cedar evaluator.
	// Injected into the orchestrator for authorization decisions and into the
	// authoring service as its cache invalidator.
	sentinelAuthzEngine := bootstrapSentinelEngine(db)

	// Sentinel authoring service: policy/role/group CRUD + attach/detach.
	sentinelService := bootstrapSentinel(db, txManager, sentinelAuthzEngine)

	organizationService := bootstrapOrganization(db, redis, mailerEngine, txManager)
	orgPropertyService := bootstrapOrgProperty(db, txManager, mailerEngine, sentinelAuthzEngine)

	// Resources aggregator: every module's mod_resources, for the FE policy builder.
	sentinelAggregator := bootstrapSentinelAggregator(orgPropertyService, sentinelService)
	orgLegalEntityService := bootstrapOrgLegalEntity(db, txManager, mailerEngine)

	/* ---------------------------- inventory engine ---------------------------- */
	inventoryEngine := bootstrapInventoryEngine(db, txManager, seroEngine)
	inventory_engine_sero_handlers.RegisterSeroHandler(seroEngine, inventoryEngine)
	/* ---------------------------- inventory engine ---------------------------- */

	employeesIdentityService := bootstrapEmployeeAuthentication(db, seroEngine, txManager, mailerEngine)
	employeesHRService := bootstrapEmployeeHR(db, seroEngine, txManager, mailerEngine)
	employeesIamEngineService := bootstrapEmployeeIamEngineService(db, seroEngine, txManager, mailerEngine)
	employeeSignUpService := bootstrapEmployeeSignUpService(db, seroEngine, txManager, mailerEngine)

	/* ------------------------------- IAM engine ------------------------------- */
	iamEngineService := bootstrapIamEngine(employeesIamEngineService)
	/* ------------------------------- IAM engine ------------------------------- */

	/* ------------------------------ HRMS engine ------------------------------ */
	hrmsEngine := bootstrapHRMSEngine(db, seroEngine, txManager, mailerEngine, employeesHRService)
	/* ------------------------------ HRMS engine ------------------------------ */

	orgCustomerService := bootStrapCustomerEngine(db, txManager, seroEngine)
	orgSupplierService := bootstrapSupplier(db, txManager, mailerEngine, seroEngine)

	authenticationService := bootstrapAuthentication(employeesIdentityService, organizationService, sentinelService)

	orgSetupWizardService := bootstrapOrgSetupWizard(
		db,
		orgLegalEntityService,
		orgPropertyService,
		organizationService,
		authenticationService,
		txManager,
		mailerEngine,
	)

	orgSignupService := bootstrapOrgSignup(
		txManager,
		mailerEngine,
		employeeSignUpService,
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
		sentinelAggregator,
		sentinelAuthzEngine,
	)

	utilitiesApiService := bootstrapUtilitiesApiService(db)

	return &AppModules{
		auditLogService:          auditLogService,
		organizationService:      organizationService,
		orgSignupService:         orgSignupService,
		orgPropertyService:       orgPropertyService,
		orgLegalEntityService:    orgLegalEntityService,
		employeesIdentityService: employeesIdentityService,
		employeesHRService:       employeesHRService,
		iamEngineService:         iamEngineService,
		hrmsEngine:               hrmsEngine,
		accountingService:        accountingService,
		authenticationService:    authenticationService,
		hospitality:              hospitality,
		hospitalityOrchestration: hospitalityOrchestration,
		orchestratorService:      orchestratorService,
		orgSetupWizard:           orgSetupWizardService,
		sseEngine:                sseEngine,
		orgCustomerService:       orgCustomerService,
		utilitiesApiService:      utilitiesApiService,
	}

}
