package organization_contracts

import "github.com/google/uuid"

type OrganizationResult struct {
	ID        uuid.UUID
	SharpID   string
	Name      string
	RootEmail string
}

type DraftOrganizationInput struct {
	Email    string
	Mobile   *string
	Password string
}

type InputDataForOrgAndEmp struct {
	Email      string
	HashedPass string
	Phone      string
}

type InputForCreateDefaultEmployee struct {
	OrgID      uuid.UUID
	Email      string
	HashedPass string
}
