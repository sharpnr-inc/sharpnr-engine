package app

import (
	"sharpnr-engine/config"
	dbtx "sharpnr-engine/pkg/dbTx"
	jwt_manager "sharpnr-engine/pkg/jwtmanager"
	"sharpnr-engine/sero"
	accounting_sero_event_handler "sharpnr-engine/sharp_accounting_engine/accounting_engine/sero_handler"
	stockwaveEngineSeroHandlers "sharpnr-engine/stockwave_engine/sero_handlers"

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
	jwtmgr jwt_manager.Manager,
	ttl config.TTLConfig,
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

	organizationService := bootstrapOrganization(db, redis, mailerEngine, txManager, ttl)
	orgPropertyService := bootstrapOrgProperty(db, txManager, mailerEngine, sentinelAuthzEngine)

	// Departments are bootstrapped ahead of the aggregator because their
	// mod_resources provider lists the org's departments for the policy builder.
	orgDepartmentService := bootstrapOrgDepartments(db, txManager, mailerEngine, seroEngine, auditLogService)

	// Resources aggregator: every module's mod_resources, for the FE policy builder.
	sentinelAggregator := bootstrapSentinelAggregator(orgPropertyService, sentinelService, orgDepartmentService)
	orgLegalEntityService := bootstrapOrgLegalEntity(db, txManager, mailerEngine)

	/* ---------------------------- inventory engine ---------------------------- */
	inventoryEngine := bootstrapInventoryEngine(db, txManager, seroEngine)
	stockwaveEngineSeroHandlers.RegisterSeroHandler(seroEngine, inventoryEngine)
	/* ---------------------------- inventory engine ---------------------------- */

	// One core employee service instance, handed to each consumer through its
	// own segregated interface.
	employeesAuthneticationService,
		employeesHRService,
		employeesIamEngineService,
		employeeSignUpService,
		employeeSetUpWizService, employeeOrchestratorService := bootstrapEmployeeServices(db, seroEngine, txManager, mailerEngine, sentinelService, sentinelAuthzEngine, auditLogService, jwtmgr)

	/* ------------------------------- IAM engine ------------------------------- */
	iamEngineService := bootstrapIamEngine(employeesIamEngineService)
	/* ------------------------------- IAM engine ------------------------------- */

	/* ------------------------------ HRMS engine ------------------------------ */
	hrmsService := bootstrapHRMSService(seroEngine, txManager, mailerEngine, employeesHRService)
	/* ------------------------------ HRMS engine ------------------------------ */

	orgCustomerService := bootStrapCustomerEngine(db, txManager, seroEngine)
	orgSupplierService := bootstrapSupplier(db, txManager, mailerEngine, seroEngine)

	authenticationService := bootstrapAuthentication(employeesAuthneticationService, organizationService, sentinelService, auditLogService)

	orgSetupWizardService := bootstrapOrgSetupWizard(
		db,
		orgLegalEntityService,
		orgPropertyService,
		organizationService,
		employeeSetUpWizService,
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
		ttl.SignupStaging,
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
		employeeOrchestratorService,
		sentinelService,
		sentinelAggregator,
		sentinelAuthzEngine,
		auditLogService,
		orgDepartmentService,
	)

	utilitiesApiService := bootstrapUtilitiesApiService(db)

	return &AppModules{
		auditLogService:                auditLogService,
		organizationService:            organizationService,
		orgSignupService:               orgSignupService,
		orgPropertyService:             orgPropertyService,
		orgLegalEntityService:          orgLegalEntityService,
		employeesAuthneticationService: employeesAuthneticationService,
		employeesHRService:             employeesHRService,
		iamEngineService:               iamEngineService,
		hrmsService:                    hrmsService,
		accountingService:              accountingService,
		authenticationService:          authenticationService,
		hospitality:                    hospitality,
		hospitalityOrchestration:       hospitalityOrchestration,
		orchestratorService:            orchestratorService,
		orgSetupWizard:                 orgSetupWizardService,
		sseEngine:                      sseEngine,
		orgCustomerService:             orgCustomerService,
		utilitiesApiService:            utilitiesApiService,
	}

}
