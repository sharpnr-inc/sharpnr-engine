package organizationrepositories

import (
	"sharpnr-engine/shared/database"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type initOrganizationRepository struct {
	redis *redis.Client
	db    *gorm.DB
}

func InitializeNewOrgRegistration(db *gorm.DB, redis *redis.Client) *initOrganizationRepository {
	return &initOrganizationRepository{
		redis: database.RedisClient,
		db:    database.DB,
	}
}

// /---Org Login----///
type orgLoginRepo struct {
	db *gorm.DB
}

// This repo for the Login
func OrgLogin(db *gorm.DB) *orgLoginRepo {
	return &orgLoginRepo{
		db: database.DB,
	}
}
