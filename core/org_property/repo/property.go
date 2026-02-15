package org_property_repo

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	org_property_models "sharpnr-engine/core/org_property/models"

	"github.com/google/uuid"
)

// CreateProperty implements OrganizationPropertyRepositoryInterface.
func (o *organizationPropertyRepo) CreateProperty(ctx context.Context, property *org_property_models.OrganizationProperty) error {
	// .Create inserts the property record.
	// GORM will automatically populate property.ID and timestamps upon success.

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = o.db
	}
	err := db.WithContext(ctx).Create(property).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *organizationPropertyRepo) GetAllPropertiesByOrgID(ctx context.Context, orgID uuid.UUID) ([]*org_property_models.OrganizationProperty, error) {
	var properties []*org_property_models.OrganizationProperty

	// Fetching all properties linked to the specific Organization
	err := o.db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Order("created_at DESC"). // Showing newest properties first
		Find(&properties).Error

	if err != nil {
		return nil, err
	}

	return properties, nil
}

func (o *organizationPropertyRepo) GetPropertyByID(ctx context.Context, propertyID uuid.UUID) (*org_property_models.OrganizationProperty, error) {
	var property org_property_models.OrganizationProperty

	// .First() adds 'LIMIT 1' to the query and filters by the primary key.
	// If no record is found, it returns gorm.ErrRecordNotFound.
	err := o.db.WithContext(ctx).
		Where("id = ?", propertyID).
		First(&property).Error

	if err != nil {
		return nil, err
	}

	return &property, nil
}
