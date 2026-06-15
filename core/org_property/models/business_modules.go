package org_property_models

import "github.com/google/uuid"

// these are models which store the business operation modules data
type PropertyBusinessOperationModules struct {
	ID                     uuid.UUID
	OrganizationPropertyID uuid.UUID
	BusinessModulesID      uuid.UUID
	BusinessModuleName     string
}
