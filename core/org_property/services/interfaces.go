package org_property_services

import (
	"context"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"

	"github.com/google/uuid"
)

type OrganizationPropertyServicesInterfaces interface {
	CreatePropertyDefault(ctx context.Context, input org_property_contracts.CreateOrgPropertyInput) error
	GetPropertiesOfOrg(ctx context.Context, orgID uuid.UUID) ([]org_property_contracts.PropertiesReader, error)
	GetPropertyDetailsByID(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (*org_property_contracts.PropertyDetailsReader, error)

	// GetPropertyFacts backs this module's Sentinel fact resolver: the owning
	// org and the attribute values policies may condition on, for several
	// properties at once. Called on the authorization path.
	GetPropertyFacts(ctx context.Context, orgID uuid.UUID, propertyIDs []uuid.UUID) ([]org_property_contracts.PropertyFacts, error)

	// GetPropertyAttributeValues returns the distinct values an attribute takes
	// across the org, for the policy builder's value pickers. Read when the
	// catalog is fetched, not when a decision is made.
	GetPropertyAttributeValues(ctx context.Context, orgID uuid.UUID, attribute string) ([]string, error)
}
