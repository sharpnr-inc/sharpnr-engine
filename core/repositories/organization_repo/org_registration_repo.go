package organizationrepositories

import (
	"context"
	"encoding/json"
	"fmt"
	"sharpnr-engine/core/models"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

//-----------------------------------------------------------------------------------------------------------------//

func (r *organizationRepository) EmailOrPhoneExistsInRedis(ctx context.Context, email, phone string) (bool, error) {
	emailKey := fmt.Sprintf("init_reg:%s", email)
	phoneKey := fmt.Sprintf("init_reg_phone:%s", phone)

	if _, err := r.redis.Get(ctx, emailKey).Result(); err == nil {
		return true, nil
	} else if err != redis.Nil {
		return false, err
	}

	if _, err := r.redis.Get(ctx, phoneKey).Result(); err == nil {
		return true, nil
	} else if err != redis.Nil {
		return false, err
	}

	return false, nil
}

//-----------------------------------------------------------------------------------------------------------------//

func (r *organizationRepository) CreateOrgInRedisDB(ctx context.Context, tempOrg *models.TempOrgRegistrationData) error {
	key := fmt.Sprintf("init_reg:%s", tempOrg.Email)

	// Serialize struct to JSON
	data, err := json.Marshal(tempOrg)
	if err != nil {
		return err
	}

	// Save in Redis with a TTL (e.g. 15 minutes)
	return r.redis.Set(ctx, key, data, 15*time.Minute).Err()
}

//-----------------------------------------------------------------------------------------------------------------//

func (r *organizationRepository) GetOrgFromRedisDB(ctx context.Context, email string) (*models.TempOrgRegistrationData, error) {
	key := fmt.Sprintf("init_reg:%s", email)
	data, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("no organisation found for email : %s", email)
		}
		return nil, err
	}

	var tempOrg models.TempOrgRegistrationData
	if err := json.Unmarshal([]byte(data), &tempOrg); err != nil {
		return nil, err
	}

	return &tempOrg, nil
}

//-----------------------------------------------------------------------------------------------------------------//

func (r *organizationRepository) DeleteOrgFromRedisDB(ctx context.Context, email string) error {
	emailKey := fmt.Sprintf("init_reg:%s", email)
	return r.redis.Del(ctx, emailKey).Err()
}

//-----------------------------------------------------------------------------------------------------------------//

func (r *organizationRepository) EmailOrPhoneExistsInMainDB(ctx context.Context, email string, phone string) (bool, error) {
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

//-----------------------------------------------------------------------------------------------------------------//

func (r *organizationRepository) CreateOrgInMainDB(ctx context.Context, Org *models.Organization) error {
	if err := r.db.WithContext(ctx).Create(Org).Error; err != nil {
		return err
	}
	return nil
}

//-----------------------------------------------------------------------------------------------------------------//

func (r *organizationRepository) FindOrgBySharpnrID(ctx context.Context, id string) (*models.Organization, bool, error) {
	var org models.Organization

	if err := r.db.WithContext(ctx).
		Where("sharpnr_org_id = ?", id).
		First(&org).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			// Not found
			return nil, false, nil
		}
		// Real DB error
		return nil, false, err
	}

	// Found
	return &org, true, nil
}

//-----------------------------------------------------------------------------------------------------------------//

func (r *organizationRepository) GetOrgUUID_By_SharpnrID(sharpnr_id string) (uuid.UUID, error) {
	// This is the function finding the orgId and return uuid
	var org models.Organization
	if err := r.db.Where("sharpnr_org_id = ?", sharpnr_id).First(&org).Error; err != nil {
		return uuid.Nil, fmt.Errorf("organization not found: %w", err)
	}
	return org.ID, nil
}
