package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	"sharpnr-engine/sero"
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

	invoiceEngine := bootStrapInvoicingEngine(db, txManager, seroEngine)

	taxEngineService := bootstrapTaxEngine(orgLegalEntityService, txManager, seroEngine)

	hospitality := bootstrapHospitality(
		db,
		txManager,
		seroEngine,
		sseEngine,
		taxEngineService,
		orgCustomerService,
		invoiceEngine,
	)

	accountingService := bootstrapAccounting(db, txManager, taxEngineService)

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
