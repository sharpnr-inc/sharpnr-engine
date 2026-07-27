package app

import (
	"sharpnr-engine/config"
	organization_repo "sharpnr-engine/core/organization/repo"
	organization_services "sharpnr-engine/core/organization/services"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sharp_mailer_engine"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func bootstrapOrganization(
	db *gorm.DB,
	redis *redis.Client,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	txManager dbtx.TxManager,
	ttl config.TTLConfig,

) organization_services.OrganizationServiceInterface {
	organizationRepo := organization_repo.NewOrganizationRepository(db)
	organizationRedisRepo := organization_repo.NewRedisOrganizationRepository(redis, ttl.SignupStaging)

	return organization_services.NewOrganizationService(
		mailerEngine,
		organizationRedisRepo,
		organizationRepo,
		txManager,
	)
}
