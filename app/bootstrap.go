package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	"sharpnr-engine/sero"
	accounting_sero_event_handler "sharpnr-engine/service_engines/accounting_engine/sero_event_handler"
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

	iAM := bootstrapSharIAM(db)
	organizationService := bootstrapOrganization(db, redis, mailerEngine, iAM)
	orgPropertyService := bootstrapOrgProperty(db, txManager)
	orgLegalEntityService := bootstrapOrgLegalEntity(db, txManager, mailerEngine)

	employeeService := bootstrapEmployee(db)
	orgSetupWizardService := bootstrapOrgSetupWizard(
		db,
		orgLegalEntityService,
		orgPropertyService,
		organizationService,
	)

	orgSignupService := bootstrapOrgSignup(
		txManager,
		mailerEngine,
		employeeService,
		organizationService,
		orgSetupWizardService,
	)

	orgCustomerService := bootStrapCustomerEngine(db, txManager, seroEngine)

	identityService := bootstrapIdentity(db, employeeService)

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
	)

	accountingService := bootstrapAccounting(db, txManager, taxEngineService, invoiceEngine, paymentsEngine)
	accounting_sero_event_handler.Register(seroEngine, accountingService)

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
	)

	return &AppModules{
		organizationService:   organizationService,
		orgSignupService:      orgSignupService,
		orgPropertyService:    orgPropertyService,
		orgLegalEntityService: orgLegalEntityService,
		employeeService:       employeeService,
		accountingService:     accountingService,
		sharpAuthService:      identityService,
		hospitality:           hospitality,
		orchestratorService:   orchestratorService,
		orgSetupWizard:        orgSetupWizardService,
		sseEngine:             sseEngine,
		taxEngineService:      taxEngineService,
		orgCustomerService:    orgCustomerService,
	}

}
