package models

import (
	property_constants "sharpnr-engine/core/org_property/constants"

	"github.com/google/uuid"
)

type OrganizationModulesRegister struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Domain         property_constants.PropertyDomainTypes
}


