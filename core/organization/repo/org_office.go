package organization_repo

import (
	"context"
	"fmt"
	"sharpnr-engine/core/organization/models"
)

// CreateOrgOffice saves a new office record.
func (r *organizationRepo) CreateOrgOffice(ctx context.Context, office *models.OrganizationOffice) error {
	// .Create() inserts the struct and populates the ID/Timestamps automatically
	err := r.db.WithContext(ctx).Create(office).Error
	if err != nil {
		return fmt.Errorf("failed to create org office: %w", err)
	}
	return nil
}

// CreateOrgOfficeAddress saves the physical address linked to an office.
func (r *organizationRepo) CreateOrgOfficeAddress(ctx context.Context, officeAddr *models.OrganizationOfficeAddress) error {
	err := r.db.WithContext(ctx).Create(officeAddr).Error
	if err != nil {
		return fmt.Errorf("failed to create office address: %w", err)
	}
	return nil
}

// CreateOrgOfficeGeoLocation saves the latitude and longitude data.
func (r *organizationRepo) CreateOrgOfficeGeoLocation(ctx context.Context, geoLocation *models.OrganizationGeoLocation) error {
	err := r.db.WithContext(ctx).Create(geoLocation).Error
	if err != nil {
		return fmt.Errorf("failed to create office geo-location: %w", err)
	}
	return nil
}
