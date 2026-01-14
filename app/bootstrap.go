package app

import (
	"context"

	dbtx "sharpnr-engine/core/dbTx"
	"sharpnr-engine/sere"
	sharp_mailer_engine "sharpnr-engine/sharp-mailer-engine"

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

	orgSignupService := bootstrapOrgSignup(db, redis, txManager, sereEngine, mailerEngine)
	hr := bootstrapHR(db)
	identityService := bootstrapIdentity(db, hr)
	hospitality := bootstrapHospitality(db, sereEngine)
	accountingModule := bootstrapAccounting(db)

	sereEngine.UseTxExecutor(func(ctx context.Context, fn func(ctx context.Context) error) error {
		return txManager.WithTx(ctx, func(tx *gorm.DB) error {
			txCtx := dbtx.InjectTx(ctx, tx)
			return fn(txCtx)
		})
	})

	RegisterAllWorkFlow(
		sereEngine,
		hr.Service,
		identityService,

		hospitality.KitchenRepo(),
		hospitality.HotelRepo(),
		hospitality.HotelService(),

		hospitality.BarRepo(),
		hospitality.BarService(),

		hospitality.RestaurantRepo(),
		hospitality.RestaurantService(),

		hospitality.MenuRepo(),
		hospitality.MenuService(),

		accountingModule.Repo,
		accountingModule.Service,
	)

	RegisterSereEventHandler(sereEngine, accountingModule.Repo)

	return &AppModules{
		orgSignupService:  orgSignupService,
		hrService:         hr.Service,
		accountingService: accountingModule.Service,
		sharpAuthService:  identityService,
		hospitality:       hospitality,
	}

}
