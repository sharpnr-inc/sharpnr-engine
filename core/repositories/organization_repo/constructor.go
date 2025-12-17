package organizationrepositories

import (
	"context"
	"sharpnr-engine/core/interfaces"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// organizationRepository is the concrete implementation of the
// OrganizationRepository interface. It provides methods to interact
// with the organization's data in both the main database (Gorm) and Redis.
type OrganizationRepository struct {
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
func NewOrganizationRepo(db *gorm.DB, redis *redis.Client) *OrganizationRepository {
	return &OrganizationRepository{
		redis: redis,
		db:    db,
	}
}

func (d *OrganizationRepository) WithTx(
	ctx context.Context,
	fn func(orgRepoTx interfaces.OrganizationRepositoryInterface) error,
) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// Create a NEW repository instance bound to the transaction
		orgRepoTx := &OrganizationRepository{
			db:    tx,      // 🔥 transaction-scoped DB
			redis: d.redis, // reuse same redis client
		}

		// Execute caller logic
		return fn(orgRepoTx)
	})
}
