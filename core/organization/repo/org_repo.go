package organization_repo

import (
	"context"
	"errors"
	"sharpnr-engine/core/organization/models"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *organizationRepo) CreateOrganization(
	ctx context.Context,
	org *models.Organization,
) error {

	if org == nil {
		return errors.New("organization is nil")
	}

	if org.ID == uuid.Nil {
		return errors.New("organization ID is required")
	}

	return r.db.
		WithContext(ctx).
		Create(org).
		Error
}

func (r *organizationRepo) GetOrganizationByID(ctx context.Context, orgID uuid.UUID) (*models.Organization, error) {
	var organization models.Organization

	// .First() searches by the primary key (ID).
	// It automatically adds "LIMIT 1" to the SQL query.
	err := r.db.WithContext(ctx).
		Where("id = ?", orgID).
		First(&organization).Error

	if err != nil {
		// If no record is found, GORM returns gorm.ErrRecordNotFound.
		// We return the error so the service layer can decide the HTTP status (e.g., 404).
		return nil, err
	}

	return &organization, nil
}

func (r *organizationRepo) GetOrganizationBySharpID(ctx context.Context, sharpID string) (*models.Organization, error) {
	var organization models.Organization

	// Convert input to uppercase to match the database format
	normalizedID := strings.ToUpper(strings.TrimSpace(sharpID))

	err := r.db.WithContext(ctx).
		Where("sharpnr_id = ?", normalizedID).
		First(&organization).Error

	if err != nil {
		return nil, err
	}

	return &organization, nil
}

func (r *organizationRepo) GetOrganizationByEmail(ctx context.Context, rootEmail string) (*models.Organization, error) {
	var organization models.Organization

	// We filter by the root_email field.
	// Using .First() ensures we get one record and returns an error if not found.
	err := r.db.WithContext(ctx).
		Where("root_email = ?", rootEmail).
		First(&organization).Error

	if err != nil {
		// Returns nil and gorm.ErrRecordNotFound if the email isn't registered
		return nil, err
	}

	return &organization, nil
}

func (r *organizationRepo) UpdateOrganizationStatus(ctx context.Context, orgID uuid.UUID, status string) error {
	// We target the Organization model and filter by its Primary Key
	// .Update() only changes the specified column and updates the 'updated_at' timestamp
	result := r.db.WithContext(ctx).
		Model(&models.Organization{}).
		Where("id = ?", orgID).
		Update("status", strings.ToUpper(status)) // Ensuring status follows your CAPS convention

	if result.Error != nil {
		return result.Error
	}

	// Check if the record actually existed
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (d *organizationRepo) AddOrgOfficeAddress(ctx context.Context, Addr *models.OrganizationOffice) error {
	return d.db.WithContext(ctx).Create(Addr).Error
}

// AddOrgVerificationDetails implements MainOrgSignUpRepositoryInterface.
func (d *organizationRepo) AddOrgVerificationDetails(ctx context.Context, orgVerification *models.OrganizationVerification) error {
	return d.db.WithContext(ctx).Create(orgVerification).Error
}

// CreateOrgSubscription implements MainOrgSignUpRepositoryInterface.
func (d *organizationRepo) CreateOrgSubscription(ctx context.Context, orgSub *models.OrganizationSubscription) error {
	return d.db.WithContext(ctx).Create(orgSub).Error
}

func (d *organizationRepo) GetOrgSubscriptionPlan(ctx context.Context, orgID uuid.UUID) (*models.OrganizationSubscription, error) {
	var subscription models.OrganizationSubscription

	// We filter by organization_id.
	// We order by 'created_at DESC' to ensure we get the most recent
	// subscription record if the history is kept in the same table.
	err := d.db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Order("created_at DESC").
		First(&subscription).Error

	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (d *organizationRepo) CreateOrganizationStatus(ctx context.Context, orgStatus *models.OrganizationSetupStatus) error {
	return d.db.WithContext(ctx).Create(orgStatus).Error
}
