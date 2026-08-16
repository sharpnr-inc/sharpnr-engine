package models

import (
	organization_constants "sharpnr-engine/core/organization/constants"
	"time"

	"github.com/google/uuid"
)

type OrganizationSharpnrServices struct {
	ID                uuid.UUID
	OrganizationID    uuid.UUID
	SharpnrServicesID uuid.UUID
	StartedAt         time.Time
}

type SharpnrServices struct {
	ID                            uuid.UUID
	OrganizationSharpnrServicesID uuid.UUID
	Service                       organization_constants.OrganizationSharpnrServices
	IsOn                          bool
}
