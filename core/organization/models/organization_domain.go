package models

import (
	"github.com/google/uuid"
)

type OrganizationDomainModules struct {
	ID                   uuid.UUID
	OrganizationID       uuid.UUID
	OrganizationDomainID uuid.UUID
	Module               HospitalityDomainModuleType // Bar, Restaurant, Hotel
	Status               string
}
