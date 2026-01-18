package organization_repo

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type organizationRepo struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *organizationRepo {
	return &organizationRepo{
		db: db,
	}
}

type organizationRedisRepo struct {
	redis *redis.Client
}

func NewRedisOrganizationRepository(redis *redis.Client) *organizationRedisRepo {
	return &organizationRedisRepo{
		redis: redis,
	}
}
