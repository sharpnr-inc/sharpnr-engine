package organizationrepositories

import (
	"context"
	"encoding/json"
	"fmt"
	"sharpnr-engine/core/models"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// -----------------------------------------------------------------------------------------------------------------//
func (r *organizationRepository) EmailOrPhoneExistsInRedis(ctx context.Context, email, phone string) (bool, error) {
	// Build keys with constants for better maintainability
	const (
		emailKeyPrefix = "init_reg:"
		phoneKeyPrefix = "init_reg_phone:"
	)

	keys := make([]string, 0, 2)
	if email != "" {
		keys = append(keys, emailKeyPrefix+email)
	}
	if phone != "" {
		keys = append(keys, phoneKeyPrefix+phone)
	}

	// Handle edge case where both are empty
	if len(keys) == 0 {
		return false, nil
	}

	// Check both keys in a single Redis call for better performance
	existsCount, err := r.redis.Exists(ctx, keys...).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check existence in Redis: %w", err)
	}

	return existsCount > 0, nil
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
	email = strings.TrimSpace(email)
	key := fmt.Sprintf("init_reg:%s", email)
	fmt.Println("Looking for Redis key:", key)

	data, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("no organization found for email: %s", email)
		}
		return nil, err
	}

	var tempOrg models.TempOrgRegistrationData
	if err := json.Unmarshal([]byte(data), &tempOrg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Redis data: %w", err)
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
		Where("sharpnr_id = ?", id).
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

// UpdateOrg updates an existing organization in the main database
func (r *organizationRepository) UpdateOrg(ctx context.Context, org *models.Organization) error {
	if org == nil {
		return fmt.Errorf("organization cannot be nil")
	}

	if org.ID == uuid.Nil {
		return fmt.Errorf("organization ID cannot be empty")
	}

	// Use Save() which updates all fields, or use Updates() for selective updates
	result := r.db.WithContext(ctx).Save(org)
	if result.Error != nil {
		return fmt.Errorf("failed to update organization: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("organization not found with ID: %s", org.ID)
	}

	return nil
}

func (r *organizationRepository) CreateBankAccounts(ctx context.Context, orgID uuid.UUID, accounts []models.OrganizationBankAccount) error {
	return r.db.WithContext(ctx).Create(&accounts).Error
}

func (r *organizationRepository) GetBankAccountsByOrgID(ctx context.Context, orgID uuid.UUID) ([]models.OrganizationBankAccount, error) {
	var accounts []models.OrganizationBankAccount
	if err := r.db.WithContext(ctx).Where("organization_id = ?", orgID).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *organizationRepository) GetPrimaryBankAccount(ctx context.Context, orgID uuid.UUID) (*models.OrganizationBankAccount, error) {
	var account models.OrganizationBankAccount
	if err := r.db.WithContext(ctx).Where("organization_id = ? AND is_primary = ?", orgID, true).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *organizationRepository) UpdateBankAccounts(ctx context.Context, orgID uuid.UUID, accounts []models.OrganizationBankAccount) error {
	// Transaction to prevent orphan data
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Delete old accounts
		if err := tx.Where("organization_id = ?", orgID).Delete(&models.OrganizationBankAccount{}).Error; err != nil {
			return err
		}

		// Insert new accounts
		for i := range accounts {
			accounts[i].OrganizationID = orgID
			if accounts[i].ID == uuid.Nil {
				accounts[i].ID = uuid.New()
			}
		}
		return tx.Create(&accounts).Error
	})
}

// GetOrg retrieves an organization by its UUID.
func (r *organizationRepository) GetOrg(ctx context.Context, orgID uuid.UUID) (*models.Organization, error) {
	var organization models.Organization
	err := r.db.WithContext(ctx).
		Where("id = ?", orgID).
		First(&organization).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to fetch organization: %w", err)
	}

	return &organization, nil
}
