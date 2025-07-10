package repositories

import (
	"github.com/google/uuid"
	"github.com/sharpnr-inc/sharpnr-engine/core/models"
	"gorm.io/gorm"
)

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *organizationRepository {
	return &organizationRepository{db}
}

// Create inserts a new organization record
func (r *organizationRepository) Create(org *models.Organization) error {
	return r.db.Create(org).Error
}

// ExistsBySharpnrOrgID checks if a SharpnrOrgID already exists
func (r *organizationRepository) ExistsBySharpnrOrgID(id string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Organization{}).
		Where("sharpnr_org_id = ?", id).
		Count(&count).Error
	return count > 0, err
}

// FindBySharpnrOrgID returns an org by its public ID
func (r *organizationRepository) FindBySharpnrOrgID(id string) (*models.Organization, error) {
	var org models.Organization
	err := r.db.Where("sharpnr_org_id = ?", id).First(&org).Error
	if err != nil {
		return nil, err
	}
	return &org, nil
}

// FindByUUID returns an org by its internal UUID
func (r *organizationRepository) FindByUUID(uuid uuid.UUID) (*models.Organization, error) {
	var org models.Organization
	err := r.db.Where("sharpnr_org_uuid = ?", uuid).First(&org).Error
	if err != nil {
		return nil, err
	}
	return &org, nil
}

// FindByEmail finds an org by its root email
func (r *organizationRepository) FindByEmail(email string, org *models.Organization) error {
	return r.db.Where("root_email = ?", email).First(org).Error
}

// Update allows partial updates using GORM's map update
func (r *organizationRepository) Update(uuid uuid.UUID, updates map[string]interface{}) error {
	return r.db.Model(&models.Organization{}).
		Where("sharpnr_org_uuid = ?", uuid).
		Updates(updates).Error
}
