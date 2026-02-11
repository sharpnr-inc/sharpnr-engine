package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BootstrapModules(
	db *gorm.DB,
	redis *redis.Client,
	txManager dbtx.TxManager,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) *AppModules {

	organizationService := bootstrapOrganization(db, redis, mailerEngine)
	orgPropertyService := bootstrapOrgProperty(db)
	orgLegalEntityService := bootstrapOrgLegalEntity(db)

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

	hospitality := bootstrapHospitality(db, txManager)

	accountingService := bootstrapAccounting(db, txManager)

	orchestratorService := bootstrapSharpnrOrchestrator(
		organizationService,
		orgPropertyService,
		orgSetupWizardService,
		hospitality,
		accountingService,
		orgLegalEntityService,
		txManager,
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
	}

}
