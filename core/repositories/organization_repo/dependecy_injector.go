package organizationrepositories

import (
	"sharpnr-engine/shared/database"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// organizationRepository is the concrete implementation of the
// OrganizationRepository interface. It provides methods to interact
// with the organization's data in both the main database (Gorm) and Redis.
type organizationRepository struct {
	// redis client for temporary data storage, OTPs, etc.
	redis *redis.Client

	// db is the Gorm database connection for persistent storage
	db *gorm.DB
}

// NewOrganizationRepo creates and initializes a new instance of organizationRepository.
//
// Parameters:
//   - db: Gorm DB instance for interacting with the main database.
//   - redis: Redis client for temporary storage.
//
// Returns:
//   - Pointer to an initialized organizationRepository.
func NewOrganizationRepo(db *gorm.DB, redis *redis.Client) *organizationRepository {
	return &organizationRepository{
		redis: database.RedisClient, // Currently using global Redis client
		db:    database.DB,          // Currently using global DB instance
	}
}
