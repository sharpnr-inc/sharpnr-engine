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

	// GetPropertiesByIDs loads several properties in one query, scoped to the
	// org. Sentinel calls this on every instance-level property check, so it
	// must not be a loop over GetPropertyByID — and unlike that method it takes
	// organizationID, so a foreign id comes back absent rather than fetched and
	// then rejected by the caller.
	GetPropertiesByIDs(ctx context.Context, orgID uuid.UUID, propertyIDs []uuid.UUID) ([]*org_property_models.OrganizationProperty, error)
}
