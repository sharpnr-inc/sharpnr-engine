package app

import (
	"context"

	dbtx "sharpnr-engine/core/dbTx"
	"sharpnr-engine/sere"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BootstrapModules(
	db *gorm.DB,
	redis *redis.Client,
	txManager dbtx.TxManager,
	sereEngine *sere.SERE,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) *AppModules {

	organizationService := bootstrapOrganization(db, redis)
	orgPropertyService := bootstrapOrgProperty(db)
	orgLegalEntityService := bootstrapOrgLegalEntity(db)

	employeeService := bootstrapEmployee(db)
	orgSignupService := bootstrapOrgSignup(db, redis, txManager, sereEngine, mailerEngine, employeeService)
	identityService := bootstrapIdentity(db)

	hospitality := bootstrapHospitality(db, sereEngine)
	accountingService := bootstrapAccounting(db)

	orgSetupWizardService := bootstrapOrgSetupWizard(db)
	orchestratorService := bootstrapSharpnrOrchestrator()

	sereEngine.UseTxExecutor(func(ctx context.Context, fn func(ctx context.Context) error) error {
		return txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)
			return fn(txCtx)
		})
	})

	RegisterAllWorkFlow(
		sereEngine,
	)

	RegisterSereEventHandler(sereEngine)

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
