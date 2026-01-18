package app

import (
	organization_repo "sharpnr-engine/core/organization/repo"
	organization_services "sharpnr-engine/core/organization/services"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func bootstrapOrganization(
	db *gorm.DB,
	redis *redis.Client,
) organization_services.OrganizationPublicServiceInterface {
	organizationRepo := organization_repo.NewOrganizationRepository(db)
	organizationRedisRepo := organization_repo.NewRedisOrganizationRepository(redis)

	return organization_services.NewOrganizationService(
		organizationRedisRepo,
		organizationRepo,
	)
}
