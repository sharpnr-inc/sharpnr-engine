package org_property_repo

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	org_property_models "sharpnr-engine/core/org_property/models"
)

type organizationPropertyRepo struct {
	db *gorm.DB
}

// AddDomainOfProperty implements OrganizationPropertyRepositoryInterface.
func (o *organizationPropertyRepo) AddDomainOfProperty(ctx context.Context, domain org_property_models.OrganizationPropertyDomain) error {
	panic("unimplemented")
}

// AddPropertyAddress implements OrganizationPropertyRepositoryInterface.
func (o *organizationPropertyRepo) AddPropertyAddress(ctx context.Context, propertyAddress org_property_models.OrganizationPropertyAddress) error {
	panic("unimplemented")
}

// CreateProperty implements OrganizationPropertyRepositoryInterface.
func (o *organizationPropertyRepo) CreateProperty(ctx context.Context, property *org_property_models.OrganizationProperty) error {
	panic("unimplemented")
}

// GetAllPropertyByOrgID implements OrganizationPropertyRepositoryInterface.
func (o *organizationPropertyRepo) GetAllPropertyByOrgID(ctx context.Context, orgID uuid.UUID) ([]org_property_models.OrganizationProperty, error) {
	panic("unimplemented")
}

// GetPropertyByID implements OrganizationPropertyRepositoryInterface.
func (o *organizationPropertyRepo) GetPropertyByID(ctx context.Context, propertyID uuid.UUID) (org_property_models.OrganizationProperty, error) {
	panic("unimplemented")
}

func NewOrganizationPropertyRepo(db *gorm.DB) *organizationPropertyRepo {
	return &organizationPropertyRepo{
		db: db,
	}
}
