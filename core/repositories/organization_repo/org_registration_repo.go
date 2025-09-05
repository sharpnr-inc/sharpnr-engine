package organizationrepositories

import (
	// "github.com/google/uuid"

	"context"
	"encoding/json"
	"fmt"
	"sharpnr-engine/core/models"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

//-----------------------------------------------------------------------------------------------------------------//

func (r *initOrganizationRepository) GetOrgFromRedisDB(ctx context.Context, email string) (*models.TempOrgRegistrationData, error) {
	key := fmt.Sprintf("init_reg:%s", email)
	data, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("no organisation founf for email : %s", email)
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

func (r *initOrganizationRepository) FindBySharpnrOrgID(id string) (bool, error) {
	return false, nil
}

//-----------------------------------------------------------------------------------------------------------------//

func (r *initOrganizationRepository) CreateOrgInMainDB(ctx context.Context, Org *models.Organization) error {
	if err := r.db.WithContext(ctx).Create(Org).Error; err != nil {
		return err
	}
	return nil
}

//-----------------------------------------------------------------------------------------------------------------//

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

//-----------------------------------------------------------------------------------------------------------------//

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

//-----------------------------------------------------------------------------------------------------------------//

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

//-----------------------------------------------------------------------------------------------------------------//

func (r *initOrganizationRepository) GetOrgUUID_By_SharpnrID(sharpnr_id string) (uuid.UUID, error) {
	// This is the function finding the orgId and return uuid
	var org models.Organization
	if err := r.db.Where("sharpnr_org_id = ?", sharpnr_id).First(&org).Error; err != nil {
		return uuid.Nil, fmt.Errorf("organization not found: %w", err)
	}
	return org.ID, nil
}
