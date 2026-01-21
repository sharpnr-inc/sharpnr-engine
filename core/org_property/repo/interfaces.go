package org_property_repo

import (
	"context"
	org_property_models "sharpnr-engine/core/org_property/models"

	"github.com/google/uuid"
)

type OrganizationPropertyRepositoryInterface interface {
	CreateProperty(ctx context.Context, property *org_property_models.OrganizationProperty) error
	GetAllPropertiesByOrgID(ctx context.Context, orgID uuid.UUID) ([]*org_property_models.OrganizationProperty, error)
	GetPropertyByID(ctx context.Context, propertyID uuid.UUID) (*org_property_models.OrganizationProperty, error)

	AddPropertyAddress(ctx context.Context, propertyAddress *org_property_models.OrganizationPropertyAddress) error
	UpdatePropertyAddress(ctx context.Context, propertyID uuid.UUID, updates map[string]interface{}) error
	GetPropertyAddress(ctx context.Context, propertyID uuid.UUID) (*org_property_models.OrganizationPropertyAddress, error)

	AddDomainOfProperty(ctx context.Context, propertyDomain *org_property_models.OrganizationPropertyDomain) error
	UpdateDomainPropertyStatus(ctx context.Context, status string, propertyID uuid.UUID) error
}
