package repositories

import (
	// "github.com/google/uuid"
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

// FindBySharpnrOrgID checks if an organization exists by its public ID.
// This method fulfills the CoreOrganizationRepository interface.
func (r *organizationRepository) FindBySharpnrOrgID(id string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Organization{}).Where("sharpnr_org_id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// FindByUUID returns an org by its internal UUID string.
func (r *organizationRepository) FindByUUID(uuid string) (*models.Organization, error) {
	var org models.Organization
	err := r.db.Where("sharpnr_org_uuid = ?", uuid).First(&org).Error
	if err != nil {
		return nil, err
	}
	return &org, nil
}

// FindByMail finds an org by its root email.
// This method fulfills the CoreOrganizationRepository interface.
func (r *organizationRepository) FindByMail(email string) (*models.Organization, error) {
	var org models.Organization
	// Assuming the column name is 'root_email' based on the previous implementation.
	err := r.db.Where("root_email = ?", email).First(&org).Error
	if err != nil {
		return nil, err
	}
	return &org, nil
}

// Update saves changes to an existing organization model.
// This method fulfills the CoreOrganizationRepository interface.
func (r *organizationRepository) Update(org *models.Organization) error {
	return r.db.Save(org).Error
}
