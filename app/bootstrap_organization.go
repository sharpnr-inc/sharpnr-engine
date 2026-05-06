package app

import (
	organization_repo "sharpnr-engine/core/organization/repo"
	organization_services "sharpnr-engine/core/organization/services"
	"sharpnr-engine/sharp_mailer_engine"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func bootstrapOrganization(
	db *gorm.DB,
	redis *redis.Client,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
) organization_services.OrganizationServiceInterface {
	organizationRepo := organization_repo.NewOrganizationRepository(db)
	organizationRedisRepo := organization_repo.NewRedisOrganizationRepository(redis)

	return organization_services.NewOrganizationService(
		mailerEngine,
		organizationRedisRepo,
		organizationRepo,
	)
}
