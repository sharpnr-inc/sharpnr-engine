package organization_contracts

import "github.com/google/uuid"

type OrganizationResult struct {
	ID        uuid.UUID
	SharpID   uuid.UUID
	Name      string
	RootEmail string
}

type DraftOrganizationInput struct {
	Email    string
	Mobile   *string
	Password string
}
