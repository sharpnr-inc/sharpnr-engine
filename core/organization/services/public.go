package organization_services

import (
	"context"
	organization_contracts "sharpnr-engine/core/organization/contracts"

	"github.com/google/uuid"
)

func (r *organizationService) GetOrgByID(
	ctx context.Context,
	orgId uuid.UUID,
) (organization_contracts.OrganizationResult, error)

func (r *organizationService) GetOrgBySharpID(
	ctx context.Context,
	sharpID string,
) (organization_contracts.OrganizationResult, error)
