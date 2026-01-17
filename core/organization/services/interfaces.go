package organization_services

import (
	"context"

	"github.com/google/uuid"
)

type OrganizationServiceInterface interface {
	GetOrgByID(ctx context.Context, orgId uuid.UUID)
	GetAllOrgProperty()
	GetOrgPropertyByID()
}
