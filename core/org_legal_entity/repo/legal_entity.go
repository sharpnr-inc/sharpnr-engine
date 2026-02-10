package org_legal_entity_repo

import (
	"context"
	"errors"
	"fmt"
	org_legal_entity_models "sharpnr-engine/core/org_legal_entity/models"

	"github.com/google/uuid"
)

// CreateLegalEntity implements OrgLegalEntityRepoInterface.
func (o *orgLegalEntityRepo) CreateLegalEntity(ctx context.Context, legalEntity *org_legal_entity_models.OrganizationLegalEntity) error {
	// .Create() inserts the legal entity record into the organization_legal_entities table.
	// If you pass a pointer, legalEntity.ID will be updated with the new ID.
	if err := o.db.WithContext(ctx).Create(legalEntity).Error; err != nil {
		return err
	}

	return nil
}

// AddAddressInLegalEntity implements OrgLegalEntityRepoInterface.
func (o *orgLegalEntityRepo) AddAddressInLegalEntity(ctx context.Context, address *org_legal_entity_models.LegalEntityAddress) error {
	// .Create inserts the new address record.
	// Ensure address.LegalEntityID is set before calling this to link it correctly.
	err := o.db.WithContext(ctx).Create(address).Error

	if err != nil {
		return err
	}

	return nil
}

func (o *orgLegalEntityRepo) GetAddressOfLegalEntity(ctx context.Context, legalEntityID uuid.UUID) (*org_legal_entity_models.LegalEntityAddress, error) {
	var address org_legal_entity_models.LegalEntityAddress

	// We filter by the specific foreign key OrganizationLegalEntityID.
	// .First() is used assuming each legal entity has a primary/singular address.
	err := o.db.WithContext(ctx).
		Where("organization_legal_entity_id = ?", legalEntityID).
		First(&address).Error

	if err != nil {
		// Returns nil and the specific error (e.g., gorm.ErrRecordNotFound)
		return nil, err
	}

	return &address, nil
}

func (o *orgLegalEntityRepo) GetLegalEntityByOrgID(ctx context.Context, orgID uuid.UUID) ([]*org_legal_entity_models.OrganizationLegalEntity, error) {
	var legalEntities []*org_legal_entity_models.OrganizationLegalEntity

	// We filter by organization_id to retrieve all associated legal entities.
	// .Find() returns an empty slice and nil error if no records are found.
	err := o.db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Find(&legalEntities).Error

	if err != nil {
		return nil, err
	}

	return legalEntities, nil
}

func (o *orgLegalEntityRepo) GetLegalEntityByID(ctx context.Context, legalEntityID uuid.UUID) (*org_legal_entity_models.OrganizationLegalEntity, error) {
	var legalEntity org_legal_entity_models.OrganizationLegalEntity

	// .First() searches for the record matching the primary key.
	// It returns gorm.ErrRecordNotFound if no entry matches the ID.
	err := o.db.WithContext(ctx).
		Where("id = ?", legalEntityID).
		First(&legalEntity).Error

	if err != nil {
		return nil, err
	}

	return &legalEntity, nil
}

func (o *orgLegalEntityRepo) UpdateLegalEntityAddress(ctx context.Context, addressID uuid.UUID, updates map[string]interface{}) error {
	// Using .Updates(map) is safer as it only changes the fields provided
	result := o.db.WithContext(ctx).
		Model(&org_legal_entity_models.LegalEntityAddress{}).
		Where("id = ?", addressID).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("address record not found")
	}
	return nil
}

func (o *orgLegalEntityRepo) UpdateLegalEntityStatus(ctx context.Context, legalEntityID uuid.UUID, isActive bool) error {
	// We use .Model() to target the table and .Update() for the specific column.
	// This is more efficient than fetching the whole record first.
	result := o.db.WithContext(ctx).
		Model(&org_legal_entity_models.OrganizationLegalEntity{}).
		Where("id = ?", legalEntityID).
		Update("is_active", isActive)

	if result.Error != nil {
		return result.Error
	}

	// Safety check to ensure the ID actually existed
	if result.RowsAffected == 0 {
		return fmt.Errorf("legal entity not found: %s", legalEntityID)
	}

	return nil
}
