package repositories

import (
	// "github.com/google/uuid"

	"context"
	"encoding/json"
	"fmt"
	"sharpnr-engine/core/models"
	"sharpnr-engine/shared/database"
	"time"

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

func (r *initOrganizationRepository) CreateOrgInMainDB(ctx context.Context, tempOrg *models.TempOrgRegistrationData) error {
	// TODO: insert into PostgreSQL (main DB)
	return nil
}

func (r *initOrganizationRepository) CreateOrgInRedisDB(ctx context.Context, tempOrg *models.TempOrgRegistrationData) error {
	key := fmt.Sprintf("init_reg:%s", tempOrg.Email)

	// Serialize struct to JSON
	data, err := json.Marshal(tempOrg)
	if err != nil {
		return err
	}

	// Save in Redis with a TTL (e.g. 15 minutes)
	return r.redis.Set(ctx, key, data, 15*time.Minute).Err()
}

func (r *initOrganizationRepository) EmailOrPhoneExistsInRedis(ctx context.Context, email string, phone string) (bool, error) {
	emailKey := fmt.Sprintf("init_reg:%s", email)
	phoneKey := fmt.Sprintf("init_reg_phone:%s", phone)

	// check email
	_, err := r.redis.Get(ctx, emailKey).Result()
	if err == nil {
		return true, nil
	}
	if err != redis.Nil {
		return false, err
	}

	// check phone
	_, err = r.redis.Get(ctx, phoneKey).Result()
	if err == nil {
		return true, nil
	}
	if err != redis.Nil {
		return false, err
	}

	return false, nil
}

func (r *initOrganizationRepository) EmailOrPhoneExistsInMainDB(ctx context.Context, email string, phone string) (bool, error) {
	var count int64

	// Check if any record exists with the given email or phone
	err := r.db.WithContext(ctx).Model(&models.Organization{}).
		Where("root_email = ? OR master_phone_number = ?", email, phone).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
