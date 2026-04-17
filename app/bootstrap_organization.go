package app

import (
	organization_repo "sharpnr-engine/core/organization/repo"
	organization_services "sharpnr-engine/core/organization/services"
	sharpIam_services "sharpnr-engine/core/sharp_iam/services"
	"sharpnr-engine/sharp_mailer_engine"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func bootstrapOrganization(
	db *gorm.DB,
	redis *redis.Client,
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	iam *sharpIam_services.SharpIamService,
) organization_services.OrganizationServiceInterface {
	organizationRepo := organization_repo.NewOrganizationRepository(db)
	organizationRedisRepo := organization_repo.NewRedisOrganizationRepository(redis)

	return organization_services.NewOrganizationService(
		mailerEngine,
		organizationRedisRepo,
		organizationRepo,
		iam,
	)
}


