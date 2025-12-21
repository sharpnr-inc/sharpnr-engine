package core_modules_events

import "github.com/google/uuid"

// FACT: Organization has been created successfully
type OrganizationCreated struct {
	OrganizationID   uuid.UUID
	OrganizationName string
	OrganizationType string
	Modules          []string
}
