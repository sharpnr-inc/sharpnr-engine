package app

import (
	dbtx "sharpnr-engine/core/dbTx"
	org_signup_services "sharpnr-engine/core/org_signup/services"
	"sharpnr-engine/sere"
	sharp_mailer_engine "sharpnr-engine/sharp_mailer_engine"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func bootstrapOrgSignup(
	db *gorm.DB,
	redis *redis.Client,
	txManager dbtx.TxManager,
	sereEngine *sere.SERE,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) org_signup_services.OrganizationSignUpServiceInterfaces {

	return org_signup_services.NewOrgSignUpService(
		txManager,
		sereEngine,
		mailerEngine,
	)
}
