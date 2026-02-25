package org_property_repo

import (
	"context"
	"errors"
	"fmt"
	org_property_models "sharpnr-engine/core/org_property/models"

	"github.com/google/uuid"
)

func (o *organizationPropertyRepo) AddPropertyAddress(ctx context.Context, propertyAddress *org_property_models.OrganizationPropertyAddress) error {
	// .Create inserts the address record.
	// Ensure propertyAddress.PropertyID is set before calling this.
	err := o.db.WithContext(ctx).Create(propertyAddress).Error

	if err != nil {
		return fmt.Errorf("failed to add property address: %w", err)
	}

	return nil
}

func (o *organizationPropertyRepo) UpdatePropertyAddress(ctx context.Context, propertyID uuid.UUID, updates map[string]interface{}) error {
	// We target the Address model but filter by the OrganizationPropertyID
	result := o.db.WithContext(ctx).
		Model(&org_property_models.OrganizationPropertyAddress{}).
		Where("organization_property_id = ?", propertyID).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	// It's important to check if the record actually existed to update
	if result.RowsAffected == 0 {
		return errors.New("property address not found")
	}

	return nil
}

func (o *organizationPropertyRepo) GetPropertyAddress(ctx context.Context, propertyID uuid.UUID) (*org_property_models.OrganizationPropertyAddress, error) {
	var address org_property_models.OrganizationPropertyAddress

	// We filter using the foreign key OrganizationPropertyID.
	// .First() is ideal here as it returns an error if the address is missing.
	err := o.db.WithContext(ctx).
		Where("organization_property_id = ?", propertyID).
		First(&address).Error

	if err != nil {
		// Returns nil and the specific error (e.g., gorm.ErrRecordNotFound)
		return nil, err
	}

	return &address, nil
}
