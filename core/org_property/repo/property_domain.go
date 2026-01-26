package org_property_repo

import (
	"context"
	"errors"
	"fmt"
	org_property_models "sharpnr-engine/core/org_property/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (o *organizationPropertyRepo) AddDomainOfProperty(ctx context.Context, propertyDomain *org_property_models.OrganizationPropertyDomain) error {
	// We insert the domain record into the organization_property_domains table.
	// Ensure propertyDomain.PropertyID is populated to link it to the correct property.
	err := o.db.WithContext(ctx).Create(propertyDomain).Error

	if err != nil {
		return fmt.Errorf("failed to link domain to property: %w", err)
	}

	return nil
}

func (o *organizationPropertyRepo) UpdateDomainPropertyStatus(ctx context.Context, status string, propertyID uuid.UUID) error {
	// We target the Domain model and filter by the OrganizationPropertyID.
	// Using .Update() ensures only the 'status' column is modified.
	result := o.db.WithContext(ctx).
		Model(&org_property_models.OrganizationPropertyDomain{}).
		Where("organization_property_id = ?", propertyID).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	// Safety check: ensure a domain record actually exists for this property
	if result.RowsAffected == 0 {
		return fmt.Errorf("no domain found for property ID: %s", propertyID)
	}

	return nil
}

func (o *organizationPropertyRepo) GetPropertyDomain(ctx context.Context, propertyID uuid.UUID) (*org_property_models.OrganizationPropertyDomain, error) {
	var domain org_property_models.OrganizationPropertyDomain

	err := o.db.
		WithContext(ctx).
		Where("organization_property_id = ?", propertyID).
		First(&domain).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // or return custom ErrNotFound
		}
		return nil, err
	}

	return &domain, nil
}
