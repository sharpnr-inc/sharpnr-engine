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

	identityService := bootstrapIdentity(db, employeeService)

	hospitality := bootstrapHospitality(db, txManager, seroEngine, sseEngine)

	accountingService := bootstrapAccounting(db, txManager)

	orchestratorService := bootstrapSharpnrOrchestrator(
		organizationService,
		orgPropertyService,
		orgSetupWizardService,
		hospitality,
		accountingService,
		orgLegalEntityService,
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
	}

}
