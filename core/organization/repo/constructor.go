package organization_repo

import (
	"time"

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

	// stagingTTL is how long a signup staging record — and its email/phone
	// reverse-lookup keys — live. Injected from config (SIGNUP_STAGING_TTL)
	// rather than fixed here, so the signup/OTP window is tunable per
	// environment instead of requiring a rebuild.
	stagingTTL time.Duration
}

func NewRedisOrganizationRepository(
	redis *redis.Client,
	stagingTTL time.Duration,
) *organizationRedisRepo {
	return &organizationRedisRepo{
		redis:      redis,
		stagingTTL: stagingTTL,
	}
}
